package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"zeotap/models"
	"zeotap/services"
)

func GetTables(c *gin.Context) {
	var tableInfo models.TableInfo
	
	if err := c.BindJSON(&tableInfo); err != nil {
		fmt.Println(err)
		return
	}
	
	tables, err := services.FetchTables(tableInfo)
	
	if (err != nil) {
		fmt.Println(err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}
