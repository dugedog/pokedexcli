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
		fmt.Printf("%v was caught!\n", pokemonTarget.Name)
		cfg.pokeapiClient.Pokedex[name] = pokemonTarget
	} else {
		fmt.Printf("%v escaped!\n", pokemonTarget.Name)

	}
	return nil
}
