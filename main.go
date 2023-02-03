package main

import (
	"hello/global"
	"hello/initall"
	"net/http"
	"time"
)

func main() {
	Router := initall.Routers()
	initall.Mysql()
	defer global.DB.Close()
	// 启动服务
	s := &http.Server{
		Addr:           ":8081",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
