package model

// Setting model.
type Setting struct {
	Model

	Category string `sql:"index" gorm:"size:32" json:"category"`
	Name     string `sql:"index" gorm:"size:64" json:"name"`
	Value    string `gorm:"type:text" json:"value"`

	BlogID uint64 `sql:"index" json:"blogID"`
}
