package entity

const (
	COMMENT_REL_TYPE_ARTICLE int8 = 1

	COMMENT_STATUS_WAIT     int8 = 1 // wait to approve
	COMMENT_STATUS_APPROVED int8 = 3 // approved
	COMMENT_STATUS_SPAM     int8 = 5 // spam
	COMMENT_STATUS_DELETE   int8 = 7 // deleted
)

type Comment struct {
	Id       int64  `json:"id"`
	AuthorId int64  `json:"author"`
	ParentId int64  `json:"parent_id"`
	Email    string `json:"email"`
	Url      string `json:"url"`

	// comment content
	Content string `json:"content"`

	// status
	Status int8 `json:"status"`

	// relative content
	RelType int8  `json:"rel_type"`
	RelId   int64 `json:"rel_id"`

	// remote info
	UserIp    string `json:"user_ip"`
	UserAgent string `json:"user_agent"`
}
