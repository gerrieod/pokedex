package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//create a struct for commands
type clicommand struct{
	name string
	description string
	callback func() error
}

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
	}
}

func main(){
	//requirements
	//create a way for the user to interact with the program
	//get user input with buffio.NewScanner function 	https://pkg.go.dev/bufio#NewScanner. and the Scan function
	scanner := bufio.NewScanner(os.Stdin)

	//get the available commands of the program
	commands := getCommands();
	//need infinite loop to keep the program running
	for {
		//promt the user for input.
		fmt.Print("pokedex >")
		scanner.Scan()
		text := scanner.Text()
		//sanitize user input
		words := sanitizeInput(text)
		command, ok := commands[words[0]]
		if !ok{
			fmt.Println("Invalid command")
			continue
		}
		command.callback()
	}
}

//extend this function to refine user input and also check for malicious inputs
func sanitizeInput(input string)[]string{
	lowerCase := strings.ToLower(input)
	words := strings.Fields(lowerCase)
	return words
}



