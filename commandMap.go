package main

import (
	"fmt"
	"github.com/arglp/pokedex/internal/pokeapi"
)

func commandMapf(cfg *config, parameter string) error {
	url := ""
	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	} 
	MapList, err := pokeapi.GetAreas(cfg.pokeapiClient, cfg.pokecacheCache, url)
	if err != nil {
		return err
	}
	locations := MapList.Locations
	for _, location := range locations {
		fmt.Printf("%v\n", location.Name)
	}

	cfg.nextLocationsURL = &MapList.NextURL
	cfg.prevLocationsURL = &MapList.PreviousURL

	return nil
}

func commandMapb(cfg *config, parameter string) error {
	if cfg.prevLocationsURL == nil || *cfg.prevLocationsURL == ""  {
		fmt.Println("you're on the first page")
	} else {
		MapList, err := pokeapi.GetAreas(cfg.pokeapiClient, cfg.pokecacheCache, *cfg.prevLocationsURL)
		if err != nil {
			return err
		}
		locations := MapList.Locations
		for _, location := range locations {
			fmt.Printf("%v\n", location.Name)
		}

		cfg.nextLocationsURL = &MapList.NextURL
		cfg.prevLocationsURL = &MapList.PreviousURL
	}
	return nil
}