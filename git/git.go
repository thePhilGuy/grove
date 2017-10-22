package git

import (
	"fmt"
	"log"
	"os"

	git2go "gopkg.in/libgit2/git2go.v26"
	"gopkg.in/urfave/cli.v1"
)

type Repository struct {
	ActiveBranch string
	BaseBranch   string
	Name         string
	Path         string
}

func getwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not read current directory.")
	}

	return cwd
}

func CheckRepository(c *cli.Context) error {
	cwd := getwd()

	repository, err := git2go.OpenRepository(cwd)
	defer repository.Free()
	if err != nil {
		fmt.Println("Nope, not a git repo.")
		return err
	}
	fmt.Println("Yup, git repo.")
	return nil
}

func InitializeRepository(c *cli.Context) error {
	cwd := getwd()

	repository, err := git2go.InitRepository(cwd, false)
	defer repository.Free()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to initialize grove repository")
		return err
	}
	fmt.Printf("Initialized empty git repository at %s\n", cwd)
	return nil
}

func ListBranches(c *cli.Context) error {
	cwd := getwd()

	repository, err := git2go.OpenRepository(cwd)
	defer repository.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open git repository at %s\n", cwd)
		return err
	}
	branchIterator, err := repository.NewBranchIterator(git2go.BranchAll)
	defer branchIterator.Free()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get branch iterator")
		return err
	}

	err = branchIterator.ForEach(func(branch *git2go.Branch, branchType git2go.BranchType) error {
		branchName, err := branch.Name()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to get branch name")
			return err
		}
		fmt.Println(branchName)
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to list branches")
		return err
	}
	return nil
}
