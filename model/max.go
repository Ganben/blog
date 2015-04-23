package model

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/entity"
)

func NewMaxer() *entity.Maxer {
	m := &entity.Maxer{
		ArticleStep:  3,
		CommentStep:  3,
		CategoryStep: 3,
		UserStep:     3,
	}

	if base.Storage.Exist(m) {
		base.Storage.Read(m)
	}

	return m
}
