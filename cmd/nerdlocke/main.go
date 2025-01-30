package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/maypok86/otter"
	"github.com/nerdwave-nick/nerdlocke/internal/frontend"
	"github.com/nerdwave-nick/nerdlocke/internal/pokeapi"
	"go.etcd.io/bbolt"
)

const boltdbPath = "./cache.bolt"

func main() {
	// persistent bolt db and bolt cache
	db, err := bbolt.Open(boltdbPath, 0600, nil)
	if err != nil {
		panic(err)
	}
	boltCache, err := NewBoltCache(db, 1*time.Minute)
	if err != nil {
		panic(err)
	}

	// in memory otter cache
	oc, err := otter.MustBuilder[string, []byte](10_000).
		WithTTL(30 * time.Second).
		DeletionListener(func(key string, _ []byte, cause otter.DeletionCause) {
			fmt.Printf("deleting %q from otter cache due to %v\n", key, cause)
		}).
		Build()
	if err != nil {
		panic(err)
	}
	otterCache := NewOtterCache(&oc)

	oc2, err := otter.MustBuilder[string, []byte](100).
		WithTTL(10 * time.Second).
		DeletionListener(func(key string, _ []byte, cause otter.DeletionCause) {
			fmt.Printf("deleting %q from otter cache due to %v\n", key, cause)
		}).
		Build()
	if err != nil {
		panic(err)
	}
	otterCache2 := NewOtterCache(&oc2)

	// multi layer cache with preference for the in memory cache
	multiCache := NewMultiLayerCache(
		otterCache2,
		otterCache,
		boltCache,
	)

	pokeapiClient := pokeapi.NewClient(multiCache, *http.DefaultClient)
	frontend, err := frontend.GetAssetFS()
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(frontend)))
	http.HandleFunc("/pokeapi", func(w http.ResponseWriter, r *http.Request) {
		route := r.URL.Query().Get("route")
		id := r.URL.Query().Get("id")
		var b any
		var err error
		switch route {
		case "berry":
			b, err = pokeapiClient.Berry(id)
		case "berry-flavor":
			b, err = pokeapiClient.BerryFlavor(id)
		case "berry-firmness":
			b, err = pokeapiClient.BerryFirmness(id)
		case "contest-type":
			b, err = pokeapiClient.ContestType(id)
		case "contest-effect":
			b, err = pokeapiClient.ContestEffect(id)
		case "super-contest-effect":
			b, err = pokeapiClient.SuperContestEffect(id)
		case "encounter-method":
			b, err = pokeapiClient.EncounterMethod(id)
		case "encounter-condition":
			b, err = pokeapiClient.EncounterCondition(id)
		case "encounter-condition-value":
			b, err = pokeapiClient.EncounterConditionValue(id)
		case "pokemon":
			if id != "" {
				b, err = pokeapiClient.Pokemon(id)
			} else {
				b, err = pokeapiClient.Pokemons(10, 0)
			}
		case "tryfucky":
			b, err = pokeapiClient.Pokemon("https://pokeapi.co/api/v2/pokemon/121")
		}
		w.WriteHeader(http.StatusOK)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		bytes, err := json.Marshal(b)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, _ = w.Write(bytes)
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
