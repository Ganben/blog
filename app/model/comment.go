package model

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
