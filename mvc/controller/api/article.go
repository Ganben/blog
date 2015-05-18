package api

import (
	"errors"
	"fmt"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/mvc/action"
	"github.com/gofxh/blog/mvc/helper"
	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
	"github.com/tango-contrib/xsrf"
)

type ArticleController struct {
	tango.Ctx
	xsrf.NoCheck
	binding.Binder
	helper.AuthController
}

func (a *ArticleController) Post() {
	form := action.ArticleForm{
		AuthorId: a.AuthUser.Id,
	}
	if e := a.Bind(&form); e.Len() > 0 {
		a.ServeJson(core.NewErrorResult(errors.New(e[0].Error())))
		return
	}
	fmt.Println(form)
	result := base.Action.Call(action.SaveArticle, &form)
	a.ServeJson(result)
}
