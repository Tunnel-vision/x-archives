package model

type Tag struct {
	Model
	Title string
	ArticleCount int
	BlogID uint64
}
