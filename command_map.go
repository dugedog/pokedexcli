package main

import (
	"encoding/json"
	"fmt"
	"github.com/dugdeog/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	queryResult := pokeapi.NewLocationAreaStruct()
	var err error
	if cfg.Next != nil {
		cacheCheck, cacheOK := cfg.pCache.Get(*cfg.Next)
		if cacheOK {
			queryResult = pokeapi.NewLocationAreaStruct()
			json.Unmarshal(cacheCheck, &queryResult)
		}
	}

	if queryResult.Results == nil {
		queryResult, err = cfg.pokeapiClient.ListLocations(cfg.Next)
		if err != nil {
			return fmt.Errorf("Error trying to get locations: %v", err)
		}
		// Marshalling json data to add to cache map.
		marshelledData, err := json.Marshal(queryResult)
		if err != nil {
			return fmt.Errorf("Error trying to convert result to json/bytes: %v", err)
		}
		// Adding marshelled data to caache map
		if cfg.Next != nil {
			cfg.pCache.Add(*cfg.Next, marshelledData)
		}

	}

	cfg.Next = queryResult.Next
	cfg.Previous = queryResult.Previous

	for _, location := range queryResult.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	queryResult := pokeapi.NewLocationAreaStruct()
	var err error
	if cfg.Previous == nil {
		return fmt.Errorf("you're on the first page")

	}
	cacheCheck, cacheOK := cfg.pCache.Get(*cfg.Previous)
	if cacheOK {
		queryResult = pokeapi.NewLocationAreaStruct()
		json.Unmarshal(cacheCheck, &queryResult)
	}

	if queryResult.Results == nil {
		queryResult, err = cfg.pokeapiClient.ListLocations(cfg.Previous)
		if err != nil {
			return fmt.Errorf("Error trying to get locations: %v", err)
		}
		// Marshalling json data to add to cache map.
		marshelledData, err := json.Marshal(queryResult)
		if err != nil {
			return fmt.Errorf("Error trying to convert result to json/bytes: %v", err)
		}
		// Adding marshelled data to caache map
		if cfg.Next != nil {
			cfg.pCache.Add(*cfg.Previous, marshelledData)
		}

	}

	cfg.Next = queryResult.Next
	cfg.Previous = queryResult.Previous
	for _, location := range queryResult.Results {
		fmt.Println(location.Name)
	}

	return nil
}
