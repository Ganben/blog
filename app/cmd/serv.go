package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/core"
	"github.com/gofxh/blog/app/log"
	"path/filepath"
	"time"
)

var Serv = cli.Command{
	Name:   "serv",
	Usage:  "start web server",
	Action: ServAction,
}

func ServAction(ctx *cli.Context) {
	t := time.Now()
	log.Info("Serv|Begin")

	// init global vars
	app.Db = core.NewDatabase(filepath.Join(app.Config.UserDirectory, app.Config.UserDataFile))
	app.Server = core.NewServer(fmt.Sprintf("%s:%s", app.Config.HttpHost, app.Config.HttpAddress))

	// start server
	core.Start(app.Server)

	log.Info("Serv|Close|%.1fms", time.Since(t).Seconds()*1000)
}
