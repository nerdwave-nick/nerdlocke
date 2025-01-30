package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	base_url     = "https://pokeapi.co/api/v2/"
	limit_param  = "limit"
	offset_param = "offset"
)

type Cache interface {
	// set, with value being a structure
	Set(endpoint string, value any) error
	// get, with return values being first an unmarshalled structure, then bool whether somethnig was found, and then error if something went wrong
	Get(endpoint string, value any) (bool, error)
}

type Client struct {
	cache  Cache
	client http.Client
}

func NewClient(cache Cache, client http.Client) *Client {
	return &Client{
		cache:  cache,
		client: client,
	}
}

func do[T any](c *Client, endpoint string) (*T, error) {
	value := new(T)
	found, err := c.cache.Get(endpoint, value)
	if err != nil {
		return nil, err
	}
	if found {
		return value, nil
	}

	resp, err := c.client.Get(base_url + endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	v := new(T)
	err = json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(endpoint, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
