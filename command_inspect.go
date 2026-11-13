package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, arg ...string) error {
	if len(arg) == 0 {
		return errors.New("Please include the name off the pokemon you are inspecting")
	}
	name := arg[0]
	inspectTarget, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Println("You have not caught that pokeon")
		return nil
	}
	fmt.Printf("Name: %v\n", inspectTarget.Name)
	fmt.Printf("Height: %v\n", inspectTarget.Height)
	fmt.Printf("Weight: %v\n", inspectTarget.Weight)
	fmt.Println("Stats:")
	for _, stats := range inspectTarget.Stats {
		fmt.Printf("-%v: %v\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range inspectTarget.Types {

		fmt.Printf("- %v\n", types.Type.Name)
	}
	return nil
}
