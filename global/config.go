package global

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Config struct {
	Endpoint string `json:"endpoint"`
}

var GlobalConfig *Config

func Initialize() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Error opening config file. Please make sure it exists.")
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	GlobalConfig = &Config{}
	err = decoder.Decode(GlobalConfig)

	if err != nil {
		log.Fatal("Error formatting config file. Please make sure it is in JSON format.")
		panic(err)
	}
	if GlobalConfig.Endpoint == "" {
		log.Fatal("Error in config file. Both 'endpoint'field must be present and non-empty.")
		panic(errors.New("Invalid config file"))
	}
}
