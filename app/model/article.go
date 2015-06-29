package model

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
