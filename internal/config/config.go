package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}
	}

	b, err := os.ReadFile(home + "/" + configFileName)
	if err != nil {
		log.Fatal(err)
	}

	// `b` contains everything your file has.
	// This writes it to the Standard Out.
	// os.Stdout.Write(b)
	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
	// You can also write it to a file as a whole.

}

func (config Config) SetUser(current_user_name string) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configInstance := &Config{DbURL: "postgres://example", CurrentUserName: current_user_name}
	b, err := json.Marshal(configInstance)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(home+"/"+configFileName, b, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

}
