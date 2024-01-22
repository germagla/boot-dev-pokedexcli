package main

import "fmt"

func helpCommand(c *config, args ...string) error {
	fmt.Println("Available commands:")
	for command, helpMessage := range helpMap {
		fmt.Printf("\t%v: %v\n", command, helpMessage)
	}
	return nil
}

var helpMap = map[string]string{
	"help": "Prints this help message",
	"exit": "Exits the program",
	"map":  "Displays the names of 20 location areas. Successive calls will display the next 20 location areas.",
	"mapb": "Displays the names of the previous 20 location areas." + " " +
		"This command will print an error if you try to go back before the first 20 location areas.",
	"explore [area]":  "Displays the names of the pokemon found in the current area. ",
	"catch [pokemon]": "Attempts to catch the specified pokemon. If successful, the pokemon will be added to your pokedex.",
	"inspect [name]":  "Displays the stats of the specified pokemon.",
	"pokedex":         "Displays the names of all pokemon in your pokedex.",
}
