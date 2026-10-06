package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	listLocations, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return fmt.Errorf("Error trying to get locations: %v", err)
	}

	cfg.Next = listLocations.Next
	cfg.Previous = listLocations.Previous
	for _, location := range listLocations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	fmt.Println(main
	if cfg.Previous == nil {
		return fmt.Errorf("you're on the first page")

	}
	listLocations, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return fmt.Errorf("Error trying to get locations: %v", err)
	}

	cfg.Next = listLocations.Next
	cfg.Previous = listLocations.Previous
	for _, location := range listLocations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
