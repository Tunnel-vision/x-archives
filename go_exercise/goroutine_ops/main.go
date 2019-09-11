package main

import (
"net/http"
_ "net/http/pprof"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitAdmin() {
	adminRouter := http.DefaultServeMux
	adminRouter.Handle("/metrics", promhttp.Handler())

	adminServer := &http.Server{
		Addr:           ":8081",
		Handler:        adminRouter,
	}

	go func() {
		if err := adminServer.ListenAndServe(); err != nil {
			println("ListenAndServe admin: ", err.Error())
		}
	}()
}

func main()  {
	InitAdmin()
	time.Sleep(1000*time.Second)
}