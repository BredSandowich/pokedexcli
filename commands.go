package main

type cliCommand struct {
	name string
	description string
	callback func() error
}

exitPokedex := map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
}