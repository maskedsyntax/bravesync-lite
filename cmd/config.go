package cmd

import (
	"fmt"

	"github.com/maskedsyntax/bravesync-lite/internal/config"
	"github.com/urfave/cli/v2"
)

func ConfigAction(c *cli.Context) error {
	cfg, err := config.Load()
	if err != nil {
		// Create default if not exists
		cfg = &config.Config{
			Branch: "main",
			Encryption: config.EncryptionConfig{
				KDF:        "argon2id",
				Iterations: 1,
			},
		}
	}

	repoURL := c.String("repo")
	if repoURL != "" {
		cfg.GitHubRepoURL = repoURL
	}

	localPath := c.String("local")
	if localPath != "" {
		cfg.LocalClonePath = localPath
	}

	branch := c.String("branch")
	if branch != "" {
		cfg.Branch = branch
	}

	pat := c.String("pat")
	if pat != "" {
		cfg.PAT = pat
	}

	if repoURL == "" && localPath == "" && branch == "" && pat == "" {
		// Just print current config
		fmt.Printf("Current Configuration:\n")
		fmt.Printf("GitHub Repo URL:  %s\n", cfg.GitHubRepoURL)
		fmt.Printf("Local Clone Path: %s\n", cfg.LocalClonePath)
		fmt.Printf("Branch:           %s\n", cfg.Branch)
		if cfg.PAT != "" {
			fmt.Printf("PAT:              %s\n", "****")
		} else {
			fmt.Printf("PAT:              Not set (will use BRAVE_SYNC_PAT env var)\n")
		}
		return nil
	}

	if err := cfg.Save(); err != nil {
		return err
	}

	fmt.Println("Configuration saved successfully.")
	return nil
}
