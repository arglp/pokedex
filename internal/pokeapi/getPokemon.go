package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/arglp/pokedex/internal/pokecache"
	"io"
	"bytes"
)



type LocationArea struct {
	Id					int					`json:"id"`
	Name				string 				`json:"name"`
	PokemonEncounters	[]PokemonEncounter 	`json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon 				Pokemon				`json:"pokemon"`
	
}

type Pokemon struct {
	Name				string				`json:"name"`
	Url					string				`json:"url"`

}

func GetPokemon (c Client, cache pokecache.Cache, area string) (LocationArea, error) {
	url := baseUrl + "/location-area/" + area 
	
	locationArea := LocationArea{}
	var buffer *bytes.Buffer

	cacheData, exists := cache.Get(url)
	if !exists{
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error making request: %v", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error getting response: %v", err)
		}

		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error decoding request: %v", err)
		}
		cache.Add(url, data)
		buffer = bytes.NewBuffer(data)
	} else {
		buffer = bytes.NewBuffer(cacheData)
	}
	decoder	:= json.NewDecoder(buffer)
	if err := decoder.Decode(&locationArea); err != nil {
		return LocationArea{}, fmt.Errorf("error deocding request: %v", err)

	}
	return locationArea, nil
}