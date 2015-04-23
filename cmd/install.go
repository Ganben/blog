package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/entity"
	"github.com/gofxh/blog/lib/log"
	"github.com/gofxh/blog/model"
	"time"
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
		result := base.Action.Call(CreateDefaultData, nil)
		if !result.Meta.Status {
			log.Fatal("Install but error : %s", result.Meta.ErrorMessage)
		}

		// install done
		log.Info("Install Success !")
	}
)

func CreateDefaultData(_ interface{}) *core.ActionResult {
	now := time.Now()

	base.Max = model.NewMaxer()

	// default user
	user := &entity.User{
		Id:         base.Max.Next(base.Max.UserId, base.Max.UserStep),
		Name:       "admin",
		Nick:       "admin",
		Email:      "admin@example.com",
		CreateTime: now.Unix(),
		Status:     entity.USER_STATUS_ACTIVE,
		Role:       entity.USER_ROLE_ADMIN,
	}
	user.Password, user.Salt = entity.GenerateUserPassword("123456")
	base.Storage.Save(user)
	base.Max.UserId = user.Id

	// write max file
	base.Storage.Save(base.Max)

	// return result
	return core.NewOKActionResult(nil)
}
