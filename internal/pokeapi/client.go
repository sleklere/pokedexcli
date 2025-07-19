package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/sleklere/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasRes, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationAreasRes LocationAreasRes

	if entry, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(entry, &locationAreasRes)
		if err != nil {
			return LocationAreasRes{}, err
		}
		return locationAreasRes, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreasRes{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasRes{}, err
	}

	c.cache.Add(url, body)

	err = json.Unmarshal(body, &locationAreasRes)
	if err != nil {
		return LocationAreasRes{}, err
	}

	return locationAreasRes, nil
}

func (c *Client) GetLocationAreaByName(name string) (*LocationAreaRes, error) {
	url := baseURL + "/location-area/" + name

	var result *LocationAreaRes

	if entry, ok := c.cache.Get(url); ok {
		result = &LocationAreaRes{}
		err := json.Unmarshal(entry, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)

	result = &LocationAreaRes{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetPokemonByName(name string) (*Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	var result *Pokemon

	if cachedData, ok := c.cache.Get(url); ok {
		result = &Pokemon{}
		err := json.Unmarshal(cachedData, result)
		if err != nil {
			return nil, err
		}
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)

	result = &Pokemon{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
