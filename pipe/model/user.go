package model

import (
	"github.com/b3log/pipe/model"
	"x-archives/pipe/util"
)

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
	UserRolePlatformAdmin // 平台管理员 1
	UserRoleBlogAdmin // 博客管理员 2
	UserRoleBlogUser  // 博客使用者 3
)

func (u *User) AvatarURLWithSize(size int) string {
	return util.ImageSize(u.AvatarURL, size, size)
}
