package main

import (
	"fmt"

	"github.com/bubbly-hb/blogSystem-gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run(":8080")
	fmt.Println("blogSystem-gin-vue")
}
