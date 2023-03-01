package main

import (
	"github.com/bubbly-hb/blogSystem-gin/model"
	"github.com/bubbly-hb/blogSystem-gin/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
