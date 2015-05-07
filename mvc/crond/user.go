package crond

import (
	"github.com/gofxh/blog/mvc/model"
	"sync"
	"time"
)

var (
	m sync.Mutex
)

func UserCron() {
	if len(model.TokenData) > 0 {
		for _, t := range model.TokenData {
			if t.ExpireTime <= time.Now().Unix() {
				model.RemoteToken(t)
			}
		}
	}
}
