package main

import (
	"errors"
	"fmt"
	"github.com/germagla/boot-dev-pokedexcli/internal/pokeapi"
	"math/rand"
)

func catchCommand(cfg *config, args ...string) error {
	switch {
	case len(args) == 0:
		return errors.New("no pokemon specified")
	case len(args) > 1:
		return errors.New("too many arguments")

	default:
		//try to get pokemon from cache
		pokemon, ok := cfg.cache.GetPokemon(args[0])
		if ok {
			catchPokemon(cfg, pokemon)
			break
		}

		//fetch pokemon
		pokemon, err := pokeapi.GetPokemon(args[0])
		if err != nil {
			return err
		}
		catchPokemon(cfg, pokemon)

	}
	return nil

}

func catchPokemon(cfg *config, pokemon pokeapi.Pokemon) {
	catchRate := arbitraryFormula(pokemon.BaseExperience)
	n := rand.Intn(500)
	if n < catchRate {
		addPokemontoPokedex(cfg, pokemon)
		fmt.Println(pokemon.Name + " was caught!\nYou may now inspect it with the inspect command.")
		return
	}
	fmt.Println(pokemon.Name + " escaped!")
}

func addPokemontoPokedex(cfg *config, pkmn pokeapi.Pokemon) {
	cfg.pokedex[pkmn.Name] = pkmn

}

func arbitraryFormula(n int) int {
	return 500 - n
}
