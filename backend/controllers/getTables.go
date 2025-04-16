package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zeotap/models"
	"zeotap/services"
)

// GetTables handles POST requests to fetch the list of tables from a connected database.
// It expects a JSON body containing connection details (models.TableInfo),
// uses the service layer to query available tables, and responds with a list of table names.
//
// Expected JSON Body:
//
//	{
//	  "connection": { ... } // DB connection information such as host, port, user, password, database
//	}
//
// Response Format (on success):
//
//	{
//	  "tables": ["table1", "table2", "table3", ...]
//	}
func GetTables(c *gin.Context) {
	var tableInfo models.TableInfo

	if err := c.BindJSON(&tableInfo); err != nil {
		c.Error(err)
		return
	}

	// Fetch tables using the service layer
	tables, err := services.FetchTables(tableInfo)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}
