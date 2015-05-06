package controller

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/mvc/controller/api"
	"github.com/lunny/tango"
)

func Init(_ interface{}) *core.ActionResult {
	/*adminGroup := tango.NewGroup()
	adminGroup.Get("/login", new(AdminLoginController))
	adminGroup.Post("/login", new(AdminLoginController))

	base.Server.Group("/admin", adminGroup)*/

	apiGroup := tango.NewGroup()
	apiGroup.Post("/user/login", new(api.LoginController))
	base.Server.Group("/api", apiGroup)

	base.Server.Get("/", new(IndexController))
	return core.NewOKActionResult(nil)
}
