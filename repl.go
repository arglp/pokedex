package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arglp/pokedex/internal/pokeapi"
	"github.com/arglp/pokedex/internal/pokecache"
)

type config struct{
	pokeapiClient		pokeapi.Client
	pokecacheCache		pokecache.Cache
	nextLocationsURL	*string
	prevLocationsURL	*string
	caughtPokemon		map[string]pokeapi.Poke
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, string) error
}


func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandRegister()
	

	for {
		print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		inputCommand := input[0]
		inputParameter := ""
		if len(input) > 1 {
			inputParameter = input[1]
		}
		command, exists := commands[inputCommand]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(cfg, inputParameter)
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
		"explore": {
			name:			"explore",
			description:	"Takes a name of a location area as an argument and shows a list of all the Pokémon located there",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Catches the Pokémon indicated as the paremeter",
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description: 	"Shows basic information of the caught Pokemon",
			callback:		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description:	"Show a list of the pokemon caught",
			callback:		commandPokedex,
		},
	}
	return commands
}