package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

type BadgerCache struct {
	db  *badger.DB
	TTL time.Duration
}

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
