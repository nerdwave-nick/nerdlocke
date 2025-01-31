package cache

import (
	"fmt"
	"log/slog"

	"github.com/nerdwave-nick/pokeapi-go"
)

type MultiLayerCache struct {
	caches []pokeapi.Cache
}

func NewMultiLayerCache(caches ...pokeapi.Cache) *MultiLayerCache {
	return &MultiLayerCache{caches: caches}
}

func (c *MultiLayerCache) Set(endpoint string, value any) error {
	slog.Debug(fmt.Sprintf("writing to multi layer cache: %q\n", endpoint))
	for _, cache := range c.caches {
		err := cache.Set(endpoint, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *MultiLayerCache) Get(endpoint string, value any) (bool, error) {
	slog.Debug(fmt.Sprintf("getting from multi layer cache: %q\n", endpoint))
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
