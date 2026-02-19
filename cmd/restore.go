package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/maskedsyntax/bravesync-lite/internal/archive"
	"github.com/maskedsyntax/bravesync-lite/internal/config"
	"github.com/maskedsyntax/bravesync-lite/internal/encryption"
	"github.com/maskedsyntax/bravesync-lite/internal/github"
	"github.com/maskedsyntax/bravesync-lite/internal/paths"
	"github.com/maskedsyntax/bravesync-lite/internal/tui"
	"github.com/maskedsyntax/bravesync-lite/internal/utils"
	"github.com/urfave/cli/v2"
)

func RestoreAction(c *cli.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w. Run 'config' first", err)
	}

	gm := github.NewGitManager(cfg.GitHubRepoURL, cfg.LocalClonePath, cfg.Branch, cfg.PAT)
	fmt.Println("Pulling latest backups from GitHub...")
	if err := gm.Clone(); err != nil {
		return err
	}
	if err := gm.Pull(); err != nil {
		return err
	}

	latestPath := filepath.Join(cfg.LocalClonePath, "latest.enc")
	if _, err := os.Stat(latestPath); os.IsNotExist(err) {
		return fmt.Errorf("latest backup not found in local repo: %s", latestPath)
	}

	isTUI := IsTUIMode
	var password string
	if isTUI {
		password, err = tui.GetPasswordTUI("Enter decryption password: ")
	} else {
		password, err = utils.GetPassword("Enter decryption password: ")
	}
	if err != nil {
		return err
	}
	if password == "" {
		return fmt.Errorf("password cannot be empty")
	}
	defer utils.ZeroMem([]byte(password))

	fmt.Println("Decrypting...")
	encryptedData, err := os.ReadFile(latestPath)
	if err != nil {
		return err
	}

	decryptedData, err := encryption.Decrypt(encryptedData, password)
	if err != nil {
		return err
	}
	defer utils.ZeroMem(decryptedData)

	braveProfilePath, err := paths.GetBraveProfilePath()
	if err != nil {
		return err
	}

	fmt.Printf("Restoring to %s. This will overwrite existing files. Continue? [y/N]: ", braveProfilePath)
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) != "y" {
		fmt.Println("Restore aborted.")
		return nil
	}

	if err := archive.ExtractArchive(decryptedData, braveProfilePath); err != nil {
		return err
	}

	fmt.Println("Restore successful! Please restart Brave browser.")
	return nil
}
