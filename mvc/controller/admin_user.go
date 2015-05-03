package controller

import (
	"github.com/gofxh/blog/mvc/helper"
	"github.com/lunny/tango"
)

type AdminLoginController struct {
	tango.Ctx
	helper.ThemeController
}

func (alc *AdminLoginController) Get() {
	alc.Render("login.html")
}

func (alc *AdminLoginController) Post() {
	println("admin login post")
}
