package main

import (
	"github.com/dugdeog/pokedexcli/internal/pokeapi"
	"github.com/dugdeog/pokedexcli/internal/pokecache"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{commands: getCommands(),
		pokeapiClient: pokeClient,
	}
	mainPokecache := pokecache.NewCache(5)
	startRepl(cfg)
}
