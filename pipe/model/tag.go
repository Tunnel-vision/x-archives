package model

// Tag model.
type Tag struct {
	Model

	Title        string `gorm:"size:128" json:"title"`
	ArticleCount int    `json:"articleCount"`

	BlogID uint64 `sql:"index" json:"blogID"`
}