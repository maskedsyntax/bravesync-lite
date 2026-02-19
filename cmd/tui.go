package cmd

import (
	"fmt"
	"github.com/maskedsyntax/bravesync-lite/internal/tui"
	"github.com/urfave/cli/v2"
)

var IsTUIMode bool

func TUIAction(c *cli.Context) error {
	IsTUIMode = true
	choice, err := tui.StartTUI()
	if err != nil {
		return err
	}

	switch choice {
	case "Backup":
		return BackupAction(c)
	case "Restore":
		return RestoreAction(c)
	case "Config":
		return ConfigAction(c)
	default:
		fmt.Println("No option selected.")
	}

	return nil
}
