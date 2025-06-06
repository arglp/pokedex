package main

import (
	"time"
	"github.com/arglp/pokedex/internal/pokeapi"
	"github.com/arglp/pokedex/internal/pokecache"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)

	cfg := &config{
		pokeapiClient: 		pokeClient,
		pokecacheCache:		pokeCache,
	}

	startRepl(cfg)
}