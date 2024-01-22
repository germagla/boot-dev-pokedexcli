package main

import (
	"fmt"
	"os"
)

func exitCommand(c *config, args ...string) error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil
}
