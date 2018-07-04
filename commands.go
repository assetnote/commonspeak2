package main

import (
	"fmt"
	"github.com/assetnote/commonspeak2/command"
	"github.com/urfave/cli"
	"os"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "clitest",
		Usage:  "",
		Action: command.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Value: "",
				Usage: "The Google Cloud Project to use for the queries.",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
