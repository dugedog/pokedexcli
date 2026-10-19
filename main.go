package main

import (
	"github.com/dugdeog/pokedexcli/internal/pokeapi"
	"github.com/dugdeog/pokedexcli/internal/pokecache"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	newCache := pokecache.NewCache(60 * time.Second)
	cfg := &config{commands: getCommands(),
		pCache:        newCache,
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
