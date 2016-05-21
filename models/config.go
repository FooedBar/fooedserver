package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Conf Config

type Config struct {
	DBPath string `json:"dbPath"`
	Port   string `json:"port"`
}

func init() {
	// Get the config file
	config_file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}
	json.Unmarshal(config_file, &Conf)
}
