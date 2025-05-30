package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			fmt.Println("An error occured")
		}
		userInput := scanner.Text()
		cleanedUserInput := cleanInput(userInput)
		fmt.Printf("Your command was: %s\n", cleanedUserInput[0])
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return words
}