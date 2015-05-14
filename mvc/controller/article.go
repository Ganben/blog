package controller

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/mvc/helper"
	"html/template"
)

type ArticleWriteController struct {
	helper.ThemeController
	helper.AuthController
}

func (awc *ArticleWriteController) GetAuthFailRedirect() string {
	return "/"
}

func (awc *ArticleWriteController) Get() {
	awc.Assign("AuthUser", awc.AuthUser)
	awc.Assign("IsSigned", awc.AuthUser != nil)

	// render rich text editor
	richEditorForm := &helper.RichEditorForm{
		Render:   awc.ThemeController,
		Template: "editor.html",
	}
	result := base.Action.Call(helper.RichEditor, richEditorForm)
	if !result.Meta.Status {
		panic(result.Meta.ErrorMessage)
	}
	awc.Assign("RichEditor", template.HTML(string(result.Data["bytes"].([]byte))))

	// render template
	awc.Render("article_write.html")
	println("get page")
}

func (awc *ArticleWriteController) Post() {
	println("get page post")
}
