package main

import "fmt"

func commandInspectPokemon(cfg *config, args ...string) error {
	pokemon, exist := cfg.pokedex[args[0]]
	if !exist {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Status:")
	for _, pokemon := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", pokemon.Stat.Name, pokemon.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemon := range pokemon.Types {
		fmt.Printf("  -%s\n", pokemon.Type.Name)
	}

	return nil
}
