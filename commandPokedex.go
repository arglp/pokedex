package main

import (
	"fmt"
)

func commandPokedex(cfg *config, parameter string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any pokemon yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, poke := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", poke.Name)
	}
	return nil
}