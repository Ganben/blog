package model

import (
	"github.com/Unknwon/cae/zip"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/entity"
	"time"
)

func CreateDefaultData(_ interface{}) *core.ActionResult {
	now := time.Now()

	base.Max = NewMaxer()

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

func CreateZipData(_ interface{}) *core.ActionResult {
	filename := time.Now().Format("20060102150405.zip")
	z, err := zip.Create(filename)
	if err != nil {
		return core.NewSystemErrorResult(err)
	}

	zip.Verbose = false

	// add real data to zip
	z.AddDir(base.Config.DataDirectory, base.Config.DataDirectory)

	if err = z.Flush(); err != nil {
		return core.NewSystemErrorResult(err)
	}

	z.Close()

	return core.NewOKActionResult(core.AData{
		"BackupFile": filename,
	})
}
