package model

import (
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/log"
)

const (
	COMMENT_STATUS_REMOVED  int8 = 15 // removed comment
	COMMENT_STATUS_SPAM     int8 = 25 // spam
	COMMENT_STATUS_APPROVED int8 = 55 // approved
)

type Comment struct {
	Id         int64
	Author     string `xorm:"not null"` // comment author
	Email      string `xorm:"not null"` // comment author email
	Url        string
	Content    string `xorm:"not null"` // comment content
	ArticleId  int64  // article id of this comment
	CreateTime int64  `xorm:"created"` // create time
	Status     int8   // status, approved or spam or removed
	ParentId   int64  // comment parent, as comment's reply
}

func NewDefaultComment(articleId int64) *Comment {
	return &Comment{
		Author:    "human",
		Email:     "human@example.com",
		Url:       "http://example.com",
		Content:   "this is a comment for default article",
		ArticleId: articleId,
		Status:    COMMENT_STATUS_APPROVED,
		ParentId:  0,
	}
}

// save comment
func SaveComment(u *Comment) error {
	var err error
	if u.Id > 0 {
		_, err = app.Db.Where("id = ?", u.Id).Update(u)
	} else {
		_, err = app.Db.Insert(u)
	}
	if err != nil {
		log.Error("Db|SaveComment|%s", err.Error())
		return err
	}
	return nil
}
