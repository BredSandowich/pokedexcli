package main

import (
	"fmt"
	"errors"

)

func commandPokedex(cfg *config, args []string) error {
		caughtPokemon := cfg.caughtPokemon
		if len(caughtPokemon) < 1 {
			return errors.New("You don't have any pokemon, go catch them all!")
		}
		
		fmt.Println("Your Pokedex:")
		
		for _, caughtPokemon := range caughtPokemon{
			fmt.Printf("   -%s\n", caughtPokemon.Name)
		}

		
		return nil

	}