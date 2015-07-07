package admin

import "github.com/gofxh/blog/app/route/base"

type Write struct {
	base.AdminPageRouter
	base.AuthRouter
	base.BindRouter
}

func (w *Write) Get() {
	w.Assign("Title", "New Article")
	w.Assign("IsArticlePage", true)
	w.MustRenderTheme(200, "write.tmpl")
}

func (w *Write) Post() {

}
