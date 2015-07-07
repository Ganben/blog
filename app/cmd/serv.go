package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/action"
	"github.com/gofxh/blog/app/core"
	"github.com/gofxh/blog/app/log"
	"github.com/gofxh/blog/app/model"
	"github.com/gofxh/blog/app/route/admin"
	"github.com/gofxh/blog/app/route/base"
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
	address := fmt.Sprintf("%s:%s", app.Config.HttpHost, app.Config.HttpAddress)
	log.Info("Serv|Begin|%s", address)

	// init global vars
	app.Db = core.NewDatabase(filepath.Join(app.Config.UserDirectory, app.Config.UserDataFile))
	app.Server = core.NewServer(address)

	// read settings
	model.ReadSettingsToGlobal()

	// set other global vars with setting
	app.Theme = core.NewTheme(filepath.Join(app.Config.UserDirectory, app.Config.UserThemeDirectory),
		model.Settings["theme"].GetString())

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
		Prefix:   "/theme",
	}))
	app.Server.Use(tango.Static(tango.StaticOptions{
		RootPath: filepath.Join(app.Config.UserDirectory, app.Config.UserUploadDirectory),
		Prefix:   "/upload",
	}))

	// set theme directory
	app.Server.Use(renders.New(renders.Options{
		Reload:     true,
		Directory:  filepath.Join(app.Config.UserDirectory, app.Config.UserThemeDirectory),
		Extensions: []string{".tmpl"},
	}))

	// binding middleware
	app.Server.Use(binding.Bind())
	app.Server.Use(base.AuthHandler())

	return action.OkResult(nil)
}

// init router,
// set route rules
func InitRoute(_ interface{}) *action.Result {
	// admin routes
	adminGroup := tango.NewGroup()
	adminGroup.Any("/login", new(admin.Login))
	adminGroup.Get("/logout", new(admin.Logout))
	adminGroup.Any("/profile", new(admin.Profile))
	adminGroup.Post("/password", new(admin.Password))
	adminGroup.Any("/article/new", new(admin.Write))
	adminGroup.Get("/", new(admin.Admin))

	app.Server.Group("/admin", adminGroup)
	return action.OkResult(nil)
}
