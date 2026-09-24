package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var register map[string]cliCommand

func init() {
	register = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if len(scanner.Text()) == 0 {
			continue
		}

		cleanedInput := cleanInput(scanner.Text())
		input := cleanedInput[0]
		command, ok := register[input]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.callback()

		}
	}

}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	lowered := strings.ToLower(text)
	stringList := strings.Fields(lowered)
	return stringList
}

func commandExit() error {
	fmt.Println("Closing the Pokedex...Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range register {
		//fmt.Printf("%s: %s", command.name, command.description)
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
