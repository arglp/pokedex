package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
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
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
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
	}
	return commands
}