package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(param string, config *config) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

func commandHelp(param string, config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cliCommand := range getCommands() {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}

	return nil
}

func commandMapForward(param string, config *config) error {
	locationAreasRes, err := config.pokeApiClient.GetLocationAreas(config.nextLocationsURL)
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

func commandMapBack(param string, config *config) error {
	if config.previousLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locationAreasRes, err := config.pokeApiClient.GetLocationAreas(config.previousLocationsURL)
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

func commandExplore(param string, config *config) error {
	if param == "" {
		return errors.New("need to provide a location area name")
	}

	fmt.Printf("Exploring %s...\n", param)

	result, err := config.pokeApiClient.GetLocationAreaByName(param)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range result.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}


type cliCommand struct {
	name string
	description string
	callback func(string, *config) error
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
		"explore": {
			name: "explore",
			description: "Lists all the Pokémon in a given location area.",
			callback: commandExplore,
		},
	}
}
