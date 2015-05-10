package action

import (
	"errors"
	"github.com/gofxh/blog/lib/core"
)

var (
	ARTICLE_BAD_DATA = errors.New("bad data")
)

type ArticleForm struct {
}

func SaveArticle(v interface{}) *core.ActionResult {
	form, ok := v.(*ArticleForm)
	if !ok {
		return core.NewErrorResult(ARTICLE_BAD_DATA)
	}
	return core.NewOKActionResult(core.AData{
		"article": form,
	})
}
