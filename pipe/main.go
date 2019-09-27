package main

import (
	"x-archives/pipe/controller"
	"x-archives/pipe/model"
	"x-archives/pipe/service"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

func init()  {
	// init db
	model.LoadConf()
}

// Entry point
func main()  {
	// connect_db
	service.Connect_db()
	// service run
	router := controller.MapRoutes()
	// 初始化 http 服务器
	server := &http.Server{
		Addr:    "0.0.0.0:" + model.Conf.Port,
		// 一个要实现 server_http 方法的 结构体
		Handler: router,
	}
	if err := server.ListenAndServe();nil != err{
		println("listen and server failed: "+err.Error())
	}
}