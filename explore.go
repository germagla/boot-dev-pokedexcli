package main

import (
	"errors"
	"fmt"
	"github.com/germagla/boot-dev-pokedexcli/internal/pokeapi"
)

func exploreCommand(c *config, args ...string) error {
	switch {
	case len(args) == 0:
		return errors.New("no location area specified")

	case len(args) > 1:
		return errors.New("too many arguments")

	default:
		//try to get location area from cache
		area, ok := c.cache.GetLocationArea(args[0])
		if ok {
			listPokemon(area)
			break
		}

		//fetch location area
		area, err := pokeapi.GetLocationArea(args[0])
		if err != nil {
			return err
		}
		listPokemon(area)

		//add location area to cache
		err = c.cache.AddLocationArea(args[0], area)
		if err != nil {
			return errors.New("error adding data to cache")
		}

	}
	return nil
}

func listPokemon(area pokeapi.LocationArea) {
	fmt.Println("Found Pokemon:")
	for _, pokemon := range area.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
}
