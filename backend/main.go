package main

import (
	"zeotap/controllers"
	"zeotap/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router instance with default middleware (logger and recovery)
	router := gin.Default()

	// Enabling Cross-Origin Resource Sharing (CORS) with default settings
	router.Use(cors.Default())

	// Adding a custom error handler middleware to handle errors globally
	router.Use(middlewares.ErrorHandler())

	// Routes for various DB related functions
	router.POST("/connect", controllers.Connect)
	router.POST("/post", controllers.SetData)
	router.POST("/get-tables", controllers.GetTables)
	router.POST("get-rows", controllers.GetRows)

	router.Run(":8080")
}
