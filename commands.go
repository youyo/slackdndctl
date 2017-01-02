package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/youyo/slackdndctl/command"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "token",
		Usage:  "Set slack-token.",
		EnvVar: "SLACK_TOKEN",
	},
}

var Commands = []cli.Command{
	{
		Name:   "end-dnd",
		Usage:  "",
		Action: command.CmdEndDND,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
