package model

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/entity"
	"sync"
	"time"
)

var (
	UserData  map[int64]*entity.User   = make(map[int64]*entity.User)
	TokenData map[string]*entity.Token = make(map[string]*entity.Token)

	userNameData  map[string]int64 = make(map[string]int64)
	userEmailData map[string]int64 = make(map[string]int64)

	m sync.Mutex
)

// load default user data
func loadUserData() {
	base.Storage.Walk(new(entity.User), func(v interface{}) {
		if u, ok := v.(*entity.User); ok {
			UserData[u.Id] = u
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
			TokenData[t.Value] = t
		}
	})
}

// get token by value, ignore expiration
func GetToken(value string) *entity.Token {
	return TokenData[value]
}

// get token by value with expiration check
func GetValidToken(value string) *entity.Token {
	token := GetToken(value)
	if token != nil {
		if token.ExpireTime <= time.Now().Unix() {
			RemoteToken(token)
			return nil
		}
	}
	return token
}

// get user by token value with expiration check
func GetUserByTokenValue(value string) (*entity.Token, *entity.User) {
	token := GetValidToken(value)
	if token == nil {
		return nil, nil
	}
	return token, UserData[token.UserId]
}

// get user by name
func GetUserByName(name string) *entity.User {
	if id := userNameData[name]; id > 0 {
		return UserData[id]
	}
	return nil
}

// create new token
func CreateToken(uid int64, ip, agent string, expire int64) *entity.Token {
	token := &entity.Token{
		Value:      entity.GenerateTokenValue(uid, ip, agent),
		UserId:     uid,
		CreateTime: time.Now().Unix(),
		ExpireTime: time.Now().Unix() + expire,
		UserIp:     ip,
		UserAgent:  agent,
	}
	base.Storage.Save(token)
	TokenData[token.Value] = token
	return token
}

// extend token's expire time
func ExtendToken(t *entity.Token) {
	t.ExpireTime = time.Now().Unix() + (t.ExpireTime - t.CreateTime)
	base.Storage.Save(t)
	TokenData[t.Value] = t
}

// remove token
func RemoteToken(t *entity.Token) {
	base.Storage.Remove(t)
	m.Lock()
	delete(TokenData, t.Value)
	m.Unlock()
}
