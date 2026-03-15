package main

import (
	"time"

	"github.com/viniciusLambert/Pokedex/internal/pokeapi"
	"github.com/viniciusLambert/Pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient:  pokeClient,
		pokecacheCache: cache,
	}
	startRepl(cfg)
}
