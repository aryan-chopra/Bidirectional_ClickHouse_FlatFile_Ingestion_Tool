package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zeotap/models"
	"zeotap/services"
)

func GetTables(c *gin.Context) {
	var tableInfo models.TableInfo
	
	if err := c.BindJSON(&tableInfo); err != nil {
		c.Error(err)
		return
	}
	
	tables, err := services.FetchTables(tableInfo)
	
	if (err != nil) {
		c.Error(err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}
