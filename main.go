package main

import (
	"github.com/dugdeog/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{commands: getCommands(),
		caughtPokemon: map[string]pokeapi.PokemonDetails{},
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
