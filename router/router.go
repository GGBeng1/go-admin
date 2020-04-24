package initrouter

import (
	"hello/middleware"
	v1 "hello/v1"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/register", v1.Register)
	r.POST("/login", v1.Login)
	r.GET("/captcha", v1.GetCaptcha)
	r.GET("/captcha/:captcha", v1.GetCaptchaPng)
	return r
}
