package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zeotap/models"
	"zeotap/services"
)

func SetData (c *gin.Context) {
	var connectionInfo models.ConnectionInfo
	
	if err := c.BindJSON(&connectionInfo); err != nil {
		return
	}
	
	err := services.WriteBatch(connectionInfo)
	
	if err != nil {
		fmt.Println(err)
	}
}
