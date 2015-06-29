package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/core"
	"github.com/gofxh/blog/app/log"
	"github.com/gofxh/blog/app/model"
	"os"
	"path/filepath"
	"time"
)

var Init = cli.Command{
	Name:   "init",
	Usage:  "init default data and settings to setup",
	Action: InitAction,
}

func InitAction(ctx *cli.Context) {
	// check config,
	// is install time > 0, show installed message
	// otherwise, set install time and write new config file
	if app.Config.AppInstallTime > 0 {
		log.Warn("Blog is installed as %s", time.Unix(app.Config.AppInstallTime, 0).Format(time.RFC1123Z))
		return
	}
	app.Config.AppInstallTime = time.Now().Unix()
	app.Config.Write()

	// make directories
	os.Mkdir(app.Config.UserDirectory, os.ModePerm)
	os.Mkdir(filepath.Join(app.Config.UserDirectory, app.Config.UserThemeDirectory), os.ModePerm)
	os.Mkdir(filepath.Join(app.Config.UserDirectory, app.Config.UserUploadDirectory), os.ModePerm)

	// init database schema
	app.Db = core.NewDatabase(filepath.Join(app.Config.UserDirectory, app.Config.UserDataFile))
	app.Db.Sync2(new(model.User), new(model.Token))

	log.Info("Blog is installed successfully !!!")
}
