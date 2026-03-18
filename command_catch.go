package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

func commandCatchPokemon(cfg *config, args ...string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := cfg.pokeapiClient.GetPokemonData(&args[0])
	if err != nil {
		switch err.Error() {
		case "pokemon not found":
			fmt.Println("Pokemon does not exits...")
			return nil
		default:
			log.Fatal(err)
		}
	}

	captureTreshHold := pokemon.BaseExperience/5 + 20
	pokeballPower := rand.IntN(101)

	if pokeballPower > captureTreshHold {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.pokedex[args[0]] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", args[0])
	}
	return nil
}
