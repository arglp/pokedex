package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	commands := getCommandRegister()

	for _, c := range commands {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}

	return nil
}