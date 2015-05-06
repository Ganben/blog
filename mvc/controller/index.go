package controller

import "github.com/gofxh/blog/mvc/helper"

type IndexController struct {
	helper.AuthController
	helper.ThemeController
}

func (idxC *IndexController) Get() {
	idxC.Render("index.html")
}
