package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapB(cfg *config) error {
	updatedURL := ""
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		updatedURL = cfg.Previous.(string)
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
