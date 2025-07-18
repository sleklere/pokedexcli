package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	cache "github.com/sleklere/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func GetLocationAreas(pageURL *string, cache *cache.Cache) (LocationAreasRes, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationAreasRes LocationAreasRes

	if entry, ok := cache.Get(url); ok {
		err := json.Unmarshal(entry, &locationAreasRes)
		if err != nil {
			return LocationAreasRes{}, err
		}
		return locationAreasRes, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreasRes{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasRes{}, err
	}

	cache.Add(url, body)

	err = json.Unmarshal(body, &locationAreasRes)
	if err != nil {
		return LocationAreasRes{}, err
	}

	return locationAreasRes, nil
}
