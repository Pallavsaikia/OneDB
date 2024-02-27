package config

import (
	"encoding/json"
	"fmt"
	"onedb-core/constants"
	"onedb-core/engine/cache"
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

func (config *Config) setCache(c *cache.Cache) error {
	if config.PORT == 0 || config.DEFAULT_USER == "" || config.DEFAULT_PASSWORD == "" {
		return fmt.Errorf("error:cannot add empty config")
	}
	c.Set(constants.CONFIG_CACHE, config, cache.TTL_IN_SECOND)
	return nil
}

func (config *Config) getCache(c *cache.Cache) (*Config, error) {

	if s, found := c.Get(constants.CONFIG_CACHE); found {
		return s.(*Config), nil
	}
	return &Config{}, fmt.Errorf("error:config not available in cache")
}



func ReadConfig(c *cache.Cache) (Config, error) {
	loadedConfig := &Config{}
	loadedConfig, err := loadedConfig.getCache(c)
	if err == nil {
		return *loadedConfig, nil
	}
	fileroot, err := filesys.GetFileLocation()
	if err != nil {
		return Config{}, fmt.Errorf("error getting file location:%v", err)
	}
	*loadedConfig, err = loadConfigFromFile(filesys.CreatePathFromStringArray([]string{fileroot, constants.CONFIG_FILE}))
	if err != nil {
		return Config{}, fmt.Errorf("error loading config:%v", err)
	}
	loadedConfig.setCache(c)
	return *loadedConfig, nil
}

func InstallConfig(PORT int, DEFAULT_USER string, DEFAULT_PASSWORD string, DATABASE_STORAGE_ROOT string, c *cache.Cache) (Config, error) {
	configuration := &Config{}
	configuration.PORT = PORT
	configuration.DEFAULT_USER = DEFAULT_USER
	configuration.DEFAULT_PASSWORD = DEFAULT_PASSWORD
	configuration.DATABASE_STORAGE_ROOT = DATABASE_STORAGE_ROOT
	_, err := WriteConfig(*configuration, c)
	if err != nil {
		return Config{}, err
	}
	return *configuration, nil
}

func WriteConfig(config Config, c *cache.Cache) (Config, error) {
	fileroot, err := filesys.GetFileLocation()
	if err != nil {
		return Config{}, fmt.Errorf("error getting file location:%v", err)
	}
	config.setCache(c)
	return config, updateConfigToFile(filesys.CreatePathFromStringArray([]string{fileroot, constants.CONFIG_FILE}), config)
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
