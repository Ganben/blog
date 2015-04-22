package entity

type Token struct {
	Value  string `json:"token"`
	UserId int64  `json:"user_id"`

	CreateTime int64 `json:"created"`
	ExpireTime int64 `json:"expired"`

	UserIp    string `json:"user_ip"`
	UserAgent string `json:"user_agent"`
}
