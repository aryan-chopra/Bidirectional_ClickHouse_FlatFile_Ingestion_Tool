package controllers

import (
	"fmt"

	"zeotap/models"
	"zeotap/services"

	"github.com/gin-gonic/gin"
)

func SetData (c *gin.Context) {
	fmt.Println("Recieved req to post")
	
	var batch models.Batch
	
	if err := c.BindJSON(&batch); err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Calling the service")
	err := services.WriteBatch(batch)
	
	if err != nil {
		fmt.Println(err)
	}
}
