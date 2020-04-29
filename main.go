package main

import (
	"hello/initall"
	"net/http"
	"time"
)

func main() {
	Router := initall.Routers()
	initall.Mysql()
	s := &http.Server{
		Addr:           ":8081",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
