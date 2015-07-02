package admin

import (
	"github.com/gofxh/blog/app/route/base"
)

type Admin struct {
	base.AdminPageRouter
	base.AuthRouter
}

func (a *Admin) Get() {
	a.Assign("Title", "Dashboard")
	a.MustRenderTheme(200, "index.tmpl")
}
