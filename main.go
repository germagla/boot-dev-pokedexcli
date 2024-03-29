package main

import (
	"bufio"
	"fmt"
	"github.com/germagla/boot-dev-pokedexcli/internal/cache"
	"github.com/germagla/boot-dev-pokedexcli/internal/pokeapi"
	"os"
	"strings"
	"time"
)

func main() {
	var conf config
	initConfig(&conf)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex CLI. Enter 'help' for a list of commands.")
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command, args := parseInput(scanner.Text())
		function, ok := funcMap[command]
		if !ok {
			fmt.Println("Unknown command.")
			continue
		}
		err := function(&conf, args...)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
	}
}

func initConfig(c *config) {
	c.cache = cache.NewCache(5 * time.Second)
	c.pokedex = make(map[string]pokeapi.Pokemon)
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var funcMap = map[string]func(*config, ...string) error{
	"help":    helpCommand,
	"exit":    exitCommand,
	"map":     mapfCommand,
	"mapb":    mapbCommand,
	"config":  configCommand,
	"explore": exploreCommand,
	"catch":   catchCommand,
	"inspect": inspectCommand,
	"pokedex": pokedexCommand,
}

func configCommand(c *config, args ...string) error {
	fmt.Println(c)
	return nil
}

func parseInput(input string) (command string, args []string) {
	words := strings.Split(input, " ")
	command = words[0]
	args = words[1:]
	return command, args
}

type config struct {
	next     *string
	previous *string
	cache    cache.Cache
	pokedex  map[string]pokeapi.Pokemon
}
