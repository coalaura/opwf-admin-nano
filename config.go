package main

import (
	_ "embed"
	"errors"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Token string `json:"token"`
}

func LoadConfig() (*Config, error) {
	path, err := ConfigPath()
	if err != nil {
		return nil, err
	}

	config, err := ReadConfig(path)

	if os.IsNotExist(err) {
		os.WriteFile(path, []byte("# your opfw api token\ntoken: \"\"\n"), 0644)

		return nil, errors.New("created empty config at ~/.config/opfw-nano.yml")
	} else if err != nil {
		return nil, err
	}

	if config.Token == "" {
		return nil, errors.New("missing token in config at ~/.config/opfw-nano.yml")
	}

	return config, nil
}

func ReadConfig(path string) (*Config, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var config Config

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func ConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".config", "opfw-nano.yml"), nil
}
