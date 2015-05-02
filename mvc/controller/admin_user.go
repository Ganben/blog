package controller

import (
	"github.com/lunny/tango"
	"github.com/tango-contrib/renders"
)

type AdminLoginController struct {
	tango.Ctx
	renders.Renderer
}

func (alc *AdminLoginController) Get() {
	alc.Render("admin/login.html")
}

func (alc *AdminLoginController) Post() {
	println("admin login post")
}
