package main

import (
	"os"

	"github.com/micnncim/readme/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "readme"
	app.Usage = "Generate README.md"
	app.Action = cmd.Readme

	app.Run(os.Args)
}
