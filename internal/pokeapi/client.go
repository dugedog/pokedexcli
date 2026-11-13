package pokeapi

import (
	"github.com/dugdeog/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

// Client
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	Pokedex    map[string]PokemonDetails
}

// NewClent
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache:      pokecache.NewCache(cacheInterval),
		httpClient: http.Client{Timeout: timeout},
		Pokedex:    make(map[string]PokemonDetails),
	}
}
