// This file contains the functions used for interacting with the config for lume

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	configDirName  = ".config"
	appConfigDir   = "lume"
	configFileName = "config.json"
)

// Config represents the application configuration
type Config struct {
	ApiKey string `json:"govee_api_key"`
}

// Global variable to hold the loaded configuration
var ApiKey string

// Init loads the configuration and sets global variables
func Init() {
	config := Load()
	ApiKey = config.ApiKey
}



func Load() Config {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	configDir := filepath.Join(home, configDirName, appConfigDir)
	configPath := filepath.Join(configDir, configFileName)

	config, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	var configData Config
	err = json.Unmarshal(config, &configData)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	return configData
}

func Save(config Config) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	configDir := filepath.Join(home, configDirName, appConfigDir)
	configPath := filepath.Join(configDir, configFileName)

	// Ensure the config directory exists
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0700)
		if err != nil {
			fmt.Printf("Failed to create config directory: %v\n", err)
			return
		}
	}

	configData, err := json.Marshal(config)
	if err != nil {
		fmt.Printf("Failed to marshal config: %v\n", err)
		return
	}

	err = os.WriteFile(configPath, configData, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetApiKey() string {
	config := Load()
	return config.ApiKey
}

func PrintConfig() {
	config := Load()
	fmt.Println(config)
}

func SetValue(key string, value string) {
	config := Load()

	switch key {
	case "govee_api_key":
		config.ApiKey = value
	}

	Save(config)
}