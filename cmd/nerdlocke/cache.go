package main

import (
	"encoding/json"
	"fmt"

	"go.etcd.io/bbolt"
)

type BoltCache struct {
	db *bbolt.DB
}

var bucket = []byte("pokeapi")

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
	fmt.Printf("writing to cache: %q: %+v\n", endpoint, value)
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
		fmt.Printf("error checking in cache: %q\n", endpoint)
		return false, err
	}
	if !found {
		fmt.Printf("not found in cache: %q\n", endpoint)
		return false, nil
	}
	fmt.Printf("found in cache: %q:%+v\n", endpoint, value)
	return true, nil
}
