package model

// Category model.
type Category struct {
	Model

	Title           string `gorm:"size:128" json:"title"`
	Path            string `gorm:"size:255" json:"path"`
	Description     string `gorm:"size:255" json:"description"`
	MetaKeywords    string `gorm:"size:255" json:"metaKeywords"`
	MetaDescription string `gorm:"type:text" json:"metaDescription"`
	Tags            string `gorm:"type:text" json:"tags"`
	Number          int    `json:"number"` // for sorting

	BlogID uint64 `sql:"index" json:"blogID"`
}
