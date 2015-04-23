package model

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/entity"
)

var (
	userData      map[int64]*entity.User = make(map[int64]*entity.User)
	userNameData  map[string]int64       = make(map[string]int64)
	userEmailData map[string]int64       = make(map[string]int64)
)

func loadUserData() {
	base.Storage.Walk(new(entity.User), func(v interface{}) {
		if u, ok := v.(*entity.User); ok {
			userData[u.Id] = u
			userNameData[u.Name] = u.Id
			userEmailData[u.Email] = u.Id
		}
	})
}
