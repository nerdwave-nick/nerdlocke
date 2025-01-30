package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

type BoltCache struct {
	db  *bbolt.DB
	TTL time.Duration
}

type boltCacheTtlItem struct {
	ValidUntil time.Time       `json:"valid_until"`
	Value      json.RawMessage `json:"value"`
}

var bucket = []byte("papi_cache")

func NewBoltCache(db *bbolt.DB, ttl time.Duration) (*BoltCache, error) {
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
	return &BoltCache{db: db, TTL: ttl}, nil
}

func (c *BoltCache) putItem(key string, value json.RawMessage) error {
	boltCacheItem := boltCacheTtlItem{ValidUntil: time.Now().Add(c.TTL), Value: value}
	bytes, err := json.Marshal(boltCacheItem)
	if err != nil {
		return err
	}
	return c.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		err = b.Put([]byte(key), bytes)
		return err
	})
}

func (c *BoltCache) Set(endpoint string, value any) error {
	fmt.Printf("writing to bolt cache: %q\n", endpoint)
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.putItem(endpoint, bytes)
}

var boltBucketNotFoundError = errors.New("bucket not found")

func (c *BoltCache) getItem(key string, value any) (bool, error) {
	var bytes []byte = nil
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return boltBucketNotFoundError
		}
		serializedItem := b.Get([]byte(key))
		if serializedItem == nil {
			fmt.Printf("a not found")
			return nil
		}
		// avoid blocking the db
		bytes = make([]byte, len(serializedItem))
		copy(bytes, serializedItem)
		fmt.Printf("copied")
		return nil
	})
	if err != nil {
		return false, err
	}
	if bytes == nil {
		fmt.Printf("bytes are nil")
		return false, nil
	}
	boltCacheItem := boltCacheTtlItem{}
	err = json.Unmarshal(bytes, &boltCacheItem)
	if err != nil {
		return false, err
	}
	// past ttl
	if boltCacheItem.ValidUntil.Before(time.Now()) {
		return false, nil
	}
	return true, json.Unmarshal(boltCacheItem.Value, value)
}

func (c *BoltCache) Get(endpoint string, value any) (bool, error) {
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
