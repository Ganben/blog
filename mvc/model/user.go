package model

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/entity"
	"time"
)

var (
	userData      map[int64]*entity.User = make(map[int64]*entity.User)
	userNameData  map[string]int64       = make(map[string]int64)
	userEmailData map[string]int64       = make(map[string]int64)

	tokenData map[string]*entity.Token = make(map[string]*entity.Token)
)

// load default user data
func loadUserData() {
	base.Storage.Walk(new(entity.User), func(v interface{}) {
		if u, ok := v.(*entity.User); ok {
			userData[u.Id] = u
			userNameData[u.Name] = u.Id
			userEmailData[u.Email] = u.Id
		}
	})
	base.Storage.Walk(new(entity.Token), func(v interface{}) {
		if t, ok := v.(*entity.Token); ok {
			if t.ExpireTime <= time.Now().Unix() {
				// remove expired token
				base.Storage.Remove(t)
				return
			}
			tokenData[t.Value] = t
		}
	})
}

// get token by value, ignore expiration
func GetToken(value string) *entity.Token {
	return tokenData[value]
}

// get token by value with expiration check
func GetValidToken(value string) *entity.Token {
	token := GetToken(value)
	if token != nil {
		if token.ExpireTime <= time.Now().Unix() {
			base.Storage.Remove(token)
			return nil
		}
	}
	return nil
}

// get user by token value with expiration check
func GetUserByTokenValue(value string) *entity.User {
	token := GetValidToken(value)
	if token == nil {
		return nil
	}
	return userData[token.UserId]
}

// get user by name
func GetUserByName(name string) *entity.User {
	if id := userNameData[name]; id > 0 {
		return userData[id]
	}
	return nil
}
