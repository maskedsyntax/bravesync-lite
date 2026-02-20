package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/maskedsyntax/bravesync-lite/internal/archive"
	"github.com/maskedsyntax/bravesync-lite/internal/config"
	"github.com/maskedsyntax/bravesync-lite/internal/encryption"
	"github.com/maskedsyntax/bravesync-lite/internal/github"
	"github.com/maskedsyntax/bravesync-lite/internal/paths"
	"github.com/maskedsyntax/bravesync-lite/internal/tui"
	"github.com/maskedsyntax/bravesync-lite/internal/utils"
	"github.com/urfave/cli/v2"
)

func BackupAction(c *cli.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w. Run 'config' first", err)
	}

	isTUI := IsTUIMode

	var password string
	if isTUI {
		password, err = tui.GetPasswordTUI("Enter encryption password: ")
	} else {
		password, err = utils.GetPassword("Enter encryption password: ")
	}
	if err != nil {
		return err
	}
	if password == "" {
		return fmt.Errorf("password cannot be empty")
	}
	defer utils.ZeroMem([]byte(password))

	bookmarkPath, loginDataPath, err := paths.GetBrowserFiles(cfg.Browser)
	if err != nil {
		return err
	}

	fmt.Printf("Backing up %s data...\n", cfg.Browser)
	fmt.Println("Creating archive...")
	archiveData, err := archive.CreateArchive([]string{bookmarkPath, loginDataPath})
	if err != nil {
		return err
	}
	defer utils.ZeroMem(archiveData)

	fmt.Println("Encrypting...")
	encryptedData, err := encryption.Encrypt(archiveData, password)
	if err != nil {
		return err
	}

	// Prepare Git Repo first to ensure directory exists
	fmt.Println("Preparing Git repository...")
	gm := github.NewGitManager(cfg.GitHubRepoURL, cfg.LocalClonePath, cfg.Branch, cfg.PAT)
	if err := gm.Clone(); err != nil {
		return err
	}
	if err := gm.Pull(); err != nil {
		fmt.Printf("Warning: pull failed: %v\n", err)
	}

	timestamp := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("backup-%s.enc", timestamp)
	localFilePath := filepath.Join(cfg.LocalClonePath, filename)

	fmt.Println("Saving to local repository...")
	if err := os.WriteFile(localFilePath, encryptedData, 0600); err != nil {
		return err
	}

	latestPath := filepath.Join(cfg.LocalClonePath, "latest.enc")
	if err := os.WriteFile(latestPath, encryptedData, 0600); err != nil {
		return err
	}

	fmt.Println("Pushing to GitHub...")
	if err := gm.AddCommitPush(filename, fmt.Sprintf("Backup %s %s", cfg.Browser, timestamp)); err != nil {
		return err
	}
	if err := gm.AddCommitPush("latest.enc", fmt.Sprintf("Update latest backup %s %s", cfg.Browser, timestamp)); err != nil {
		return err
	}

	fmt.Printf("Backup successful: %s\n", filename)
	return nil
}
