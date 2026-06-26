package main

import (
	"fmt"
	"bufio"
	"os"
	"time"

	"github.com/BredSandowich/pokedexcli/internal/pokecache"
	"github.com/BredSandowich/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{
		caughtPokemon : make(map[string]pokeapi.PokemonResp),
	}
	cfg.cache = *pokecache.NewCache(6 * time.Second)
	commands := getCommands(cfg)

	for ;; {
		fmt.Print("Pokedex > ")
		//Wait for cli prompt to scan and then return the string
		scanner.Scan()
		entry := scanner.Text()

		input := cleanInput(entry)
		if len(input) == 0 {
			continue
		}
		
		command := input[0]
		args := input[1:]

		cmd, ok := commands[command]
		
		if ok {
			err := cmd.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}