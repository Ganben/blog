package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/helper"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/log"
)

var (
	backupCommand cli.Command = cli.Command{
		Name:   "backup",
		Usage:  "backup blog engine",
		Action: backupCommandFunc,
	}

	/*
	   backup steps
	   1. load config
	   2. add file to backup zip
	   3. output zip
	*/
	backupCommandFunc = func(ctx *cli.Context) {
		// load config
		base.Config = core.NewConfig()
		// if not exist, need install
		if !base.Config.Exist() {
			log.Fatal("Blog was not installed yet !")
		}

		result := base.Action.Call(helper.CreateZipData, nil)
		if !result.Meta.Status {
			log.Fatal("Backup|Fail|%s", result.Meta.ErrorMessage)
		}

		log.Info("Backup to file %s", result.Data["BackupFile"])

	}
)
