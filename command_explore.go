package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.argument) == 0 {
		return errors.New("Please provide area name when exploring eg 'explore <area name>'")
	}

	pokemonList, err := cfg.pokeapiClient.ListPokemon(cfg.argument[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", cfg.argument[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonList.PokemonEncounters {

		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
