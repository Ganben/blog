package admin

import (
	"github.com/gofxh/blog/app/action"
	"github.com/gofxh/blog/app/model"
	"github.com/gofxh/blog/app/route/base"
	"github.com/lunny/tango"
	"net/http"
	"time"
)

type (
	// login router
	Login struct {
		base.AdminPageRouter
		base.BindRouter
	}

	Logout struct {
		tango.Ctx
	}
)

func (l *Login) Get() {
	if l.Cookie("x-token") != "" {
		l.Redirect("/admin/")
		return
	}
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

	// set cookie
	token := result.Data["token"].(*model.Token)
	l.Cookies().Set(&http.Cookie{
		Name:     "x-token",
		Value:    token.Value,
		Path:     "/",
		Expires:  time.Unix(token.ExpireTime, 0),
		MaxAge:   int(token.ExpireTime - time.Now().Unix()),
		HttpOnly: true,
	})

	// success, redirect
	l.Redirect("/admin/")
}

func (l *Logout) Get() {
	// remove token
	if token := l.Cookie("x-token"); token != "" {
		action.Call(action.UserLogout, token)
	}

	// remove cookie if exist
	l.Cookies().Set(&http.Cookie{
		Name:     "x-token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   0,
	})

	// redirect go login
	l.Redirect("/admin/login")
}
