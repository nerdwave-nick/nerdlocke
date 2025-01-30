package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/maypok86/otter"
	"github.com/nerdwave-nick/nerdlocke/internal/pokeapi"
)

type MultiLayerCache struct {
	caches []pokeapi.Cache
}

func NewMultiLayerCache(caches ...pokeapi.Cache) *MultiLayerCache {
	return &MultiLayerCache{caches: caches}
}

func (c *MultiLayerCache) Set(endpoint string, value any) error {
	fmt.Printf("writing to multi layer cache: %q\n", endpoint)
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
	indexFound := -1
	for i, cache := range c.caches {
		found, err := cache.Get(endpoint, value)
		if err != nil {
			return found, err
		}
		if found {
			indexFound = i
			break
		}
	}
	if indexFound >= 0 {
		for _, cache := range c.caches[:indexFound] {
			_ = cache.Set(endpoint, value)
		}
	}
	return indexFound >= 0, nil
}

type BadgerCache struct {
	db  *badger.DB
	TTL time.Duration
}

var bucket = []byte("papi_cache")

func NewBadgerCache(db *badger.DB, ttl time.Duration) BadgerCache {
	return BadgerCache{db: db, TTL: ttl}
}

func (c *BadgerCache) putItem(key string, value []byte) error {
	return c.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), value).WithTTL(c.TTL)
		err := txn.SetEntry(e)
		return err
	})
}

func (c *BadgerCache) Set(endpoint string, value any) error {
	fmt.Printf("writing to badger cache: %q\n", endpoint)
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.putItem(endpoint, bytes)
}

var boltBucketNotFoundError = errors.New("bucket not found")

func (c *BadgerCache) getItem(key string, value any) (bool, error) {
	var bytes []byte = nil
	err := c.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		bytes, err = item.ValueCopy(bytes)
		if err != nil {
			return err
		}
		return nil
	})
	if errors.Is(err, badger.ErrKeyNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if bytes == nil {
		return false, nil
	}
	return true, json.Unmarshal(bytes, value)
}

func (c *BadgerCache) Get(endpoint string, value any) (bool, error) {
	found, err := c.getItem(endpoint, value)
	if err != nil {
		fmt.Printf("error checking in bolt cache: %q, %v\n", endpoint, err)
		return false, err
	}
	if !found {
		fmt.Printf("not found in bolt cache: %q\n", endpoint)
		return false, nil
	}
	fmt.Printf("found in bolt cache: %q\n", endpoint)
	return true, nil
}

type OtterCache struct {
	cache *otter.Cache[string, []byte]
}

func NewOtterCache(c *otter.Cache[string, []byte]) *OtterCache {
	return &OtterCache{cache: c}
}

func (c *OtterCache) Set(endpoint string, value any) error {
	fmt.Printf("writing to otter cache: %q\n", endpoint)
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
	fmt.Printf("found in otter cache: %q\n", endpoint)
	return true, nil
}
