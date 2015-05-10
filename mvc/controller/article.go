package controller

import "github.com/gofxh/blog/mvc/helper"

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
	awc.Render("article_write.html")
	println("get page")
}

func (awc *ArticleWriteController) Post() {
	println("get page post")
}
