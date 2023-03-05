package routes

import (
	v1 "github.com/bubbly-hb/blogSystem-gin/api/v1"
	"github.com/bubbly-hb/blogSystem-gin/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("user/get", v1.GetUsers)
		router.PUT("user/:id", v1.UpdateUser)
		router.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("category/get", v1.GetCategorys)
		router.PUT("category/:id", v1.UpdateCategory)
		router.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块的路由接口
		router.POST("article/add", v1.AddArticle)
		router.GET("article/get", v1.GetArticles)             // 获取文章列表
		router.GET("article/info/:id", v1.GetArticle)         // 获取单篇文章
		router.GET("article/list/", v1.GetArticlesByCategory) // 获取分类下的所有文章
		router.PUT("article/:id", v1.UpdateArticle)
		router.DELETE("article/:id", v1.DeleteArticle)
	}

	r.Run(utils.HttpPort)
}
