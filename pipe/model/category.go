package model


type Category struct {
	Model
	Title string
	Path string
	Description string
	MetaKeywords string
	MetaDescription string
	Tags string
	Number int
	BlogID uint64

}