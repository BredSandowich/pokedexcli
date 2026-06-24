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
		firstWord := input[0]
		returnText := fmt.Sprintf("Your command was: %s", firstWord)
		fmt.Println(returnText)

	}
}