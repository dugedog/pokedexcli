package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	updatedURL := ""
	if cfg.Next == "" {
		updatedURL = baseURL
	} else {
		updatedURL = cfg.Next
	}
	req, err := http.NewRequest("GET", updatedURL, nil)
	if err != nil {
		return fmt.Errorf("Error building http request: %v", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error executin http request: %v", err)
	}
	defer res.Body.Close()

	var webJson locationArea

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&webJson); err != nil {
		return err
	}
	cfg.Next = webJson.Next
	cfg.Previous = webJson.Previous
	for _, location := range webJson.Results {
		fmt.Println(location.Name)
	}

	return nil
}
