package model

import (
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/log"
	"github.com/gofxh/blog/app/utils"
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

	Nick   string // user nick name
	Url    string // personal address
	Bio    string // profile content
	Avatar string // avatar link, default is gravatar

	CreateTime int64 `xorm:"created"`

	Role   int8 // role status
	Status int8 // status, activated or not or deleted
}

// check password string
func (u *User) CheckPassword(str string) bool {
	pwd := utils.Sha256String(str + u.Salt)
	return pwd == u.Password
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
		Avatar: utils.GravatarLink("admin@example.com"),
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

// get user by column and value
func GetUserBy(col string, value interface{}) (*User, error) {
	u := new(User)
	if _, err := app.Db.Where(col+" = ?", value).Get(u); err != nil {
		log.Error("Db|GetUserBy|%s", err.Error())
		return nil, err
	}
	return u, nil
}

// get an user with name but without id
func GetUserByUniqueName(id int64, name string) (*User, error) {
	u := new(User)
	if _, err := app.Db.Where("id != ? AND name = ?", id, name).Get(u); err != nil {
		log.Error("Db|GetUserByUniqueName|%s", err.Error())
		return nil, err
	}
	if u.Id == 0 {
		return nil, nil
	}
	return u, nil
}

// get an user with email but without id
func GetUserByUniqueEmail(id int64, email string) (*User, error) {
	u := new(User)
	if _, err := app.Db.Where("id != ? AND email = ?", id, email).Get(u); err != nil {
		log.Error("Db|GetUserByUniqueEmail|%s", err.Error())
		return nil, err
	}
	if u.Id == 0 {
		return nil, nil
	}
	return u, nil
}

// update user profile
func UpdateUserProfile(u *User) (*User, error) {
	// update
	if _, err := app.Db.Cols("name,nick,email,url,bio").Where("id = ?", u.Id).Update(u); err != nil {
		log.Error("Db|UpdateUserProfile|%s", err.Error())
		return nil, err
	}
	// return updated user data
	user, err := GetUserBy("id", u.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// update user password
func UpdateUserPassword(u *User) error {
	u.Password, u.Salt = EncodePassword(u.Password)
	// update password
	if _, err := app.Db.Cols("password,salt").Where("id = ?", u.Id).Update(u); err != nil {
		log.Error("Db|UpdateUserPassword|%s", err.Error())
		return err
	}
	return nil
}
