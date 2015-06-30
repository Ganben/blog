package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/action"
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
	t := time.Now()
	// check config,
	// is install time > 0, show installed message
	// otherwise, set install time and write new config file
	if app.Config.AppInstallTime > 0 {
		log.Info("Blog|Installed|%s", time.Unix(app.Config.AppInstallTime, 0).Format(time.RFC3339))
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
	action.Call(InitDbSchema, nil)
	action.Call(InitDbDefault, nil)

	log.Info("Blog|Install|Success|%.1fms", time.Since(t).Seconds()*1000)

}

func InitDbSchema(_ interface{}) *action.Result {
	app.Db.Sync2(new(model.User), new(model.Token), new(model.Article), new(model.Tag), new(model.Comment))
	return action.OkResult(nil)
}

func InitDbDefault(_ interface{}) *action.Result {
	// init administrator user
	user := model.NewDefaultUser()
	if err := model.SaveUser(user); err != nil {
		return action.ErrorResult(err)
	}

	// init welcome article
	article := model.NewDefaultArticle(user.Id)
	if err := model.SaveArticle(article); err != nil {
		return action.ErrorResult(err)
	}

	// init first comment
	comment := model.NewDefaultComment(article.Id)
	if err := model.SaveComment(comment); err != nil {
		return action.ErrorResult(err)
	}

	return action.OkResult(nil)
}
