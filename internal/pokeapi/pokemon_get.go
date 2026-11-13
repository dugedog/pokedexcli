package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(target string) (PokemonDetails, error) {
	url := baseURL + "/pokemon/" + target
	if val, ok := c.cache.Get(url); ok {
		response := PokemonDetails{}
		err := json.Unmarshal(val, &response)
		if err != nil {
			return PokemonDetails{}, err
		}

		return response, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	pokeResponse := PokemonDetails{}
	err = json.Unmarshal(dat, &pokeResponse)
	if err != nil {
		return PokemonDetails{}, err
	}
	if pokeResponse.Name == "" {
		return PokemonDetails{}, errors.New("Pokemon name not recognised. Please try again")
	}

	c.cache.Add(url, dat)
	return pokeResponse, nil
}
