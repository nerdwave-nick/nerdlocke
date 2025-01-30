package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/dgraph-io/badger"
	"github.com/maypok86/otter"
	"github.com/nerdwave-nick/nerdlocke/internal/api"
	"github.com/nerdwave-nick/nerdlocke/internal/api/health"
	intapi "github.com/nerdwave-nick/nerdlocke/internal/api/pokeapi"
	"github.com/nerdwave-nick/nerdlocke/internal/frontend"
	"github.com/nerdwave-nick/nerdlocke/internal/pokeapi"
)

func badgerBackgroundHandling(ctx context.Context, db *badger.DB, gcInterval time.Duration) {
	go func() {
		for {
			select {
			case <-time.After(gcInterval):
				err := db.RunValueLogGC(0.5)
				if err != nil {
					if !errors.Is(err, badger.ErrNoRewrite) {
						slog.Error("running the badger db gc", slog.Any("error", err))
					}
				}
			case <-ctx.Done():
				slog.Debug("badger gc loop shut down")
				return
			}
		}
	}()
}

func startServer(ctx context.Context, opts *Options, server *http.Server, db *badger.DB) func() {
	return func() {
		// start background gc and closing handler
		badgerBackgroundHandling(ctx, db, time.Duration(opts.GCInterval)*time.Second)
		slog.Info("badger db background gc started...")

		slog.Info("server ready to listen...", slog.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			slog.Error("error in listen and serve", slog.Any("error", err))
			log.Fatal(err)
		}
	}
}

func stopServerWithTimeout(cancelRunningProcesses context.CancelFunc, server *http.Server, db *badger.DB) func() {
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			slog.Error("shutting down http server", slog.Any("error", err))
		}

		cancelRunningProcesses()
		err = db.Close()
		if err != nil {
			slog.Error("shutting down db", slog.Any("error", err))
		}
	}
}

type Options struct {
	DBPath      string `doc:"The path of the badger db folder" default:".badger"`
	GCInterval  int    `doc:"The garbage collection interval of the badger db in seconds" default:"600"`
	L2CacheTTL  int    `doc:"The ttl of the larger l2 cache in seconds" default:"86400"`
	L1CacheTTL  int    `doc:"The ttl of the smaller l1 in-memory cache in seconds" default:"7200"`
	L1CacheSize int    `doc:"The size of the smaller l1 in-memory cache in number of items" default:"2000"`
	Port        int    `doc:"Port to listen on." short:"p" default:"8080"`
}

func (o *Options) Validate() error {
	var err error
	if o.DBPath == "" {
		fmt.Printf("The DBPath can't be empty")
		err = fmt.Errorf("db path messed up: %q - %w", o.DBPath, err)
	}
	if o.L2CacheTTL <= 0 {
		fmt.Printf("The L2CacheTTL must be greater than 0")
		err = fmt.Errorf("l2 ttl messed up: %q - %w", o.L2CacheTTL, err)
	}
	if o.L1CacheTTL <= 0 {
		fmt.Printf("The L1CacheTTL must be greater than 0")
		err = fmt.Errorf("l1 ttl messed up: %q - %w", o.L1CacheTTL, err)
	}
	if o.L1CacheSize <= 0 {
		fmt.Printf("The L1CacheSize must be greater than 0")
		err = fmt.Errorf("l1 size messed up: %q - %w", o.L1CacheTTL, err)
	}
	if o.GCInterval <= 0 {
		fmt.Printf("The GCInterval must be greater than 0")
		err = fmt.Errorf("gc interval messed up: %d - %w", o.GCInterval, err)
	}
	if o.Port <= 0 {
		fmt.Printf("The GCInterval must be valid")
		err = fmt.Errorf("port messed up: %d - %w", o.Port, err)
	}
	return err
}

type BadgerLoggerWrapper struct{}

func (*BadgerLoggerWrapper) Errorf(format string, args ...interface{}) {
	slog.Error(strings.TrimSuffix(fmt.Sprintf(format, args...), "\n"), slog.String("module", "badger"))
}

func (*BadgerLoggerWrapper) Warningf(format string, args ...interface{}) {
	slog.Warn(strings.TrimSuffix(fmt.Sprintf(format, args...), "\n"), slog.String("module", "badger"))
}

