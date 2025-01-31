package cache

import (
	"encoding/json"
	"fmt"

	"github.com/maypok86/otter"
)

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
