package admin

import (
	"github.com/gofxh/blog/app/action"
	"github.com/gofxh/blog/app/route/base"
)

type (
	// login router
	Login struct {
		base.AdminPageRouter
		base.BindRouter
	}
)

func (l *Login) Get() {
	l.Assign("Title", "Login")
	l.MustRenderTheme(200, "login.tmpl")
}

func (l *Login) Post() {
	l.Assign("Title", "Login")

	// validate form
	form := &action.LoginForm{}
	if err := l.BindAndValidate(form); err != nil {
		l.Assign("Error", err.Error())
		l.MustRenderTheme(200, "login.tmpl")
		return
	}
	// call UserLogin action
	result := action.Call(action.UserLogin, form)
	if !result.Status {
		l.Assign("Error", result.Error)
		l.MustRenderTheme(200, "login.tmpl")
		return
	}

	// success, redirect
	l.Redirect("/admin/")
}
