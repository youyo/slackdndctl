package main

import (
	"os"

	"github.com/urfave/cli"
)

var Name string
var Version string

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "youyo"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
