package main

import (
	"bufio"
	"fmt"
	"github.com/dugdeog/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		var arg []string
		if len(cleanedInput) > 1 {
			arg = cleanedInput[1:]

		}

		menu, ok := getCommands()[input]

		if !ok {
			fmt.Println("Unknown command")
		} else {
			if err := menu.callback(config, arg...); err != nil {
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
	caughtPokemon map[string]pokeapi.PokemonDetails
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
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
		"inspect": {
			name:        "inspect",
			description: "Displays informaion about a pokemon if you have caught it",
			callback:    commandInspect,
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
		"explore": {
			name:        "explore",
			description: "Look around and stuff",
			callback:    commandExplore,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List Pokemon from your pokedex",
			callback:    commandPokedex,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"list": {
			name:        "list",
			description: "List Pokemone in pokedex",
			callback:    commandList,
		},
	}

}
