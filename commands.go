package main

import (
	"fmt"
	"os"
	api "github.com/sleklere/pokedexcli/internal/pokeapi"
	cache "github.com/sleklere/pokedexcli/internal/pokecache"
)

func commandExit(config *config, cache *cache.Cache) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

func commandHelp(config *config, cache *cache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cliCommand := range getCommands() {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}

	return nil
}

func commandMapForward(config *config, cache *cache.Cache) error {
	locationAreasRes, err := api.GetLocationAreas(config.nextLocationsURL, cache)
	if err != nil {
		fmt.Printf("Error getting location areas: %v\n", err)
		return err
	}

	config.nextLocationsURL = locationAreasRes.Next
	config.previousLocationsURL = locationAreasRes.Previous

	for _, location := range locationAreasRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(config *config, cache *cache.Cache) error {
	if config.previousLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locationAreasRes, err := api.GetLocationAreas(config.previousLocationsURL, cache)
	if err != nil {
		fmt.Printf("Error getting location areas: %v\n", err)
		return err
	}

	config.nextLocationsURL = locationAreasRes.Next
	config.previousLocationsURL = locationAreasRes.Previous

	for _, location := range locationAreasRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}


type cliCommand struct {
	name string
	description string
	callback func(*config, *cache.Cache) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: 	"exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: 	"help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: 	"map",
			description: "Displays the names of the next 20 location areas in the Pokemon world.",
			callback: commandMapForward,
		},
		"mapb": {
			name: 	"mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			callback: commandMapBack,
		},
	}
}
