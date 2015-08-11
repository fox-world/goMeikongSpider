package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Model struct {
	Number  int
	Name    string
	Click   int
	Page    string
	Address []string
}

type Config struct {
	DBuri string
}

func (config *Config) ReadConfig() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/src/config.json")
	if err != nil {
		fmt.Println("Open config file error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		fmt.Println("Decode config file error:", err)
	}
}
