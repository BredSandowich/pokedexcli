package main

import (
	"fmt"
	"net/http"
	"io"
	"errors"
	"encoding/json"

	"github.com/BredSandowich/pokedexcli/internal/pokeapi"
)


func commandExplore(cfg *config, args []string) error {
		if len(args) < 1 {
			return errors.New("location area required")
		}
		area := args[0]

		url := "https://pokeapi.co/api/v2/location-area/" + area

		fmt.Printf("Exploring Pokemon in %s...\n", area)

		var body []byte
		var err error

		//Check if URL is already in cache
		if val, ok := cfg.cache.Get(url); ok{
			body = val
		} else {
			res, err := http.Get(url)
			if err != nil {
				return err
			}
			defer res.Body.Close()		
			
			if res.StatusCode == 404 {
				return errors.New("location area not found")
        	}

        body, err = io.ReadAll(res.Body)
        if err != nil {
            return err
        }

        // Save it to the cache for next time!
        cfg.cache.Add(url, body)
    	}

		// 4. Unmarshal into our corrected struct
		var encounterResp pokeapi.LocationAreaResp
		err = json.Unmarshal(body, &encounterResp)
		if err != nil {
			return err
		}

		// 5. Print the Pokemon found in this area
		fmt.Println("Found Pokemon:")
		for _, encounter := range encounterResp.PokemonEncounters {
			fmt.Printf(" - %s\n", encounter.Pokemon.Name)
		}

		return nil

	}