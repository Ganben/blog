package controller

import (
	"github.com/gofxh/blog/lib/base"
	// "github.com/lunny/tango"
)

func Init() {
	/*adminGroup := tango.NewGroup()
	adminGroup.Get("/login", new(AdminLoginController))
	adminGroup.Post("/login", new(AdminLoginController))

	base.Server.Group("/admin", adminGroup)*/
	base.Server.Get("/", new(IndexController))
}
