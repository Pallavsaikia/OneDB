package config

import (
	"encoding/json"
	"fmt"
	"onedb-core/filesys"
	"os"
)
type Config struct {
	PORT                  string `json:"database_port"`
	DEFAULT_USER          string `json:"default_user"`
	DEFAULT_PASSWORD      string `json:"default_password"`
	DATABASE_STORAGE_ROOT string `json:"database_storage_root"`
}
const (
	CONFIG_FILE = "config.json"
)
func ReadConfig() (Config, error) {
	fileroot, err := filesys.GetFileLocation()
	if err != nil {
		fmt.Println("Error getting file location:", err)
		return Config{}, nil
	}
	loadedConfig, err := loadConfigFromFile(fileroot + "/" + CONFIG_FILE)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return Config{}, nil
	}
	return loadedConfig, nil
}

func loadConfigFromFile(filePath string) (Config, error) {
	configJSON, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file %s: %v", filePath, err)
	}
	var config Config
	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return config, nil
}
