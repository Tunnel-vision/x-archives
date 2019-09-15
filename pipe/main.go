package main

import (
	"x-archives/pipe/model"
	"x-archives/pipe/service"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init()  {
	// init db
	model.LoadConf()
}


func main()  {
	// connect_db
	service.Connect_db()
	// service run

}