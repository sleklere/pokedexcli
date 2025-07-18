package main

import (
	cache "github.com/sleklere/pokedexcli/internal/pokecache"
	"strings"
	"time"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.TrimSpace(text))
}

func main() {
	apiConfig := &config{}

	cache := cache.NewCache(5*time.Second)

	startRepl(apiConfig, cache)
}
