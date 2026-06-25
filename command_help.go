package main

import (
	"fmt"
)


func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for key, value := range getCommands(cfg) {
		fmt.Printf("\t%s: %s\n", key, value.description)
	}

	return nil
}