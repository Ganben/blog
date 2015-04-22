package cmd

import "github.com/codegangsta/cli"

var (
	installCommand cli.Command = cli.Command{
		Name:   "install",
		Usage:  "install blog engine",
		Action: installCommandFunc,
	}

	installCommandFunc = func(ctx *cli.Context) {

	}
)
