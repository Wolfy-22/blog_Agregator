package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
)

type Config struct {
	Db_url          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const (
	configFileName = ".gatorconfig.json"
)

func Read() Config {
	configPath, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("error retrieving home directory: %v", err)
	}

	file, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return Config{}
	}

	var config Config
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		fmt.Printf("error decoding json file: %v\n", err)
		return Config{}
	}

	return config
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := fmt.Sprintf("%v/%v", homeDir, configFileName)

	return configPath, nil
}

func (c *Config) SetUser() error {
	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("error getting user name: %v", err)
	}

	c.CurrentUserName = currentUser.Username

	err = write(c)

	return nil
}

func write(cfg *Config) error {

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error marshalling cfg: %v", err)
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %v", err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Println("Shit Written Correctly!")
	return nil
}
