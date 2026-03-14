package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/viniciusLambert/Pokedex/intnternal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		newScan := scanner.Scan()
		if newScan {
			text := scanner.Text()
			command := cleanInput(text)
			commandFound := false
			for cmd, structure := range commands {
				if cmd == command[0] {
					commandFound = true
					_ = structure.callback()
				}
			}

			if !commandFound {
				fmt.Println("Unknown command")
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			callback:    diplayHelp,
		},
		"map": {
			name:        "map",
			description: "get next 20 maps",
			callback:    getMap,
		},
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	cleanText := strings.Fields(text)
	return cleanText
}
