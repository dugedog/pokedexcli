package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocation(area string) (pokeArea, error) {
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
