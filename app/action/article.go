package action

import (
	"github.com/gofxh/blog/app/model"
	"github.com/gofxh/blog/app/utils"
	"strings"
	"time"
)

type ArticleForm struct {
	Id      int64  `form:"id"`
	Title   string `form:"title" binding:"Required"`
	Content string `form:"content" binding:"Required"`
	Tag     string `form:"tag"`
	Status  string `form:"status"`
	UserId  int64  `form:"-"`
}

func ArticleSave(v interface{}) *Result {
	form, ok := v.(*ArticleForm)
	if !ok {
		return ErrorResult(paramTypeError(new(ArticleForm)))
	}
	article := &model.Article{
		Title:         form.Title,
		Content:       form.Content,
		UserId:        form.UserId,
		ContentType:   model.ARTICLE_CNT_TYPE_MARKDOWN,
		CommentStatus: model.ARTICLE_CMT_STATUS_OPEN,
		HitCount:      1,
		CommentCount:  0,
		TagString:     strings.TrimSpace(strings.Replace(form.Tag, "，", ",", -1)),
	}
	if form.Status == "publish" {
		article.Status = model.ARTICLE_STATUS_PUBLISH
	} else {
		article.Status = model.ARTICLE_STATUS_DRAFT
	}

	// generate link
	article.Link = generateLink()

	// save article
	if err := model.SaveArticle(article); err != nil {
		return ErrorResult(err)
	}

	// means update, remove old tags
	if form.Id > 0 {
		if err := model.RemoveArticleTags(form.Id); err != nil {
			return ErrorResult(err)
		}
	}

	// save tags
	if err := model.SaveTags(article.Id, article.UserId, strings.Split(article.TagString, ",")); err != nil {
		return ErrorResult(err)
	}

	return OkResult(nil)
}

func generateLink() string {
	link := utils.Md5String(time.Now().Format(time.RFC3339Nano))[8:16]
	// read database, make sure link is unique
	if article, _ := model.GetArticleByLink(link); article != nil {
		return generateLink()
	}
	return link
}
