package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db_url          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const (
	configFileName = ".gatorconfig.json"
)

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("error retrieving home directory: %v", err)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := fmt.Sprintf("%v/%v", homeDir, configFileName)

	return configPath, nil
}

func (c *Config) SetUser(username string) error {

	c.CurrentUserName = username

	err := Write(c)
	if err != nil {
		return err
	}

	return nil
}

func Write(cfg *Config) error {

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
