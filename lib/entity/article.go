package entity

const (
	ARTICLE_COMMENT_STATUS_ENABLE  int8 = 1
	ARTICLE_COMMENT_STATUS_EXPIRE  int8 = 3
	ARTICLE_COMMENT_STATUS_DISABLE int8 = 5

	ARTICLE_TOP_STATUS_ORDER    int8 = 1
	ARTICLE_TOP_STATUS_EXTERNAL int8 = 3

	ARTICLE_STATUS_PUBLIC  int8 = 1
	ARTICLE_STATUS_PRIVATE int8 = 3
	ARTICLE_STATUS_DELETED int8 = 5

	ARTICLE_TEXT_TYPE_MARKDOWN int8 = 1
	ARTICLE_TEXT_TYPE_HTML     int8 = 3
)

type Article struct {
	Id       int64  `json:"id"`
	AuthorId int64  `json:"author"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Text     string `json:"text"`
	TextType int8   `json:"text_type"`
	// brief text
	briefText string
	// rendered text
	renderedText string
	// category
	CategoryId int64 `json:"category_id"`
	category   *Category
	// tags
	Tags []string `json:"tags"`

	CreateTime int64 `json:"created"`
	UpdateTime int64 `json:"updated"`
	DeleteTime int64 `json:"deleted"`

	Status int8 `json:"status"`

	// CommentStatus set comment-enable
	CommentStatus int8
	// TopStatus set top link
	TopStatus int8

	// Comments
	CommentCount int `json:"comments"`
	comments     []*Comment

	// view count
	ViewCount int `json:"views"`
}
