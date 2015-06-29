package model

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
