package controller

import (
	"github.com/gin-gonic/gin"
	"x-archives/pipe/util"
)

func MapRoutes() *gin.Engine  {
	ret := gin.New()

	ret.Use(gin.Recovery())

	// cookies 管理

	// 1 层命名空间
	api := ret.Group(util.PathAPI)
	api.GET("/oauth/github/redirect",redirectGithubLoginAction)
	return ret
}