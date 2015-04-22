package entity

type Category struct {
	Id          int64  `json:"id"`
	AuthorId    int64  `json:"author"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"desc"`

	// article count
	ArticleCount int `json:"article_count"`
}
