package github

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GitManager handles interactions with the system git.
type GitManager struct {
	RepoURL        string
	LocalClonePath string
	Branch         string
	PAT            string
}

// NewGitManager creates a new GitManager.
func NewGitManager(repoURL, localPath, branch, pat string) *GitManager {
	return &GitManager{
		RepoURL:        repoURL,
		LocalClonePath: localPath,
		Branch:         branch,
		PAT:            pat,
	}
}

// AuthenticatedURL returns the repo URL with the PAT embedded.
func (m *GitManager) AuthenticatedURL() string {
	// Assuming URL is https://github.com/user/repo.git
	if m.PAT == "" {
		return m.RepoURL
	}
	parts := strings.Split(m.RepoURL, "https://")
	if len(parts) < 2 {
		return m.RepoURL
	}
	return fmt.Sprintf("https://%s@%s", m.PAT, parts[1])
}

// Clone clones the repository if it doesn't exist.
func (m *GitManager) Clone() error {
	if _, err := os.Stat(m.LocalClonePath); err == nil {
		return nil // Already cloned
	}

	parentDir := filepath.Dir(m.LocalClonePath)
	if err := os.MkdirAll(parentDir, 0700); err != nil {
		return fmt.Errorf("failed to create parent directory: %w", err)
	}

	cmd := exec.Command("git", "clone", "-b", m.Branch, m.AuthenticatedURL(), m.LocalClonePath)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git clone failed: %s: %w", string(output), err)
	}

	return nil
}

// Pull pulls the latest changes.
func (m *GitManager) Pull() error {
	cmd := exec.Command("git", "-C", m.LocalClonePath, "pull", "origin", m.Branch)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git pull failed: %s: %w", string(output), err)
	}
	return nil
}

// AddCommitPush adds a file, commits, and pushes to the repository.
func (m *GitManager) AddCommitPush(filename, commitMsg string) error {
	// Add
	addCmd := exec.Command("git", "-C", m.LocalClonePath, "add", filename)
	if output, err := addCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git add failed: %s: %w", string(output), err)
	}

	// Commit
	commitCmd := exec.Command("git", "-C", m.LocalClonePath, "commit", "-m", commitMsg)
	if output, err := commitCmd.CombinedOutput(); err != nil {
		// It might fail if there are no changes, but since we add a new timestamped file, it should work.
		return fmt.Errorf("git commit failed: %s: %w", string(output), err)
	}

	// Push
	pushCmd := exec.Command("git", "-C", m.LocalClonePath, "push", "origin", m.Branch)
	if output, err := pushCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git push failed: %s: %w", string(output), err)
	}

	return nil
}
