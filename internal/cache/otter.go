package cache

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/maypok86/otter"
)

type OtterCache struct {
	cache *otter.Cache[string, []byte]
}

func NewOtterCache(c *otter.Cache[string, []byte]) *OtterCache {
	return &OtterCache{cache: c}
}

func (c *OtterCache) Set(endpoint string, value any) error {
	slog.Debug(fmt.Sprintf("writing to otter cache: %q\n", endpoint))
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
		slog.Debug(fmt.Sprintf("not found in otter cache: %q\n", endpoint))
		return false, nil
	}
	err := json.Unmarshal(bytes, value)
	if err != nil {
		slog.Debug(fmt.Sprintf("error unmarshalling from otter cache: %q\n", endpoint))
		return true, err
	}
	slog.Debug(fmt.Sprintf("found in otter cache: %q\n", endpoint))
	return true, nil
}
