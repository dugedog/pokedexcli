package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func commandMap(cfg *config) error {
	updatedURL := baseURL + "?offset=" + strconv.Itoa(cfg.page) + "&limit=20"
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
	/*
		webBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error reading http request: %v", err)
		}
	*/
	var webJson locationArea

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&webJson); err != nil {
		return err
	}
	for _, location := range webJson.Results {
		fmt.Println(location.Name)
	}

	if cfg.page+20 < webJson.Count {
		cfg.page += 20
	} else {
		fmt.Println("At end of list...")
	}

	return nil
}
