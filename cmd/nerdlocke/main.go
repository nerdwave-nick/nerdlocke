package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nerdwave-nick/nerdlocke/internal/frontend"
	"github.com/nerdwave-nick/nerdlocke/internal/pokeapi"
	"go.etcd.io/bbolt"
)

const boltdbPath = "./cache.bolt"

func main() {
	//...

	db, err := bbolt.Open(boltdbPath, 0600, nil)
	if err != nil {
		panic(err)
	}
	cache, err := NewBoltCache(db)
	if err != nil {
		panic(err)
	}
	pokeapiClient := pokeapi.NewClient(cache, *http.DefaultClient)
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

	//...
}
