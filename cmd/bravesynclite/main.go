package main

import (
	"log"
	"os"

	"github.com/maskedsyntax/bravesync-lite/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "bravesynclite",
		Usage: "Brave browser data backup and restore tool",
		Commands: []*cli.Command{
			{
				Name:   "backup",
				Usage:  "Backup Brave data to GitHub",
				Action: cmd.BackupAction,
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "tui-mode", Hidden: true},
				},
			},
			{
				Name:   "restore",
				Usage:  "Restore Brave data from GitHub",
				Action: cmd.RestoreAction,
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "tui-mode", Hidden: true},
				},
			},
			{
				Name:  "config",
				Usage: "Manage configuration",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "repo", Usage: "GitHub repository URL"},
					&cli.StringFlag{Name: "local", Usage: "Local clone directory"},
					&cli.StringFlag{Name: "branch", Usage: "GitHub branch name"},
					&cli.StringFlag{Name: "pat", Usage: "GitHub Personal Access Token (PAT)"},
					&cli.BoolFlag{Name: "tui-mode", Hidden: true},
				},
				Action: cmd.ConfigAction,
			},
			{
				Name:   "tui",
				Usage:  "Run TUI mode",
				Action: cmd.TUIAction,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
