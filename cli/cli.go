package cli

import (
	"github.com/thePhilGuy/grove/git"
	urfaveCli "gopkg.in/urfave/cli.v1"
)

func Initialize() *urfaveCli.App {
	grover := urfaveCli.NewApp()
	grover.Name = "grove"
	grover.Usage = "Work across multiple git repositories"
	grover.Version = "0.0.1"

	grover.Commands = []urfaveCli.Command{
		{
			Name:   "check",
			Usage:  "check if the current directory is a git repository",
			Action: git.CheckRepository,
		},
		{
			Name:   "init",
			Usage:  "initializes a git repository",
			Action: git.InitializeRepository,
		},
		{
			Name:  "branch",
			Usage: "list, create, or delete branches",
			Subcommands: []urfaveCli.Command{
				{
					Name:    "list",
					Aliases: []string{"ls", "l"},
					Usage:   "list branches in current repository",
					Action:  git.ListBranches,
				},
			},
		},
	}

	return grover
}
