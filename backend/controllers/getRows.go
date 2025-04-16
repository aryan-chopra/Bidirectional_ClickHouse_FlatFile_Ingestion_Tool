package controllers

import (
	"fmt"
	"net/http"
	"zeotap/models"
	"zeotap/services"

	"github.com/gin-gonic/gin"
)

func GetRows (c *gin.Context) {
	var rowInfo models.RowInfo
	
	err := c.BindJSON(&rowInfo)
	
	if err != nil {
		fmt.Println(err)
		c.Error(err)
		return
	}
	
	columnNames, fetchedRows, items, err := services.FetchRows(rowInfo)
	
	if err != nil {
		fmt.Println(err)
		c.Error(err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"columnNames": columnNames,
		"rows": fetchedRows,
	})
}
