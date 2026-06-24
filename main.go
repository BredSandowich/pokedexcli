package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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

		commands :=getCommands()
		cmd, ok := commands[firstWord]
		
		if ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}