package utils

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

// GetPassword prompts the user for a password without echoing.
func GetPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(bytePassword), nil
}
