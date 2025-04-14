package main

import "github.com/gin-gonic/gin"
import "zeotap/controllers"

func main() {
	router := gin.Default()
	
	router.POST("/post", controllers.SetData)
	
	router.Run(":8080")
}
