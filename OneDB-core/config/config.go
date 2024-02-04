package config

import (
	"encoding/json"
	"fmt"
	"onedb-core/filesys"
	"os"
)

type Config struct {
	PORT                  int    `json:"database_port"`
	DEFAULT_USER          string `json:"default_user"`
	DEFAULT_PASSWORD      string `json:"default_password"`
	DATABASE_STORAGE_ROOT string `json:"database_storage_root"`
}

func (c *Config) Validate() (bool, error) {
	if c.PORT == 0 || c.DEFAULT_USER == "" || c.DEFAULT_PASSWORD == "" {
		return false, fmt.Errorf("fields cannot be empty")
	}
	return true, nil
}
func (c *Config) Update(value Config) {
	if value.PORT != 0 {
		c.PORT = value.PORT
	}
	if value.DEFAULT_USER != "" {
		c.DEFAULT_USER = value.DEFAULT_USER
	}
	if value.DEFAULT_PASSWORD != "" {
		c.DEFAULT_PASSWORD = value.DEFAULT_PASSWORD
	}
	c.DATABASE_STORAGE_ROOT = value.DATABASE_STORAGE_ROOT
}

const (
	CONFIG_FILE = "config.json"
)

func ReadConfig() (Config, error) {
	fileroot, err := filesys.GetFileLocation()
	if err != nil {
		return Config{}, fmt.Errorf("error getting file location:%v", err)
	}
	loadedConfig, err := loadConfigFromFile(fileroot + "/" + CONFIG_FILE)
	if err != nil {
		return Config{}, fmt.Errorf("error loading config:%v", err)
	}
	return loadedConfig, nil
}

func InstallConfig(PORT int, DEFAULT_USER string, DEFAULT_PASSWORD string, DATABASE_STORAGE_ROOT string) (Config, error) {
	configuration, err := ReadConfig()
	if err != nil {
		return Config{}, err
	}
	configuration.PORT = PORT
	configuration.DEFAULT_USER = "root"
	configuration.DEFAULT_PASSWORD = "root"
	configuration.DATABASE_STORAGE_ROOT = DATABASE_STORAGE_ROOT
	_, err = WriteConfig(configuration)
	if err != nil {
		return Config{}, err
	}
	return configuration, nil
}

func WriteConfig(config Config) (Config, error) {
	fileroot, err := filesys.GetFileLocation()
	if err != nil {
		return Config{}, fmt.Errorf("error getting file location:%v", err)
	}
	configuration, err := ReadConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error reading file")
	}
	configuration.Update(config)
	return configuration, updateConfigToFile(fileroot+"/"+CONFIG_FILE, configuration)
}

func updateConfigToFile(filePath string, config Config) error {
	isValid, _ := config.Validate()
	if isValid {
		configJSON, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return fmt.Errorf("error marshaling JSON:%v", err)
		}
		err = os.WriteFile(filePath, configJSON, 0644)
		if err != nil {
			return fmt.Errorf("error overwriting file:%v", err)

		}
	} else {
		return fmt.Errorf("file config is not valid")
	}
	return nil
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
