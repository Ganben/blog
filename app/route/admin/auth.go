package admin

import (
	"github.com/gofxh/blog/app/route/base"
)

type (
	Login struct {
		base.AdminPageRouter
		base.BindRouter
	}
	LoginForm struct {
		Username string `form:"username" binding:"Required"`
		Password string `form:"password" binding:"Required;AlphaDash"`
		Remember int    `form:"remember"`
	}
)

func (l *Login) Get() {
	l.Assign("Title", "Login")
	l.MustRenderTheme(200, "login.tmpl")
}

func (l *Login) Post() {
	l.Assign("Title", "Login")

	// validate form
	form := &LoginForm{}
	if err := l.BindAndValidate(form); err != nil {
		l.Assign("Error", err.Error())
		l.MustRenderTheme(200, "login.tmpl")
		return
	}
}
