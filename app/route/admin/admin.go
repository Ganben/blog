package admin

import (
	"github.com/gofxh/blog/app/route/base"
	"time"
)

type Admin struct {
	base.PageRouter
}

func (a *Admin) Get() {
	a.Assign("Title", time.Now().Format(time.RFC1123))
	a.MustRenderTheme(200, "index.tmpl")
}
