package entity

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"
)

type Token struct {
	Value  string `json:"token"`
	UserId int64  `json:"user_id"`

	CreateTime int64 `json:"created"`
	ExpireTime int64 `json:"expired"`

	UserIp    string `json:"user_ip"`
	UserAgent string `json:"user_agent"`
}

func (t *Token) SKey() string {
	return fmt.Sprintf("token/%s.json", t.Value)
}

// generate token value
func GenerateTokenValue(uid int64, ip, agent string) string {
	content := fmt.Sprintf("%d_%s_%s_%d", uid, ip, agent, time.Now().Unix())
	s := sha1.New()
	s.Write([]byte(content))
	return hex.EncodeToString(s.Sum(nil))
}
