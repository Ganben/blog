package model

import (
	"coding.net/fuxiaohei/goplay.git/log"
	"github.com/gofxh/blog/app"
)

const (
	ARTICLE_CNT_TYPE_MARKDOWN int8 = 12 // markdown content
	ARTICLE_CNT_TYPE_HTML     int8 = 22 // html content

	ARTICLE_STATUS_REMOVED int8 = 13 // article removed
	ARTICLE_STATUS_DRAFT   int8 = 23 // draft
	ARTICLE_STATUS_PUBLISH int8 = 33 // published

	ARTICLE_CMT_STATUS_CLOSED int8 = 14 // close comment
	ARTICLE_CMT_STATUS_MONTH  int8 = 24 // close after month
	ARTICLE_CMT_STATUS_OPEN   int8 = 34 // open forever
)

// article struct
type Article struct {
	Id          int64
	Title       string `xorm:"not null"`              // article title
	Link        string `xorm:"not null unique(link)"` // article unique link
	Content     string `xorm:"not null"`              // article content
	UserId      int64  `xorm:"not null"`              // article author id
	ContentType int8   `xorm:"not null"`              // article content type, markdown or html or other

	CreateTime int64 `xorm:"created"` // create time
	UpdateTime int64 `xorm:"updated"` // update time
	RemoveTime int64

	Status        int8 // article status, publish or draft or removed
	CommentStatus int8 // comment status, open or closed or others
	HitCount      int  // hit count
	CommentCount  int  // comment count

	Tags      []*Tag `xorm:"-"` // tag data
	TagString string `xorm:"-"` // tag data as string
}

// new default article
func NewDefaultArticle(uid int64) *Article {
	article := &Article{
		Title: "welcome to Gofxh.Blog",
		Link:  "welcome",
		Content: `## 欢迎使用

欢迎使用 Gofxh.Blog 。如果您看到这篇文章，表示您的Blog已经在SAE安装成功。您现在可以编辑或者删除掉这篇文章，然后开始您的博客之旅！`,
		UserId:        uid,
		ContentType:   ARTICLE_CNT_TYPE_MARKDOWN,
		Status:        ARTICLE_STATUS_PUBLISH,
		CommentStatus: ARTICLE_CMT_STATUS_OPEN,
		HitCount:      1,
		CommentCount:  0,
	}
	return article
}

// save article
func SaveArticle(a *Article) error {
	var err error
	if a.Id > 0 {
		_, err = app.Db.Where("id = ?", a.Id).Update(a)
	} else {
		_, err = app.Db.Insert(a)
	}
	if err != nil {
		log.Error("Db|SaveArticle|%s", err.Error())
		return err
	}
	return nil
}
