package controller

import "github.com/gofxh/blog/mvc/helper"

type IndexController struct {
	helper.AuthController
	helper.ThemeController
}

func (idxC *IndexController) Get() {
	idxC.Assign("IsSigned", idxC.AuthUser != nil)
	idxC.Render("index.html")
}
