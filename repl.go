package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/sleklere/pokedexcli/internal/pokeapi"
)


type config struct {
	pokeApiClient pokeapi.Client
	nextLocationsURL *string
	previousLocationsURL *string
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()

		cleaned := cleanInput(userInput)

		if len(cleaned) == 0 {
			fmt.Println("Please enter a command")
			continue
		}

		command := cleaned[0]

		supportedCommands := getCommands()

		supportedCommand, ok := supportedCommands[command]
		if ok {
			err := supportedCommand.callback(config, cleaned[1:]...)
			if err != nil {
				fmt.Printf("Error executing command '%s': %v\n", supportedCommand.name, err)
			}
		}
		if !ok {
			fmt.Println("Unknown command")
		}
	}
}
