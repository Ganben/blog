package model

// article tag
type Tag struct {
    ArticleId int64 `xorm:"not null"`
    Tag string `xorm:"not null"`
    UserId int64  `xorm:"not null"`
}