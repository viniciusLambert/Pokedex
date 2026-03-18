package main

import (
	"fmt"
)

func commandDisplayPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You have no pokemons yet.")
		return nil
	}
	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  -%s\n", pokemon.Name)
	}

	return nil
}
