package main

import (
	"zeotap/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.POST("/post", controllers.SetData)
	
	router.Run(":8080")
}
