package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/maskedsyntax/bravesync-lite/internal/config"
	"github.com/urfave/cli/v2"
)

func promptInput(prompt string, current string) string {
	reader := bufio.NewReader(os.Stdin)
	if current != "" {
		fmt.Printf("%s [%s]: ", prompt, current)
	} else {
		fmt.Printf("%s: ", prompt)
	}
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return current
	}
	return input
}

func ConfigAction(c *cli.Context) error {
	cfg, err := config.Load()
	if err != nil {
		defaultLocal, _ := config.GetDefaultLocalClonePath()
		cfg = &config.Config{
			LocalClonePath: defaultLocal,
			Branch:         "main",
			Browser:        "Brave",
			Encryption: config.EncryptionConfig{
				KDF:        "argon2id",
				Iterations: 1,
			},
		}
	}

	repoURL := c.String("repo")
	localPath := c.String("local")
	branch := c.String("branch")
	pat := c.String("pat")
	browser := c.String("browser")

	// If no flags are provided, enter interactive mode
	if repoURL == "" && localPath == "" && branch == "" && pat == "" && browser == "" {
		fmt.Println("Entering interactive configuration mode. Press enter to keep current values.")
		cfg.Browser = promptInput("Browser (Brave/Helium)", cfg.Browser)
		cfg.GitHubRepoURL = promptInput("GitHub Repo URL", cfg.GitHubRepoURL)
		cfg.LocalClonePath = promptInput("Local Clone Path", cfg.LocalClonePath)
		cfg.Branch = promptInput("Branch", cfg.Branch)
		cfg.PAT = promptInput("GitHub PAT (leave blank to use BRAVE_SYNC_PAT env var)", cfg.PAT)
	} else {
		// Update only provided flags
		if browser != "" {
			cfg.Browser = browser
		}
		if repoURL != "" {
			cfg.GitHubRepoURL = repoURL
		}
		if localPath != "" {
			cfg.LocalClonePath = localPath
		}
		if branch != "" {
			cfg.Branch = branch
		}
		if pat != "" {
			cfg.PAT = pat
		}
	}

	if err := cfg.Save(); err != nil {
		return err
	}

	fmt.Println("Configuration saved successfully.")
	return nil
}
