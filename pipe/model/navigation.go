package model

// Navigation model.
type Navigation struct {
	Model

	Title      string `gorm:"size:128" json:"title"`
	URL        string `gorm:"size:255" json:"url"`
	IconURL    string `gorm:"size:255" json:"iconURL"`
	OpenMethod string `gorm:"size:32" json:"openMethod"`
	Number     int    `json:"number"` // for sorting

	BlogID uint64 `sql:"index" json:"blogID"`
}

// Navigation open methods.
const (
	NavigationOpenMethodBlank = "_blank"
	NavigationOpenMethodSelf  = "_self"
)
