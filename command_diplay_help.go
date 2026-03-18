package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	commands := getCommands()
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for name, cmd := range commands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}
