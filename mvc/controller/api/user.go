package api

import (
	"errors"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/entity"
	"github.com/gofxh/blog/mvc/action"
	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
	"github.com/tango-contrib/xsrf"
	"net/http"
	"time"
)

// login controller
type LoginController struct {
	tango.Ctx
	xsrf.NoCheck
	binding.Binder
}

// login controller post method
func (l *LoginController) Post() {
	var form action.LoginForm
	if e := l.Bind(&form); e.Len() > 0 {
		l.ServeJson(core.NewErrorResult(errors.New(e[0].Error())))
		return
	}
	form.Ip = l.Req().RemoteAddr
	form.UserAgent = l.Req().UserAgent()
	form.Expire = 3600 * 24 * 7
	result := base.Action.Call(action.Login, &form)

	redirect := l.Req().FormValue("redirect")
	if result.Meta.Status && redirect != "" {
		tk := result.Data["token"].(*entity.Token)
		l.Cookies().Set(&http.Cookie{
			Name:     "token",
			Value:    tk.Value,
			Path:     "/",
			Expires:  time.Unix(tk.ExpireTime, 0),
			MaxAge:   int(tk.ExpireTime - time.Now().Unix()),
			HttpOnly: true,
		})
		l.Redirect(redirect, 302)
		return
	}
	l.ServeJson(result)
}
