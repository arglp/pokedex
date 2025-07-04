package main

import (
	"time"
	"github.com/arglp/pokedex/internal/pokeapi"
	"github.com/arglp/pokedex/internal/pokecache"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(1 * time.Hour)

	cfg := &config{
		pokeapiClient: 		pokeClient,
		pokecacheCache:		pokeCache,
		caughtPokemon:		make(map[string]pokeapi.Poke),
	}

	startRepl(cfg)
}