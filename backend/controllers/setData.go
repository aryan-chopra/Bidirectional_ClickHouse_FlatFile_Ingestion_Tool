package controllers

import (
	"fmt"
	"net/http"

	"zeotap/models"
	"zeotap/services"

	"github.com/gin-gonic/gin"
)

func SetData (c *gin.Context) {
	var batch models.Batch
	
	if err := c.BindJSON(&batch); err != nil {
		fmt.Println(err)
		return
	}
	
	count, err := services.WriteBatch(batch)
	
	if err != nil {
		fmt.Println(err)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
