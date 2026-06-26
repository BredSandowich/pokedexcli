package main

import (
	"fmt"
	"net/http"
	"io"
	"errors"
	"encoding/json"
	"math/rand/v2"

	"github.com/BredSandowich/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
		//Check for arguments ie) catch + pokemon name
		if len(args) < 1 {
			return errors.New("pokemon name required")
		}
		pokemon := args[0]

		url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

		fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

		var body []byte
		var err error

		//Check if URL is already in cache for faster load
		if val, ok := cfg.cache.Get(url); ok{
			body = val
		} else {
			res, err := http.Get(url)
			if err != nil {
				return err
			}
			defer res.Body.Close()		
			
			if res.StatusCode == 404 {
				return errors.New("pokemon too rare to exist")
        	}

        body, err = io.ReadAll(res.Body)
        if err != nil {
            return err
        }

        // Save it to the cache for next time!
        cfg.cache.Add(url, body)
    	}

		// 4. Unmarshal into our corrected struct
		var pokemonResp pokeapi.PokemonResp
		err = json.Unmarshal(body, &pokemonResp)
		if err != nil {
			return err
		}

		// 5. Randomize if pokemon was caught
		catchRandomizer := rand.IntN(pokemonResp.BaseExperience)
		if  15 > catchRandomizer {
			fmt.Println("Pokemon got away!")
			return nil
		} else {
			cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
			fmt.Printf("Caught %s\n", pokemonResp.Name)
			return nil
		}


		return nil

	}