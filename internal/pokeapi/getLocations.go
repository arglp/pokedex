package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
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

func GetAreas (c Client, url string) (AreaLocationList, error) {
	fullUrl := ""
	if url == "" {
		fullUrl = baseUrl + "/location-area"
	} else {
		fullUrl = url
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Errorf("Error making Request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt. Errorf("Error getting response: %v", err)
	}

	defer res.Body.Close()
	var areaLocationList AreaLocationList

	decoder	:= json.NewDecoder(res.Body)
	if err := decoder.Decode(&areaLocationList); err != nil {
		fmt.Errorf("Error decoding: %v", err)
	}

	return areaLocationList, nil
}