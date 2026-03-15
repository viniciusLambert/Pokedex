package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/viniciusLambert/Pokedex/internal/pokeapi"
	"github.com/viniciusLambert/Pokedex/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokecacheCache   *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}

			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Call help command",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "get next 20 maps",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previuous 20 maps",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	cleanText := strings.Fields(text)
	return cleanText
}
