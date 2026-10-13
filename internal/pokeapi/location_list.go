package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (locationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea{}, fmt.Errorf("Error building http request: %v", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)

	if err != nil {
		return locationArea{}, err
	}

	locationResponse := locationArea{}
	err = json.Unmarshal(dat, &locationResponse)
	if err != nil {
		return locationArea{}, err
	}

	return locationResponse, nil
}
