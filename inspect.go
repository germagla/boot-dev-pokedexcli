package main

import (
	"errors"
	"fmt"
	"github.com/germagla/boot-dev-pokedexcli/internal/pokeapi"
)

func inspectCommand(c *config, args ...string) error {
	switch len(args) {
	case 0:
		return errors.New("no pokemon specified")
	case 1:
		pokemon, ok := c.pokedex[args[0]]
		if !ok {
			return errors.New("you have to catch " + args[0] + " first")
		}
		printStats(pokemon)

	default:
		return errors.New("too many arguments")
	}

	return nil
}

func printStats(pokemon pokeapi.Pokemon) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%v\n", t.Type.Name)
	}
}
