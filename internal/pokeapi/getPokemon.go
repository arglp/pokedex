package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/arglp/pokedex/internal/pokecache"
	"io"
	"bytes"
)

type Poke struct {
	Id					int			`json:"id"`
	Name				string 		`json:"name"`
	BaseExperience		int			`json:"base_experience"`
	Height				int			`json:"height"`
	Weight				int			`json:"weight"`
	Stats				[]PokeStats		`json:"stats"`		
	Types				[]PokeType	`json:"types"`
}

type PokeStats struct {
	Stat					Stat		`json:"stat"`
	BaseStat				int			`json:"base_stat"`
}

type Stat struct {
	Name					string			`json:"name"`
}

type PokeType struct {
	Type					Type		`json:"type"`
}

type Type struct {
	Name					string		`json:"name"`
}




func GetPokemon (c Client, cache pokecache.Cache, name string) (Poke, error) {
	url := baseUrl + "/pokemon/" + name
	
	poke := Poke{}
	var buffer *bytes.Buffer

	cacheData, exists := cache.Get(url)
	if !exists{
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Poke{}, fmt.Errorf("error making request: %v", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Poke{}, fmt.Errorf("error getting response: %v", err)
		}

		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return Poke{}, fmt.Errorf("error decoding request: %v", err)
		}
		cache.Add(url, data)
		buffer = bytes.NewBuffer(data)
	} else {
		buffer = bytes.NewBuffer(cacheData)
	}
	decoder	:= json.NewDecoder(buffer)
	if err := decoder.Decode(&poke); err != nil {
		return Poke{}, fmt.Errorf("error deocding request: %v", err)

	}
	return poke, nil
}