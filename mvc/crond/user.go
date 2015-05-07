package crond

import (
	"github.com/gofxh/blog/lib/base"
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
				base.Storage.Remove(t)
				m.Lock()
				delete(model.TokenData, t.Value)
				m.Unlock()
			}
		}
	}
}
