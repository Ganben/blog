package action

import (
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/model"
)

func InitDbSchema(_ interface{}) *Result {
	app.Db.Sync2(new(model.User), new(model.Token), new(model.Article), new(model.Tag), new(model.Comment))
	return OkResult(nil)
}

func InitDbDefault(_ interface{}) *Result {
	// init administrator user
	user := model.NewDefaultUser()
	if err := model.SaveUser(user); err != nil {
		return ErrorResult(err)
	}

	// init welcome article
	article := model.NewDefaultArticle(user.Id)
	if err := model.SaveArticle(article); err != nil {
		return ErrorResult(err)
	}

	// init first comment
	comment := model.NewDefaultComment(article.Id)
	if err := model.SaveComment(comment); err != nil {
		return ErrorResult(err)
	}

	return OkResult(nil)
}
