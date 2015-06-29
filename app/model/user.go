package model

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
