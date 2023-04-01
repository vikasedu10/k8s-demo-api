package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Role string
}

var Config Configuration

func ConfigurationServer(serverMode string) {
	if serverMode == "dev" {
		os.Setenv("KEY", "##")
		os.Setenv("VALUE", "##")
	}
	file, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}
