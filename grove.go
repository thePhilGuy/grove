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

	checkRepository := func(c *cli.Context) error {
		repository, err := git.OpenRepository(cwd)
		if err != nil {
			fmt.Println("Nope, not a git repo.")
			return fmt.Errorf("Not a git repo.")
		}
		fmt.Println("Yup, git repo.")
		repository.Free()
		return nil
	}

	initRepository := func(c *cli.Context) error {
		repository, err := git.InitRepository(cwd, false)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to initialize grove repository")
			return err
		}
		fmt.Printf("Initialized empty Git repository at %s\n", cwd)
		repository.Free()
		return nil
	}

	listBranches := func(c *cli.Context) error {
		repository, err := git.OpenRepository(cwd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open git repository at %s\n", cwd)
			return err
		}
		branchIterator, err := repository.NewBranchIterator(git.BranchAll)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to get branch iterator")
			repository.Free()
			return err
		}

		err = branchIterator.ForEach(func(branch *git.Branch, branchType git.BranchType) error {
			branchName, err := branch.Name()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to get branch name")
				repository.Free()
				return err
			}
			fmt.Println(branchName)
			return nil
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to list branches")
			branchIterator.Free()
			repository.Free()
			return err
		}
		branchIterator.Free()
		repository.Free()
		return nil
	}

	grover := cli.NewApp()
	grover.Commands = []cli.Command{
		{
			Name:   "check",
			Usage:  "check if the current directory is a git repository",
			Action: checkRepository,
		},
		{
			Name:   "init",
			Usage:  "initializes a git repository",
			Action: initRepository,
		},
		{
			Name:  "branch",
			Usage: "list, create, or delete branches",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"ls", "l"},
					Usage:   "list branches in current repository",
					Action:  listBranches,
				},
			},
		},
	}

	grover.Run(os.Args)
}
