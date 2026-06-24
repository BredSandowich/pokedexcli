package main

import (
	"fmt"
)


func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for key, value := range getCommands() {
		fmt.Printf("\t%s: %s\n", key, value.description)
	}

	return nil
}