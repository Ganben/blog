package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/action"
	"github.com/gofxh/blog/app/core"
	"github.com/gofxh/blog/app/log"
	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
	"github.com/tango-contrib/renders"
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

	// init server
	action.Call(InitServer, nil)
	// init router
	action.Call(InitRoute, nil)

	// start server
	core.Start(app.Server)

	log.Info("Serv|Close|%.1fms", time.Since(t).Seconds()*1000)
}

// init server,
// set static files, theme and middleware
func InitServer(_ interface{}) *action.Result {

	// set static directory
	app.Server.Use(tango.Static(tango.StaticOptions{
		RootPath: filepath.Join(app.Config.UserDirectory, app.Config.UserThemeDirectory),
		Prefix:   "theme",
	}))
	app.Server.Use(tango.Static(tango.StaticOptions{
		RootPath: filepath.Join(app.Config.UserDirectory, app.Config.UserUploadDirectory),
		Prefix:   "upload",
	}))

	// set theme directory
	app.Server.Use(renders.New(renders.Options{
		Reload:     true,
		Directory:  filepath.Join(app.Config.UserDirectory, app.Config.UserThemeDirectory),
		Extensions: []string{".tmpl"},
	}))

	// binding middleware
	app.Server.Use(binding.Bind())

	return action.OkResult(nil)
}

// init router,
// set route rules
func InitRoute(_ interface{}) *action.Result {
	return action.OkResult(nil)
}
