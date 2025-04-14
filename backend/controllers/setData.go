package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zeotap/models"
	"zeotap/services"
)

func SetData (c *gin.Context) {
	var batch models.Batch
	
	if err := c.BindJSON(&batch); err != nil {
		return
	}
	
	err := services.WriteBatch(batch)
	
	if err != nil {
		fmt.Println(err)
	}
}
