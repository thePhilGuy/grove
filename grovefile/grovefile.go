/*
	use a grovedir instead, backed by a git repository
	the branch will be detected using the git api
*/
package grovefile

import (
	"bufio"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/thePhilGuy/grove/git"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Grovefile struct {
	Path         string
	Repositories []git.Repository
	Branch       string
	File         *os.File
}

func Load(relativePath string) *Grovefile {
	usr, err := user.Current()
	check(err)

	path := usr.HomeDir + relativePath

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	check(err)
	defer file.Close()

	grovefile := Grovefile{
		Path:   path,
		File:   file,
		Branch: "master", // TODO when using a grovedir
	}

	fileinfo, err := file.Stat()
	check(err)
	if fileinfo.Size() > 0 {
		grovefile.loadRepositories()
	}

	return &grovefile
}

func isNameString(s string) bool {
	return strings.HasPrefix(nameString, "[") && strings.HasSuffix(nameString, "]")
}

func (grovefile *Grovefile) loadRepositories() {

	scanner := bufio.NewScanner(grovefile.File)

	repositories := make([]git.Repository, 2)

	nameString := scanner.Text()

	if !isNameString(nameString) {
		log.Fatal("Failed to parse repository name string: ", nameString)
	}

	for scanner.Scan() {
		// Scan for base branch and path
	}
	check(scanner.Err())

	grovefile.Repositories = repositories
}
