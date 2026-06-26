package main

type cliCommand struct {
	name string
	description string
	callback func(*config, []string) error
}


func getCommands(cfg *config) map[string]cliCommand {
	return map[string]cliCommand {
	"map": {
			name: "map",
			description: "List of location end-points",
			callback: getMap,
	},
	
	"mapb": {
		name: "map",
		description: "List of location end-points",
		callback: getMapBack,
},
	
	"help": {
		name: "help",
		description: "Help index of the Pokedex",
		callback: commandHelp,
	},
	
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},

	"explore": {
		name: "explore",
		description: "list all pokemon located in selected location",
		callback: commandExplore,
	},

	"catch": {
		name: "catch",
		description: "add a pokemon to your catalogue",
		callback: commandCatch,
	},

	"inspect": {
		name: "inspect",
		description: "investigate catalogue of pokemon you have caught",
		callback: commandInspect,
	},

}
}