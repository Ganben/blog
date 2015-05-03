package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/log"
	"github.com/gofxh/blog/mvc/controller"
	"github.com/gofxh/blog/mvc/helper"
	"github.com/gofxh/blog/mvc/model"
)

var (
	servCommand cli.Command = cli.Command{
		Name:   "serv",
		Usage:  "serv blog engine",
		Action: servCommandFunc,
	}

	/*
	   server steps:
	   1. load config, check install
	   2. load data to memory
	   3. start cron
	   4. start http server
	*/
	servCommandFunc = func(ctx *cli.Context) {
		// load config
		base.Config = core.NewConfig()
		// if not exist, need install
		if !base.Config.Exist() {
			log.Fatal("Blog was not installed yet !")
		}

		// load data
		base.Storage = core.NewStorage(base.Config.DataDirectory)
		model.Init()

		// init helper
		helper.Init()

		// start cron

		// start server
		log.Info("Http server is running in %s", base.Config.HttpAddress)
		base.Server = core.NewServer(base.Config)
		controller.Init()

		core.Start(base.Server)
	}
)
