package config

import (
	"encoding/json"
	"log"
	"nkn-web/src/helpers"
	"os"
)

const (
	Version string = "0.0.1"
)

var (
	ConfigPath = "./config.json"
	Parameters = &Configuration{
		Version:    "0.0.0",
		ApiBaseUri: "/api",
		Port: "8080",
	}
)

type Configuration struct {
	Version    string `json:"version"`
	ApiBaseUri string `json:"apiBaseUri"`
	Port string `json:"port"`
}

func Init() error {
	if _, err := os.Stat(ConfigPath); err == nil {
		file, _ := helpers.LoadFile(ConfigPath)

		err = json.Unmarshal(file, Parameters)
		if err != nil {
			return err
		}
	} else {
		log.Println("Config file not exists, use default parameters.")
	}

	return nil
}
