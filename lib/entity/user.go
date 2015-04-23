package entity

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const (
	USER_STATUS_ACTIVE    int8 = 1 // active
	USER_STATUS_NO_ACTIVE int8 = 3 // not active
	USER_STATUS_BLOCK     int8 = 5 // blocked, not allow to login
	USER_STATUS_DELETED   int8 = 7 // deleted

	USER_ROLE_ADMIN  int8 = 1
	USER_ROLE_WRITER int8 = 3
	USER_ROLE_READER int8 = 5
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Salt     string `json:"salt"`

	Nick  string `json:"nick"`
	Email string `json:"email"`
	Url   string `json:"url"`
	Bio   string `json:"bio"`

	CreateTime int64 `json:"created"`
	LoginTime  int64 `json:"logged"`
	DeleteTime int64 `json:"deleted"`

	Status int8 `json:"status"`
	Role   int8 `json:"role"`
}

func (s *User) SKey() string {
	return fmt.Sprintf("user/user_%d", s.Id)
}

func GenerateUserPassword(password string) (string, string) {
	// md5 to create salt
	t := strconv.FormatInt(time.Now().UnixNano(), 10)
	m := md5.New()
	m.Write([]byte(t))
	salt := hex.EncodeToString(m.Sum(nil))[:16]

	// sha256 to real password
	s := sha256.New()
	s.Write([]byte(password))
	s.Write([]byte(salt))
	return hex.EncodeToString(s.Sum(nil)), salt
}
