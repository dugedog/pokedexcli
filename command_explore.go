package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, arg ...string) error {
	if len(arg) == 0 {
		return errors.New("Please provide area name when exploring eg 'explore <area name>'")
	}

	pokemonList, err := cfg.pokeapiClient.GetLocation(arg[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", arg[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonList.PokemonEncounters {

		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
