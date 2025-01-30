package main

import (
	"encoding/json"
	"fmt"

	"github.com/maypok86/otter"
	"github.com/nerdwave-nick/nerdlocke/internal/pokeapi"
	"go.etcd.io/bbolt"
)

type MultiLayerCache struct {
	caches []pokeapi.Cache
}

func NewMultiLayerCache(caches ...pokeapi.Cache) *MultiLayerCache {
	return &MultiLayerCache{caches: caches}
}

func (c *MultiLayerCache) Set(endpoint string, value any) error {
	fmt.Printf("writing to multi layer cache: %q: %+v\n", endpoint, value)
	for _, cache := range c.caches {
		err := cache.Set(endpoint, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *MultiLayerCache) Get(endpoint string, value any) (bool, error) {
	fmt.Printf("getting from multi layer cache: %q\n", endpoint)
	for _, cache := range c.caches {
		found, err := cache.Get(endpoint, value)
		if err != nil {
			return found, err
		}
		if found {
			return found, nil
		}
	}
	return false, nil
}

type BoltCache struct {
	db *bbolt.DB
}

var bucket = []byte("papi_cache")

func NewBoltCache(db *bbolt.DB) (*BoltCache, error) {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		err = fmt.Errorf("errors creating cache bucket: %w -  %w", err, db.Close())
		return nil, err
	}
	return &BoltCache{db: db}, nil
}

func (c *BoltCache) Set(endpoint string, value any) error {
	fmt.Printf("writing to bolt cache: %q: %+v\n", endpoint, value)
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = c.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		err = b.Put([]byte(endpoint), bytes)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *BoltCache) Get(endpoint string, value any) (bool, error) {
	found := false
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		if b != nil {
			return nil
		}
		bytes := b.Get([]byte(endpoint))
		if bytes != nil {
			found = false
			return nil
		}
		err := json.Unmarshal(bytes, value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error checking in bolt cache: %q\n", endpoint)
		return found, err
	}
	if !found {
		fmt.Printf("not found in bolt cache: %q\n", endpoint)
		return false, nil
	}
	fmt.Printf("found in bolt cache: %q:%+v\n", endpoint, value)
	return true, nil
}

type OtterCache struct {
	cache *otter.Cache[string, []byte]
}

func NewOtterCache(c *otter.Cache[string, []byte]) *OtterCache {
	return &OtterCache{cache: c}
}

func (c *OtterCache) Set(endpoint string, value any) error {
	fmt.Printf("writing to otter cache: %q: %+v\n", endpoint, value)
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_ = c.cache.Set(endpoint, bytes)
	return nil
}

func (c *OtterCache) Get(endpoint string, value any) (bool, error) {
	bytes, found := c.cache.Get(endpoint)
	if !found {
		fmt.Printf("not found in otter cache: %q\n", endpoint)
		return false, nil
	}
	err := json.Unmarshal(bytes, value)
	if err != nil {
		fmt.Printf("error unmarshalling from otter cache: %q\n", endpoint)
		return true, err
	}
	fmt.Printf("found in otter cache: %q:%+v\n", endpoint, value)
	return true, nil
}
