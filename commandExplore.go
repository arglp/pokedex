package main

import (
	"fmt"
	"github.com/arglp/pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, parameter string) error {
	if parameter == "" {
		fmt.Println("Please enter location name as a parameter")
	} else {
		fmt.Printf("Exploring %v...\n", parameter)
		area, err := pokeapi.GetPokemon(cfg.pokeapiClient, cfg.pokecacheCache, parameter)
		if err != nil {
			return err
		}
		encounterList := area.PokemonEncounters
		for _, encounter := range encounterList {
			fmt.Printf("%v\n", encounter.Pokemon.Name)
		}
	}	
	return nil
}