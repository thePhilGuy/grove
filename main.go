package main

import (
	"fmt"
	"os"

	"github.com/thePhilGuy/grove/cli"
	"github.com/thePhilGuy/grove/grovefile"
)

func main() {
	globalGrovefile := grovefile.Open()
	fmt.Println(globalGrovefile)

	groverCli := cli.Initialize()
	groverCli.Run(os.Args)
}
