package main

import (
	"fmt"
	"io"
	"strings"
	"github.com/sleklere/pokedexcli/internal/pokeapi"

	"github.com/chzyer/readline"
)


type config struct {
	pokeApiClient pokeapi.Client
	nextLocationsURL *string
	previousLocationsURL *string
	catchedPokemons map[string]pokeapi.Pokemon
}

func startRepl(config *config) {
	supportedCommands := getCommands()

	autoCompleteItems := make([]readline.PrefixCompleterInterface, 0)

	for _, v := range supportedCommands {
		autoCompleteItems = append(autoCompleteItems, readline.PcItem(v.name))
	}


	completer := readline.NewPrefixCompleter(
		autoCompleteItems...
	)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "Pokedex > ",
		HistoryFile:     "/tmp/pokedex_history",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		userInput := strings.TrimSpace(line)
		if userInput == "exit" {
			break
		}

		cleaned := cleanInput(userInput)
		if len(cleaned) == 0 {
			fmt.Println("Please enter a command")
			continue
		}

		command := cleaned[0]
		supportedCommand, ok := supportedCommands[command]
		if ok {
			err := supportedCommand.callback(config, cleaned[1:]...)
			if err != nil {
				fmt.Printf("Error executing command '%s': %v\n", supportedCommand.name, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
