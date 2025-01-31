package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/maypok86/otter"
	"github.com/nerdwave-nick/nerdlocke/internal/api"
	"github.com/nerdwave-nick/nerdlocke/internal/api/health"
	intapi "github.com/nerdwave-nick/nerdlocke/internal/api/pokeapi"
	"github.com/nerdwave-nick/nerdlocke/internal/cache"
	"github.com/nerdwave-nick/nerdlocke/internal/frontend"
	"github.com/nerdwave-nick/pokeapi-go"
	"github.com/spf13/cobra"
)

type RootOptions struct {
	LogLevel    string
	DBPath      string
	GCInterval  int
	L2CacheTTL  int
	L1CacheTTL  int
	L1CacheSize int
	Port        int
}

func (o *RootOptions) Validate() error {
	concatErr := func(err error, olderr error) error {
		if olderr != nil {
			return fmt.Errorf("%s\n%w", err.Error(), olderr)
		}
		return err
	}
	var err error
	if o.DBPath == "" {
		err = concatErr(fmt.Errorf("db-path can't be empty"), err)
	}
	if o.L2CacheTTL <= 0 {
		err = concatErr(fmt.Errorf("l2-ttl must be greater than 0"), err)
	}
	if o.L1CacheTTL <= 0 {
		err = concatErr(fmt.Errorf("l1-ttlmust be greater than 0"), err)
	}
	if o.L1CacheSize <= 0 {
		err = concatErr(fmt.Errorf("l1-size must be greater than 0"), err)
	}
	if o.GCInterval <= 0 {
		err = concatErr(fmt.Errorf("gc-interval must be greater than 0"), err)
	}
	if o.Port <= 0 {
		err = concatErr(fmt.Errorf("port must be greater than 0"), err)
	}
	return err
}

var rootOpts = &RootOptions{}

func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&rootOpts.DBPath, "db-path", ".badger", "The path of the badger db folder. Will be created when it doesn't exist.")
	rootCmd.Flags().IntVar(&rootOpts.GCInterval, "gc-interval", 600, "The garbage collection interval of the badger db in seconds. Needs to be greater than 0.")
	rootCmd.Flags().IntVar(&rootOpts.L2CacheTTL, "l2-ttl", 86400, "The ttl of the larger l2 cache in seconds. Needs to be greater than 0.")
	rootCmd.Flags().IntVar(&rootOpts.L1CacheTTL, "l1-ttl", 7200, "The ttl of the smaller l1 cache in seconds. Needs to be greater than 0.")
	rootCmd.Flags().IntVar(&rootOpts.L1CacheSize, "l1-size", 2000, "The size of the smaller l1 cache in number of items. Needs to be greater than 0.")
	rootCmd.Flags().IntVarP(&rootOpts.Port, "port", "p", 8080, "The port to listen on")
	rootCmd.Flags().StringVarP(&rootOpts.LogLevel, "level", "l", "info", "The log level. Valid levels are debug, info, warn, and error.")
}

var rootCmd = &cobra.Command{
	Use:   "nerdlocke",
	Short: "nerdlocke - the tracking and knowledge tool for nuzlockes and soullockes in pokemon",
	Long:  "nerdlocke - the tracking and knowledge tool for nuzlockes and soullockes in pokemon\n\nProvides a frontend for user interaction and a backend api including persistent K/V database and caching of pokeapi",
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		if err := rootOpts.Validate(); err != nil {
			return fmt.Errorf("incorrect command usage:\n%w\n", err)
		}
		switch strings.ToLower(rootOpts.LogLevel) {
		case "debug":
			slog.SetLogLoggerLevel(slog.LevelDebug)
		case "info":
			slog.SetLogLoggerLevel(slog.LevelInfo)
		case "warn":
			slog.SetLogLoggerLevel(slog.LevelWarn)
		case "error":
			slog.SetLogLoggerLevel(slog.LevelError)
		default:
			slog.Warn("no/invalid log level provided, setting to info")
			slog.SetLogLoggerLevel(slog.LevelInfo)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, _ []string) {
		err := rootMain(cmd.Context())
		if err != nil {
			os.Exit(1)
		}
	},
}

func badgerBackgroundGC(ctx context.Context, db *badger.DB, gcInterval time.Duration) {
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

func stopServerWithTimeout(server *http.Server) error {
	slog.Debug("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("shutting down http server", slog.Any("error", err))
		return err
	}
	return nil
}

type BadgerLoggerWrapper struct{}

func (*BadgerLoggerWrapper) Errorf(format string, args ...interface{}) {
	slog.Error("badger " + strings.TrimSuffix(fmt.Sprintf(format, args...), "\n"))
}

func (*BadgerLoggerWrapper) Warningf(format string, args ...interface{}) {
	slog.Warn("badger " + strings.TrimSuffix(fmt.Sprintf(format, args...), "\n"))
}

func (*BadgerLoggerWrapper) Infof(format string, args ...interface{}) {
	slog.Info("badger " + strings.TrimSuffix(fmt.Sprintf(format, args...), "\n"))
}

func (*BadgerLoggerWrapper) Debugf(format string, args ...interface{}) {
	slog.Debug("badger " + strings.TrimSuffix(fmt.Sprintf(format, args...), "\n"))
}

func rootMain(parentCtx context.Context) error {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()
	// persistent badger db and cache wrapper
	db, err := badger.Open(badger.DefaultOptions(rootOpts.DBPath).WithLogger(&BadgerLoggerWrapper{}))
	if err != nil {
		return err
	}
	defer func() {
		err = db.Close()
		if err != nil {
			slog.Error("shutting down db", slog.Any("error", err))
		}
	}()

	boltCache := cache.NewBadgerCache(db, time.Duration(rootOpts.L2CacheTTL)*time.Second)

	// in memory otter cache
	oc, err := otter.MustBuilder[string, []byte](rootOpts.L1CacheSize).
		WithTTL(time.Duration(rootOpts.L2CacheTTL) * time.Second).
		Build()
	if err != nil {
		return err
	}
	otterCache := cache.NewOtterCache(&oc)

	// multi layer cache with preference for the in memory cache
	multiCache := cache.NewMultiLayerCache(
		otterCache,
		&boltCache,
	)

	papiClient := pokeapi.NewClient(multiCache, *http.DefaultClient)
	b, err := papiClient.Berry("0")
	slog.Info("berry", slog.Any("berry", b), slog.Any("err", err))

	frontend, err := frontend.GetAssetFS()
	if err != nil {
		return err
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
		Addr:         fmt.Sprintf(":%d", rootOpts.Port),
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

	badgerBackgroundGC(ctx, db, time.Duration(rootOpts.GCInterval)*time.Second)
	slog.Info("badger db background gc started...")
	go func() {
		defer cancel()
		slog.Info("server ready to listen...", slog.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			slog.Error("error in listen and serve", slog.Any("error", err))
		}
	}()

	<-ctx.Done()
	return stopServerWithTimeout(server)
}
