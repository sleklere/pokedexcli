package main

import (
	api "github.com/sleklere/pokedexcli/internal/pokeapi"
	"strings"
	"time"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.TrimSpace(text))
}

func main() {
	client := api.NewClient(5*time.Second, 5*time.Second)
	apiConfig := &config{
		pokeApiClient: client,
	}

	startRepl(apiConfig)
}
