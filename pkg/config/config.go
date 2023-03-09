package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type DockerConfig struct {
	Item    string `json:"item"`
	Section string `json:"section"`
}

type Config struct {
	Docker DockerConfig `json:"docker"`
}

func (c *Config) Read() error {
	path, err := c.Path()
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err = createDefaultConfig(path); err != nil {
			return err
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func (c *Config) Path() (string, error) {
	path := os.Getenv("CRED_FILE")
	if path == "" {
		userDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		path = filepath.Join(userDir, ".config", "creds.json")
	}

	return path, nil
}

func createDefaultConfig(path string) error {
	config := &Config{
		Docker: DockerConfig{
			Item:    "Docker",
			Section: "Credentials",
		},
	}

	content, err := json.Marshal(config)
	if err != nil {
		return err
	}
	if err = os.WriteFile(path, content, 0777); err != nil {
		return err
	}

	return nil
}
