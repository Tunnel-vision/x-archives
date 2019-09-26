package model

import (
	"time"
)

type Article struct {
	Model
	Title string
	AuthorID uint64 `json:"author_id"`
	Abstract string
	Tags string
	Content string
	Path string
	Status string
	Topped bool
	Commentable bool
	ViewCout int
	CommentCount int
	IP string
	userAgent string
	PushedAt time.Time
	BlogID uint64
}

const ArticleStatusOK  = iota