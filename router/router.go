package router

import (
	"github.com/bubbly-hb/blogSystem-gin/controller"
	"github.com/bubbly-hb/blogSystem-gin/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("v1")
	{
		api.POST("/user/register", controller.Register)
		api.POST("/user/login", controller.Login)
		api.GET("/user/info", middleware.AuthMiddleware(), controller.Info)
	}
	return r
}
