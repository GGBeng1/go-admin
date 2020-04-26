package router

import (
	"hello/middleware"
	v1 "hello/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.JWTAuth())
	{
		userRouter.POST("userinfo", v1.TestToken)
	}
}
