package api

import (
	"errors"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/entity"
	"github.com/gofxh/blog/mvc/action"
	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
	"net/http"
	"time"
)

// login controller
type LoginController struct {
	tango.Ctx
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
	if result.Meta.Status {
		tk := result.Data["token"].(*entity.Token)
		l.Cookies().Set(&http.Cookie{
			Name:   "token",
			Value:  tk.Value,
			MaxAge: int(tk.ExpireTime - time.Now().Unix()),
		})
	}
	l.ServeJson(result)
}
