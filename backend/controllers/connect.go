package controllers

import (
	"fmt"
	"net/http"
	"zeotap/models"
	"zeotap/services"

	"github.com/gin-gonic/gin"
)

func Connect(c *gin.Context) {
	var connectionInfo models.ConnectionInfo
	
	err := c.BindJSON(&connectionInfo)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	_, err = services.Connect(connectionInfo)
	
	if err != nil {
		fmt.Println(err)
		c.Error(err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"staus": "ok",
	})
}
