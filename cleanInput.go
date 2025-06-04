package main

import(
	"strings"
)

func cleanInput (text string) []string {
	words := strings.Fields(text)
	cleanedWords := []string{}
	for _, w := range words {
		cleanedWords = append(cleanedWords, strings.TrimSpace(strings.ToLower(w)))
	}

	return cleanedWords
}