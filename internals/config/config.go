package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Package struct {
	Name     string `json:"name"`
	Repo     string `json:"repo"`
	Branch   string `json:"branch"`
	OnCommit string `json:"onCommit"`
}

type Config struct {
	ViewScript bool      `json:"viewScript"`
	Packages   []Package `json:"packages"`
}

func LoadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}
