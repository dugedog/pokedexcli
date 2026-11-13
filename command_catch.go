package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arg ...string) error {
	name := arg[0]
	if len(arg) == 0 {
		return errors.New("Please provide the name of the pokemon you are trying to catch")
	}
	pokemonTarget, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	catchDifficulty := (pokemonTarget.BaseExperience / 10) + 100
	catchAttempt := rand.Intn(catchDifficulty)
	if catchAttempt <= 40 {
		cfg.caughtPokemon[name] = pokemonTarget
		fmt.Printf("%v was caught!\n", pokemonTarget.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%v escaped!\n", pokemonTarget.Name)

	}
	return nil
}
