package cmd

import "github.com/codegangsta/cli"

var (
	installCommand cli.Command = cli.Command{
		Name:   "install",
		Usage:  "install blog engine",
		Action: installCommandFunc,
	}

	/*
	   install steps:
	   1. download static files
	   2. write default config file
	   3. init default data
	*/
	installCommandFunc = func(ctx *cli.Context) {

	}
)
