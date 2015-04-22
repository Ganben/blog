package entity

type Maxer struct {
	ArticleId   int64 `json:"article_id"`
	articleStep int64

	CategoryId   int64 `json:"category_id"`
	categoryStep int64

	CommentId   int64 `json:"comment_id"`
	commentStep int64

	UserId   int64 `json:"user_id"`
	userStep int64
}
