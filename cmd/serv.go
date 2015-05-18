package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/log"
	"github.com/gofxh/blog/mvc/controller"
	"github.com/gofxh/blog/mvc/crond"
	"github.com/gofxh/blog/mvc/helper"
	"github.com/gofxh/blog/mvc/model"
	"github.com/gofxh/blog/plugin"
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

		// init base vars
		base.Storage = core.NewStorage(base.Config.DataDirectory)
		base.Cron = core.NewCron()
		base.Server = core.NewServer(base.Config)
		base.Max = model.NewMaxer()

		// load data
		base.Action.Call(model.Init, nil)

		// start cron
		base.Action.Call(crond.Init, nil)

		// start server
		base.Action.Call(controller.Init, nil)
		log.Info("Http server is running in %s", base.Config.HttpAddress)

		// init helper
		base.Action.Call(helper.Init, nil)

		// init plugin
		base.Action.Call(plugin.Init, nil)

		core.Start(base.Server)
	}
)
