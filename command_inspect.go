package main

import (
	"fmt"
	"errors"

)

func commandInspect(cfg *config, args []string) error {
		//Check for arguments ie) catch + pokemon name
		if len(args) < 1 {
			return errors.New("pokemon name required")
		}
		pokemonname := args[0]

		pokemon, ok := cfg.caughtPokemon[pokemonname]
		if !ok {
			fmt.Println("You have not caught that pokemon")
			return nil
		}
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.PokemonStats{
			fmt.Printf("   -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, ptype := range pokemon.PokemonTypes{
			fmt.Printf("   -%s\n", ptype.Type.Name)
		}

		return nil

	}