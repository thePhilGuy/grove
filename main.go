package main

import (
	"fmt"
	"os"

	"github.com/thePhilGuy/grove/cli"
	"github.com/thePhilGuy/grove/grovefile"
)

func main() {
	// TODO get grovefile location from config
	globalGrovefile := grovefile.Load(".grovefile")
	fmt.Println(globalGrovefile)

	// TODO pass the grovefile context along to cli command parser
	groverCli := cli.Initialize()
	groverCli.Run(os.Args)
}
