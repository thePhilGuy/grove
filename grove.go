package main

import (
	"fmt"
	"log"
	"os"

	git "gopkg.in/libgit2/git2go.v26"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panicln("Could not read current directory.")
	}

	grover := cli.NewApp()
	grover.Commands = []cli.Command{
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "check if the current directory is a git repository",
			Action: func(c *cli.Context) error {
				repository, err := git.OpenRepository(cwd)
				if err != nil {
					fmt.Println("Nope, not a git repo.")
					return fmt.Errorf("Not a git repo.")
				} else {
					fmt.Println("Yup, git repo.")
					repository.Free()
					return nil
				}
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initializes a git repository",
			Action: func(c *cli.Context) error {
				repository, err := git.InitRepository(cwd, false)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Failed to initialize grove repository")
					return err
				} else {
					fmt.Printf("Initialized empty Git repository in %s\n", cwd)
					repository.Free()
					return nil
				}
			},
		},
	}

	grover.Run(os.Args)
}
