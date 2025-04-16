package controllers

import (
	"net/http"

	"zeotap/models"
	"zeotap/services"

	"github.com/gin-gonic/gin"
)

func SetData (c *gin.Context) {
	var batch models.Batch
	
	if err := c.BindJSON(&batch); err != nil {
		c.Error(err)
		return
	}
	
	count, err := services.WriteBatch(batch)
	
	if err != nil {
		c.Error(err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
