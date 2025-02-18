package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	filepath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filepath += "/" + configFileName
	return filepath, nil
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	dat, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = json.Unmarshal(dat, &config)
	if err != nil {
		return config, err
	}
	fmt.Println(filepath)
	fmt.Println(dat)
	fmt.Println(config)
	return config, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	return write(*cfg)
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, dat, 'w')
	if err != nil {
		return err
	}
	return nil
}
