package utils

import (
	"encoding/json"
	"log"
	"os"
)

func ReadEndpointList() (map[string]string, error) {
	jsonFile, err := os.Open("src/endpointList.json")
	if err != nil {
		log.Println("Error reading the endpoint list file", err)
		return nil, err
	}
	defer jsonFile.Close()

	var endpoints map[string]string
	err = json.NewDecoder(jsonFile).Decode(&endpoints)
	if err != nil {
		log.Println("Error reading the endpoint list into a map", err)
		return nil, err
	}
	return endpoints, nil
}
