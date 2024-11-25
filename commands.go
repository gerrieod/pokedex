package main

import (
	"fmt"
	"os"
)

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
