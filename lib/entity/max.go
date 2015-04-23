package entity

import "time"

type Maxer struct {
	ArticleId   int64 `json:"article_id"`
	ArticleStep int64 `json:"article_step"`

	CategoryId   int64 `json:"category_id"`
	CategoryStep int64 `json:"category_step"`

	CommentId   int64 `json:"comment_id"`
	CommentStep int64 `json:"comment_step"`

	UserId   int64 `json:"user_id"`
	UserStep int64 `json:"user_step"`
}

func (m *Maxer) SKey() string {
	return "max"
}

func (m *Maxer) Next(max, step int64) int64 {
	return max + time.Now().Unix()%step + 1
}
