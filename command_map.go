package main

import (
	"fmt"
	"log"
)

func commandMapf(cfg *config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, cfg.pokecacheCache)
	if err != nil {
		log.Fatal(err)
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL, cfg.pokecacheCache)
	if err != nil {
		log.Fatal(err)
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
