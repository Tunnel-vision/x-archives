package model


type Archive struct {
	Model
	Year string `gorm:"size:4"json:"year"`
	Month string `gorm:"size:2"json:"month"`
	ArticleCount int `json:"articleCount"`
	BlogID uint64 `sql:"sql" json:"blogID"`
}