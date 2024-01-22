package main

import (
	"errors"
	"fmt"
	"github.com/germagla/boot-dev-pokedexcli/internal/pokeapi"
)

func mapfCommand(cfg *config, args ...string) error {

	switch {

	case cfg.next == nil:
		//fetch first page
		firstPage, err := pokeapi.GetLocationAreaList(pokeapi.LocationAreaEndpoint)
		if err != nil {
			return err
		}

		//add it to cache
		err = cfg.cache.AddAreaList(pokeapi.LocationAreaEndpoint, firstPage)
		if err != nil {
			return errors.New("error adding data to cache")
		}
		printResults(firstPage)
		cfg.next = firstPage.Next
		cfg.previous = firstPage.Previous
		break

	default:
		//fetch next page
		nextPage, err := pokeapi.GetLocationAreaList(*cfg.next)
		if err != nil {
			return errors.New("error getting data from PokeAPI")
		}

		//add next page to cache
		err = cfg.cache.AddAreaList(*cfg.next, nextPage)
		if err != nil {
			return errors.New("error adding data to cache")
		}

		printResults(nextPage)
		cfg.next = nextPage.Next
		cfg.previous = nextPage.Previous
	}

	return nil
}

func mapbCommand(cfg *config, args ...string) error {

	switch {

	case cfg.previous == nil:
		return errors.New("on First Page")

	default:
		//try to get previous page from cache
		previousPage, ok := cfg.cache.GetAreaList(*cfg.previous)
		if ok {
			printResults(previousPage)
			cfg.next = previousPage.Next
			cfg.previous = previousPage.Previous
			break
		}

		//fetch previous page
		prevPage, err := pokeapi.GetLocationAreaList(*cfg.previous)
		if err != nil {
			return err
		}
		printResults(prevPage)
		cfg.next = prevPage.Next
		cfg.previous = prevPage.Previous
	}

	return nil
}

func printResults(list pokeapi.LocationAreaList) {
	fmt.Println("Found Location Areas:")
	for _, area := range list.Results {
		fmt.Println(area.Name)
	}

}
