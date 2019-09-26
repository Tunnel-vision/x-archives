package model

import "html/template"

type ThemeArticle struct {
	ID uint64
	Abstract template.HTML
	Author string
	CreateAt string
	CreateAtYear string
	CreateAtMonth string
	CreateAtDay string
	Title string
	Tags string
	URL string
	Topped bool
	ViewCount int
	CommentCount int
	ThumbnailURL string
	Content template.HTML
	Editable bool

}