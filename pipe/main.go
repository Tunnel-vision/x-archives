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
	server := &http.Server{
		Addr:    "0.0.0.0:" + model.Conf.Port,
		Handler: router,
	}
	if err := server.ListenAndServe();nil != err{
		println("listen and server failed: "+err.Error())
	}
}