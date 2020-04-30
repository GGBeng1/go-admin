package initall

import (
	"hello/middleware"
	"hello/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors()).Use(middleware.LoggerToFile())
	ApiGroup := Router.Group("")
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	return Router
}
