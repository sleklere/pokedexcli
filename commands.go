package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(config *config, args ...string) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

func commandHelp(config *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cliCommand := range getCommands() {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}

	return nil
}

func commandMapForward(config *config, args ...string) error {
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

func commandMapBack(config *config, args ...string) error {
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

func commandExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You need to provide a location area name")
	}

	fmt.Printf("Exploring %s...\n", args[0])

	result, err := config.pokeApiClient.GetLocationAreaByName(args[0])
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
	callback func(*config, ...string) error
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
