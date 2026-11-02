package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Node struct {
		ID     string `yaml:"id"`
		Role   string `yaml:"role"`
		Branch string `yaml:"branch"`
	} `yaml:"node"`

	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Path string `yaml:"path"`
	} `yaml:"database"`

	Auth struct {
		Token string `yaml:"token"`
	} `yaml:"auth"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	if cfg.Auth.Token == "" {
		cfg.Auth.Token = os.Getenv("PHARMACY_AUTH_TOKEN")
	}
	if cfg.Auth.Token == "" {
		cfg.Auth.Token = "dev-pharmacist-token"
	}

	return &cfg, nil
}
