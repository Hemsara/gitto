package keys

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	APIKey string `json:"api_key"`
}

func GetConfigFilePath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(configDir, "gitto", "config.json")
}

func SaveAPIKey(apiKey string) error {
	config := Config{APIKey: apiKey}
	configPath := GetConfigFilePath()

	if err := os.MkdirAll(filepath.Dir(configPath), 0700); err != nil {
		return err
	}

	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(config); err != nil {
		return err
	}

	fmt.Println("API key saved successfully!")
	return nil
}

func LoadAPIKey() (string, error) {
	configPath := GetConfigFilePath()
	file, err := os.Open(configPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return "", err
	}

	return config.APIKey, nil
}
