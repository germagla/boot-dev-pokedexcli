package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const LocationAreaEndpoint = "https://pokeapi.co/api/v2/location-area/"

type LocationAreaList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreaList(url string) (LocationAreaList, error) {
	var locationAreaList LocationAreaList
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaList{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaList{}, err
	}
	merr := json.Unmarshal(body, &locationAreaList)
	if merr != nil {
		return LocationAreaList{}, err
	}
	err = resp.Body.Close()
	if err != nil {
		return LocationAreaList{}, err
	}

	return locationAreaList, nil
}
