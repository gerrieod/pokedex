package main

import (
	"fmt"
	"os"
)



func getCommands() map[string]clicommand {
	return map[string]clicommand{
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
		},
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
    "map":{
      name: "map",
      description: "Get next pokemon map locations",
      callback: commandMap,
    },
    "mapb":{
      name: "mapb",
      description: "Get previous pokemon map locations",
      callback: commandMapb,
    },
	}
}


//command to display help information : can modify to display help for a command instead of just one command 
func commandHelp() error{
	fmt.Println("This is the Pokedex help command")
	fmt.Println("Available commands: help | exit")
	return nil
}

//stop pokedex execution
func commandExit() error{
	fmt.Println("Pokedex closing....")
	os.Exit(0)
	return nil
}

func commandMap() error{
  return nil
}

func commandMapb() error{
  return nil
}
