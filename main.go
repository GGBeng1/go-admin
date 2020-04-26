package main

import (
	"hello/initall"
	"net/http"
	"time"
)

func main() {
	// r := gin.Default()
	// r.Use(middleware.Cors())
	// r.GET("/users/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.String(200, "the user'id =", id)
	// })
	// r.GET("/user/info", func(c *gin.Context) {
	// 	c.String(200, c.Query("cc"))
	// })
	// r.GET("/user/as", func(c *gin.Context) {
	// 	c.String(200, c.Query("bb"))
	// })
	// r.POST("/register", func(c *gin.Context) {
	// 	var R request.RigisterStruct
	// 	err := c.ShouldBindJSON(&R)
	// 	if err != nil {
	// 		c.JSON(400, err.Error())
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, R)
	// })
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	Router := initall.Routers()
	// r.Run(":8081")
	s := &http.Server{
		Addr:           ":8081",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
