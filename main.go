package main

import (
	"blogger/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Static("/static", "static")
	engine.LoadHTMLGlob("views/*")
	engine.GET("/", controller.IndexHandler)
	engine.GET("/category", controller.CategoryHandler)

	err := engine.Run()
	if err != nil {
		panic(err)
	}
}
