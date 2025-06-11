package main

import (
	"fmt"
	"github.com/arglp/pokedex/internal/pokeapi"
	"math/rand"
)

func commandCatch(cfg *config, parameter string) error {
	if parameter == "" {
		fmt.Println("Please enter pokémon name as a parameter")
	} else {
		fmt.Printf("Throwing a Pokeball at %v...\n", parameter)
		pokemon, err := pokeapi.GetPokemon(cfg.pokeapiClient, cfg.pokecacheCache, parameter)

		if err != nil {
			return err
		}
		chance := rand.Intn(700)
		if chance >= pokemon.BaseExperience{
			fmt.Printf("%s was caught!\n", pokemon.Name)
			cfg.caughtPokemon[pokemon.Name] = pokemon
		} else {
			fmt.Printf("%s escaped!\n", pokemon.Name)
		}
	}	
	return nil
}