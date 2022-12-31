package router

import (
	"github.com/bubbly-hb/blogSystem-gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("v1")
	{
		api.POST("/user/register", controller.Register)
	}
	return r
}