package action

import (
	"errors"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/entity"
	"github.com/gofxh/blog/mvc/model"
	"strings"
)

var (
	ARTICLE_BAD_DATA = errors.New("bad data")
)

type ArticleForm struct {
	Id       int64 // if id > 0, try update
	AuthorId int64
	Title    string
	Content  string
	Type     string
	Tags     string
	Slug     string
}

func SaveArticle(v interface{}) *core.ActionResult {
	form, ok := v.(*ArticleForm)
	if !ok {
		return core.NewErrorResult(ARTICLE_BAD_DATA)
	}
	if form.Id > 0 {
		return saveModifiedArticle(form)
	}
	// assign default data
	a := new(entity.Article)
	a.AuthorId = form.AuthorId
	a.Title = form.Title
	a.Slug = form.Slug
	a.Text = form.Content
	if form.Type == "markdown" {
		a.TextType = entity.ARTICLE_TEXT_TYPE_MARKDOWN
	}
	a.CategoryId = 0
	form.Tags = strings.Replace(form.Tags, "，", ",", -1) // replace chinese comma
	a.Tags = strings.Split(form.Tags, ",")
	a.Status = entity.ARTICLE_STATUS_PUBLIC
	a.CommentStatus = entity.ARTICLE_STATUS_PUBLIC
	a.TopStatus = 0
	a.CommentCount = 0
	a.ViewCount = 1
	model.SaveArticle(a)

	return core.NewOKActionResult(core.AData{
		"article": a,
	})
}

func saveModifiedArticle(form *ArticleForm) *core.ActionResult {
	// assign form data
	a := new(entity.Article)
	a.Title = form.Title
	a.Slug = form.Slug
	a.Text = form.Content
	if form.Type == "markdown" {
		a.TextType = entity.ARTICLE_TEXT_TYPE_MARKDOWN
	}
	a.CategoryId = 0
	form.Tags = strings.Replace(form.Tags, "，", ",", -1) // replace chinese comma
	a.Tags = strings.Split(form.Tags, ",")
	a.Status = entity.ARTICLE_STATUS_PUBLIC
	a.CommentStatus = entity.ARTICLE_STATUS_PUBLIC
	a.TopStatus = 0
	model.UpdateArticle(a)

	return core.NewOKActionResult(core.AData{
		"article": a,
	})
}
