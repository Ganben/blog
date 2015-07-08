package model

import (
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/log"
)

// article tag
type Tag struct {
	ArticleId int64  `xorm:"not null"`
	Tag       string `xorm:"not null"`
	UserId    int64  `xorm:"not null"`
}

// remove tags belongs to an article
func RemoveArticleTags(articleId int64) error {
	if _, err := app.Db.Where("article_id = ?", articleId).Delete(new(Tag)); err != nil {
		log.Error("Db|RemoveArticleTags|%d|%s", articleId, err.Error())
		return err
	}
	return nil
}

// save tags with article and user id
func SaveTags(articleId, userId int64, tags []string) error {
	var err error
	for _, t := range tags {
		if len(t) == 0 {
			continue
		}
		tag := &Tag{articleId, t, userId}
		if _, err = app.Db.Insert(tag); err != nil {
			log.Error("Db|SaveTags|%d|%s", articleId, err.Error())
			return err
		}
	}
	return nil
}
