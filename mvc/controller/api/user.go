package api

import (
	"errors"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/mvc/action"
	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
)

// login controller
type LoginController struct {
	tango.Context
	binding.Binder
}

// login controller post method
func (l *LoginController) Post() {
	var form action.LoginForm
	if e := l.Bind(&form); e.Len() > 0 {
		l.ServeJson(core.NewErrorResult(errors.New(e[0].Error())))
		return
	}
	result := base.Action.Call(action.Login, &form)
	l.ServeJson(result)
}
