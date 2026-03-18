package main

import (
	"fmt"
	"log"
)

func commandExploreArea(cfg *config, args ...string) error {
	fmt.Printf("Exploring %s...", args[0])

	locationsArea, err := cfg.pokeapiClient.GetLocationAreaData(&args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationsArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
