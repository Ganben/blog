package model

import (
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/log"
	"github.com/gofxh/blog/app/utils"
)

// user auth token struct
type Token struct {
	Id         int64
	UserId     int64  `xorm:"not null"` // token's owner
	Value      string `xorm:"not null"` // token value
	CreateTime int64  `xorm:"created"`  // create time
	ExpireTime int64  `xorm:"not null"` // token expire time

	Ip        string // token request ip
	UserAgent string // token user agent
	From      string // from type, web or app or other
}

// encode token value to hash
func (t *Token) EncodeValue() {
	t.Value = utils.Md5String(t.Value)
	t.Value = utils.Md5String(t.Value)
	t.Value = utils.Md5String(t.Value)
}

// save new token, only insert, not update
func SaveToken(t *Token) error {
	if _, err := app.Db.Insert(t); err != nil {
		log.Error("Db|SaveToken|%s", err.Error())
		return err
	}
	return nil
}
