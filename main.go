package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{}
	cfg.cache := *pokecache.NewCache(6 * time.Seconds)
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
		firstWord := input[0]

		cmd, ok := commands[firstWord]
		
		if ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}