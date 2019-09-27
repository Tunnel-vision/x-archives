package model

import "html/template"

type ThemeArticle struct {
	ID             uint64        `json:",omitempty"`
	Abstract       template.HTML `json:"abstract"`
	Author         *ThemeAuthor  `json:",omitempty"`
	CreatedAt      string        `json:",omitempty"`
	CreatedAtYear  string        `json:",omitempty"`
	CreatedAtMonth string        `json:",omitempty"`
	CreatedAtDay   string        `json:",omitempty"`
	Title          string        `json:"title"`
	Tags           []*ThemeTag   `json:"tags"`
	URL            string        `json:"url"`
	Topped         bool          `json:",omitempty"`
	ViewCount      int           `json:",omitempty"`
	CommentCount   int           `json:",omitempty"`
	ThumbnailURL   string        `json:",omitempty"`
	Content        template.HTML `json:",omitempty"`
	Editable       bool          `json:",omitempty"`
}

type ThemeTag struct {
	Title        string `json:"title"`
	URL          string `json:"url"`
	ArticleCount int    `json:",omitempty"`
}

type ThemeAuthor struct {
	Name         string
	AvatorURL    string
	URL          string
	ArticleCount int
}
