package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port              int
	GoogleAPIKey      string
	EncodedTwitterKey string
}

func ReadFromFile(path string) (Configuration, error) {
	configuration := Configuration{}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return configuration, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	return configuration, err
}
