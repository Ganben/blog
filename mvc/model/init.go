package model

import "github.com/gofxh/blog/lib/core"

func Init(_ interface{}) *core.ActionResult {
	loadUserData()
	loadArticleData()

	return core.NewOKActionResult(core.AData{
		"users":    UserData,
		"articles": ArticleData,
	})
}
