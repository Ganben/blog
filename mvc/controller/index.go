package controller

import "github.com/gofxh/blog/mvc/helper"

type IndexController struct {
	helper.ThemeController
}

func (idxC *IndexController) Get() {
	idxC.Render("index.html")
}
