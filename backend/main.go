package main

import (
	"zeotap/controllers"
	"zeotap/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	router.Use(middlewares.ErrorHandler())

	router.POST("/connect", controllers.Connect)
	router.POST("/post", controllers.SetData)
	router.POST("/get-tables", controllers.GetTables)
	router.POST("get-rows", controllers.GetRows)

	router.Run(":8080")
}
