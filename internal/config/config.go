package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigDir  = ".config/bravesynclite"
	ConfigFile = "config.json"
	PATEnvVar  = "BRAVE_SYNC_PAT"
)

// EncryptionConfig holds the KDF settings.
type EncryptionConfig struct {
	KDF        string `json:"kdf"`
	Iterations int    `json:"iterations"`
}

// Config is the main application configuration.
type Config struct {
	GitHubRepoURL  string           `json:"github_repo_url"`
	LocalClonePath string           `json:"local_clone_path"`
	Branch         string           `json:"branch"`
	Encryption     EncryptionConfig `json:"encryption"`
	PAT            string           `json:"pat,omitempty"`
}

// GetConfigPath returns the standard configuration file path.
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home dir: %w", err)
	}
	return filepath.Join(home, ConfigDir, ConfigFile), nil
}

// Load loads the configuration from the disk.
func Load() (*Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %w", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// Try to get PAT from environment if not set
	if cfg.PAT == "" {
		cfg.PAT = os.Getenv(PATEnvVar)
	}

	return &cfg, nil
}

// Save saves the configuration to the disk.
func (c *Config) Save() error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return fmt.Errorf("failed to create config dir: %w", err)
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}
