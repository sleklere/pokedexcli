package main

import (
	"bufio"
	"fmt"
	"os"
	cache "github.com/sleklere/pokedexcli/internal/pokecache"
)


type config struct {
	nextLocationsURL *string
	previousLocationsURL *string
}

func startRepl(config *config, cache *cache.Cache) {
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
			err := supportedCommand.callback(config, cache)
			if err != nil {
				fmt.Printf("Error executing command '%s': %v", supportedCommand.name, err)
			}
		}
		if !ok {
			fmt.Println("Unknown command")
		}
	}
}
