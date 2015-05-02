package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/log"
	"github.com/gofxh/blog/mvc/helper"
)

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

		// create default config
		base.Config = core.NewConfig()
		// if exist, throw error
		if base.Config.Exist() {
			log.Fatal("Blog was installed !")
		}
		if err := base.Config.WriteFile(); err != nil {
			log.Fatal("Config|WriteFile|%s", err.Error())
		}

		// init default data
		base.Storage = core.NewStorage(base.Config.DataDirectory)
		result := base.Action.Call(helper.CreateDefaultData, nil)
		if !result.Meta.Status {
			log.Fatal("Install but error : %s", result.Meta.ErrorMessage)
		}

		log.Info("Data are in %s", base.Config.DataDirectory)
		log.Info("Http server will be running in %s", base.Config.HttpAddress)

		// install done
		log.Info("Install Success !")
	}
)
