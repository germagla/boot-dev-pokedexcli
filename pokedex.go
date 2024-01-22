package main

import "fmt"

func pokedexCommand(c *config, args ...string) error {
	switch len(args) {
	case 0:
		fmt.Println("Your pokedex:")
		for name := range c.pokedex {
			fmt.Println(name)
		}
	default:
		fmt.Println("This command takes no arguments.")

	}
	return nil
}
