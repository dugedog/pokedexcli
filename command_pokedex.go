package main

import (
	"fmt"
)

func commandPokedex(cfg *config, arg ...string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You do not have any caught pokemon.")
	}
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %v\n", pokemon.Name)
	}

	return nil
}
