package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (locationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := locationArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return locationArea{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationArea{}, err
	}

	locationsResp := locationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return locationArea{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}

func (c *Client) ListPokemon(area string) (pokeArea, error) {
	url := baseURL + "/location-area/" + area
	if val, ok := c.cache.Get(url); ok {
		pokeReponse := pokeArea{}
		err := json.Unmarshal(val, &pokeReponse)
		if err != nil {
			return pokeArea{}, err
		}

		return pokeReponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokeArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokeArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokeArea{}, err
	}

	pokeResponse := pokeArea{}
	err = json.Unmarshal(dat, &pokeResponse)
	if err != nil {
		return pokeArea{}, err
	}
	if pokeResponse.Name == "" {
		return pokeArea{}, errors.New("Area name not recognised. Please try again")
	}

	c.cache.Add(url, dat)
	return pokeResponse, nil
}
