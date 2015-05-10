package api

import (
	"errors"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/mvc/action"
	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
	"github.com/tango-contrib/xsrf"
)

type ArticleController struct {
	tango.Ctx
	xsrf.NoCheck
	binding.Binder
}

func (a *ArticleController) Post() {
	var form action.ArticleForm
	if e := a.Bind(&form); e.Len() > 0 {
		a.ServeJson(core.NewErrorResult(errors.New(e[0].Error())))
		return
	}
	result := base.Action.Call(action.SaveArticle, &form)
	a.ServeJson(result)
}
