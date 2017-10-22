package grovefile

import (
	"log"
	"os/user"
)

type Grovefile struct {
	Path   string
	Gits   []string
	Branch string
}

func Open() *Grovefile {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	gofilePath := usr.HomeDir + ".grovefile"

	return &Grovefile{
		Path:   gofilePath,
		Branch: "master",
	}
}