func (*BadgerLoggerWrapper) Infof(format string, args ...interface{}) {
	slog.Info(strings.TrimSuffix(fmt.Sprintf(format, args...), "\n"), slog.String("module", "badger"))
}

func (*BadgerLoggerWrapper) Debugf(format string, args ...interface{}) {
	slog.Debug(strings.TrimSuffix(fmt.Sprintf(format, args...), "\n"), slog.String("module", "badger"))
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		err := opts.Validate()
		if err != nil {
			os.Exit(1)
		}
		// persistent badger db and cache wrapper
		db, err := badger.Open(badger.DefaultOptions(opts.DBPath).WithLogger(&BadgerLoggerWrapper{}))
		if err != nil {
			panic(err)
		}

		boltCache := NewBadgerCache(db, time.Duration(opts.L2CacheTTL)*time.Second)

		// in memory otter cache
		oc, err := otter.MustBuilder[string, []byte](opts.L1CacheSize).
			WithTTL(time.Duration(opts.L2CacheTTL) * time.Second).
			Build()
		if err != nil {
			panic(err)
		}
		otterCache := NewOtterCache(&oc)

		// multi layer cache with preference for the in memory cache
		multiCache := NewMultiLayerCache(
			otterCache,
			&boltCache,
		)

		_ = pokeapi.NewClient(multiCache, *http.DefaultClient)

		frontend, err := frontend.GetAssetFS()
		if err != nil {
			panic(err)
		}

		mux := http.NewServeMux()
		mux.Handle("/", http.FileServer(http.FS(frontend)))

		healthController := health.MakeController()
		pokeapiController := intapi.MakeController()

		router := api.MakeRouter(
			mux,
			[]api.Controller{
				healthController,
				pokeapiController,
			},
		)
		slog.Debug("router created, proceeding to start backend...")

		server := &http.Server{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 30 * time.Second,
			Addr:         fmt.Sprintf(":%d", opts.Port),
			Handler:      router,
		}

		// http.HandleFunc("/pokeapi", func(w http.ResponseWriter, r *http.Request) {
		// 	route := r.URL.Query().Get("route")
		// 	id := r.URL.Query().Get("id")
		// 	var b any
		// 	var err error
		// 	switch route {
		// 	case "berry":
		// 		b, err = pokeapiClient.Berry(id)
		// 	case "berry-flavor":
		// 		b, err = pokeapiClient.BerryFlavor(id)
		// 	case "berry-firmness":
		// 		b, err = pokeapiClient.BerryFirmness(id)
		// 	case "contest-type":
		// 		b, err = pokeapiClient.ContestType(id)
		// 	case "contest-effect":
		// 		b, err = pokeapiClient.ContestEffect(id)
		// 	case "super-contest-effect":
		// 		b, err = pokeapiClient.SuperContestEffect(id)
		// 	case "encounter-method":
		// 		b, err = pokeapiClient.EncounterMethod(id)
		// 	case "encounter-condition":
		// 		b, err = pokeapiClient.EncounterCondition(id)
		// 	case "encounter-condition-value":
		// 		b, err = pokeapiClient.EncounterConditionValue(id)
		// 	case "pokemon":
		// 		if id != "" {
		// 			b, err = pokeapiClient.Pokemon(id)
		// 		} else {
		// 			b, err = pokeapiClient.Pokemons(10, 0)
		// 		}
		// 	case "tryfucky":
		// 		b, err = pokeapiClient.Pokemon("https://pokeapi.co/api/v2/pokemon/121")
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// 	if err != nil {
		// 		_, _ = w.Write([]byte(err.Error()))
		// 		return
		// 	}
		// 	bytes, err := json.Marshal(b)
		// 	if err != nil {
		// 		_, _ = w.Write([]byte(err.Error()))
		// 		return
		// 	}
		// 	_, _ = w.Write(bytes)
		// })

		hooks.OnStart(startServer(ctx, opts, server, db))
		hooks.OnStop(stopServerWithTimeout(cancel, server, db))
	})

	// Run the thing!
	cli.Run()
}
