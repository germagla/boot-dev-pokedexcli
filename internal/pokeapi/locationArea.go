package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int           `json:"min_level"`
				MaxLevel        int           `json:"max_level"`
				ConditionValues []interface{} `json:"condition_values"`
				Chance          int           `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					Url  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationArea(name string) (LocationArea, error) {
	var locationAreaDetails LocationArea
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	url := baseURL + name
	resp, err := http.Get(url)
	if err != nil {
		return locationAreaDetails, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaDetails, err
	}
	err = json.Unmarshal(body, &locationAreaDetails)
	if err != nil {
		return locationAreaDetails, err
	}
	return locationAreaDetails, nil
}
