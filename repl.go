package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arglp/pokedex/internal/pokeapi"
)

type config struct{
	pokeapiClient		pokeapi.Client
	nextLocationsURL	*string
	prevLocationsURL	*string
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandRegister()

	for {
		print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())[0]
		
		command, exists := commands[input]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func getCommandRegister() map[string]cliCommand{
	commands := map[string]cliCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"map": {
			name:			"map",
			description:    "Get the next page of locations",
			callback:		commandMapf,
		},
		"mapb": {
			name:			"mapb",
			description:    "Get the previous page of locations",
			callback:		commandMapb,
		},
	}
	return commands
}