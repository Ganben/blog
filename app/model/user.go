package model

import (
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/utils"
	"github.com/gofxh/note/lib/log"
	"time"
)

const (
	USER_ROLE_BLOCK  int8 = 9  // blocked user
	USER_ROLE_READER int8 = 19 // normal reader
	USER_ROLE_WRITER int8 = 29 // writer, can post article
	USER_ROLE_ADMIN  int8 = 99 // administrator, do any thing

	USER_STATUS_REMOVED       int8 = 7  // removed user
	USER_STATUS_NOT_ACTIVATED int8 = 17 // not-active
	USER_STATUS_ACTIVATE      int8 = 77 // active
)

// user struct
type User struct {
	Id       int64  // pk field
	Name     string `xorm:"unique(name) not null"`  // user name, unique, not null
	Password string `xorm:"not null"`               // user password
	Salt     string `xorm:"not null"`               // user password salt, use to encrypt some data by this user
	Email    string `xorm:"unique(email) not null"` // user email, unique, not null

	Nick string // user nick name
	Url  string // personal address
	Bio  string // profile content

	CreateTime int64 `xorm:"created"`

	Role   int8 // role status
	Status int8 // status, activated or not or deleted
}

// new default user
func NewDefaultUser() *User {
	user := &User{
		Name:   "admin",
		Email:  "admin@example.com",
		Nick:   "admin",
		Url:    "http://example.com",
		Bio:    "this is a default administrator user",
		Role:   USER_ROLE_ADMIN,
		Status: USER_STATUS_ACTIVATE,
	}
	user.Password, user.Salt = EncodePassword("123456789")
	return user
}

// encode password string, return encoded string and salt
func EncodePassword(pwd string) (string, string) {
	tmp := time.Now().Format(time.RFC3339)
	salt := utils.Md5String(tmp)[8:24]
	return utils.Sha256String(pwd + salt), salt
}

// save user
func SaveUser(u *User) error {
	var err error
	if u.Id > 0 {
		_, err = app.Db.Where("id = ?", u.Id).Update(u)
	} else {
		_, err = app.Db.Insert(u)
	}
	if err != nil {
		log.Error("Db|SaveUser|%s", err.Error())
		return err
	}
	return nil
}
