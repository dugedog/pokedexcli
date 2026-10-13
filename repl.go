package main

import (
	"bufio"
	"fmt"
	"github.com/dugdeog/pokedexcli/internal/pokeapi"
	"github.com/dugdeog/pokedexcli/internal/pokecache"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(config *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if len(scanner.Text()) == 0 {
			continue
		}

		cleanedInput := cleanInput(scanner.Text())
		input := cleanedInput[0]
		menu, ok := getCommands()[input]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			if err := menu.callback(config); err != nil {
				fmt.Println(err)
			}

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

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	pCache        *pokecache.Cache
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Obtain a list of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Move backwards in a list of location areas",
			callback:    commandMapB,
		},
	}

}
