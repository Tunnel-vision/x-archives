package service

import (
	"fmt"
	"time"
	"x-archives/pipe/model"
	"log"
	"os"
	"github.com/gorm"
)
var db *gorm.DB
var logger =  log.New(os.Stdout,"pipe",log.LstdFlags)
var userSQLite bool


func Connect_db()  {
	db, err := gorm.Open("sqlite3", model.Conf.SQLite)
	if nil != err{
		log.Fatal("err:"+err.Error())
	}
	userSQLite = false
	if err := db.AutoMigrate(model.Models...).Error;nil != err{
		fmt.Println("0"+err.Error())
		logger.Output(2,"auto migrate tables false")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(5*time.Minute)
	db.LogMode(false)

}

func Disconect_db()  {
	if err := db.Close(); nil != err{
		logger.Output(2,err.Error())
	}
}

func Datebase() string {
	if userSQLite{
		return "SQLite"
	}
	return "MYSQL"
}