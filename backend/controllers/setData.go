package controllers

import (
	"fmt"
	"net/http"

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
	count, err := services.WriteBatch(batch)
	
	if err != nil {
		fmt.Println(err)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
