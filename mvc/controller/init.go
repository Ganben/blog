package controller

import (
	"github.com/gofxh/blog/lib/base"
	// "github.com/lunny/tango"
	"github.com/gofxh/blog/lib/core"
)

func Init(_ interface{}) *core.ActionResult {
	/*adminGroup := tango.NewGroup()
	adminGroup.Get("/login", new(AdminLoginController))
	adminGroup.Post("/login", new(AdminLoginController))

	base.Server.Group("/admin", adminGroup)*/
	base.Server.Get("/", new(IndexController))
	return core.NewOKActionResult(nil)
}
