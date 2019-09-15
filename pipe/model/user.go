package model

import "github.com/b3log/pipe/model"

type User struct {
	model.Model
	Name              string `gorm:"size:32" json:"name"`
	NickName          string `gorm:"size:32" json:"nickname"`
	AvatarURL         string `gorm:"size:255" json:"avatarURL"`
	B3Key             string `gorm:"size:32" json:"b3key"`
	Locale            string `gorm:"size:32" json:"locale"`
	TotalArticleCount int    `json:"totalArticleCount"`
	GithubId          string `gorm:"255" json:"githubId"`
}

// user role
const (
	UserRoleNoLogin = iota
	UserRolePlatformAdmin
	UserRoleBlogAdmin
	UserRoleBlogUser
)

func (u *User) AvatarURLWithSize (size int) string {
	return "xxx"
}