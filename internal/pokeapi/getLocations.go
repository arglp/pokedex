package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/arglp/pokedex/internal/pokecache"
	"io"
	"bytes"
)



type AreaLocationList struct {
	NextURL	string `json:"next"`
	PreviousURL	string `json:"previous"`
	Locations	[]Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

func GetAreas (c Client, cache pokecache.Cache, url string) (AreaLocationList, error) {
	fullUrl := ""
	if url == "" {
		fullUrl = baseUrl + "/location-area"
	} else {
		fullUrl = url
	}

	areaLocationList := AreaLocationList{}
	var buffer *bytes.Buffer

	cacheData, exists := cache.Get(fullUrl)
	if !exists{
		req, err := http.NewRequest("GET", fullUrl, nil)
		if err != nil {
			return AreaLocationList{}, fmt.Errorf("error making request: %v", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return AreaLocationList{}, fmt.Errorf("error getting response: %v", err)
		}

		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return AreaLocationList{}, fmt.Errorf("error decoding request: %v", err)
		}
		cache.Add(fullUrl, data)
		buffer = bytes.NewBuffer(data)
	} else {
		buffer = bytes.NewBuffer(cacheData)
	}
	decoder	:= json.NewDecoder(buffer)
	if err := decoder.Decode(&areaLocationList); err != nil {
		return AreaLocationList{}, fmt.Errorf("error deocding request: %v", err)

	}
	return areaLocationList, nil
}