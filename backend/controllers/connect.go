package controllers

import (
	"net/http"
	"zeotap/models"
	"zeotap/services"

	"github.com/gin-gonic/gin"
)

func Connect(c *gin.Context) {
	var connectionInfo models.ConnectionInfo
	
	err := c.BindJSON(&connectionInfo)
	
	if err != nil {
		return
	}
	
	_, err = services.Connect(connectionInfo)
	
	if err != nil {
		c.Error(err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"staus": "ok",
	})
}
