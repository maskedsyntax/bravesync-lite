package paths

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	BraveConfigPath = ".config/BraveSoftware/Brave-Browser/Default"
	BookmarksFile   = "Bookmarks"
	LoginDataFile   = "Login Data"
)

// GetBraveProfilePath returns the path to the Brave profile directory.
func GetBraveProfilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %w", err)
	}
	return filepath.Join(home, BraveConfigPath), nil
}

// GetBraveFiles returns the absolute paths to Bookmarks and Login Data.
func GetBraveFiles() (string, string, error) {
	profilePath, err := GetBraveProfilePath()
	if err != nil {
		return "", "", err
	}
	return filepath.Join(profilePath, BookmarksFile), filepath.Join(profilePath, LoginDataFile), nil
}

// EnsureDir checks if a directory exists and creates it if it doesn't.
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0700)
}
