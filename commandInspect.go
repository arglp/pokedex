package main

import (
	"fmt"
)

func commandInspect(cfg *config, parameter string) error {
	if parameter == "" {
		fmt.Println("Please enter a pokemon Name")
		return nil	
	}
	pokemon, exists := cfg.caughtPokemon[parameter]
	if !exists {
		fmt.Println("This pokemon has not been caught yet")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  -%s\n", typ.Type.Name)
	}

	return nil
}