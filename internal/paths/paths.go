package paths

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	BraveConfigPath  = ".config/BraveSoftware/Brave-Browser/Default"
	HeliumConfigPath = ".config/net.imput.helium/Default"
	BookmarksFile    = "Bookmarks"
	LoginDataFile    = "Login Data"
)

// GetBrowserProfilePath returns the path to the specified browser's profile directory.
func GetBrowserProfilePath(browserName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %w", err)
	}

	var relPath string
	switch browserName {
	case "Helium":
		relPath = HeliumConfigPath
	case "Brave":
		fallthrough
	default:
		relPath = BraveConfigPath
	}

	return filepath.Join(home, relPath), nil
}

// GetBrowserFiles returns the absolute paths to Bookmarks and Login Data for the browser.
func GetBrowserFiles(browserName string) (string, string, error) {
	profilePath, err := GetBrowserProfilePath(browserName)
	if err != nil {
		return "", "", err
	}
	return filepath.Join(profilePath, BookmarksFile), filepath.Join(profilePath, LoginDataFile), nil
}

// EnsureDir checks if a directory exists and creates it if it doesn't.
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0700)
}
