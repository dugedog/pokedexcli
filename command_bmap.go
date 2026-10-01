package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func commandMapB(cfg *config) error {
	if cfg.page == 0 {
		fmt.Println("You're on the first page")
		return nil
	} else {
		cfg.page -= 20
	}
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

	var webJson locationArea

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&webJson); err != nil {
		return err
	}
	for _, location := range webJson.Results {
		fmt.Println(location.Name)
	}

	return nil
}
