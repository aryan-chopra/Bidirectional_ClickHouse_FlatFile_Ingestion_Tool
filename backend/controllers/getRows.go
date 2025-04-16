package controllers

import (
	"net/http"
	"zeotap/models"
	"zeotap/services"

	"github.com/gin-gonic/gin"
)

// GetRows handles POST requests to fetch rows from a data source.
// It expects a JSON body with connection and query details (models.RowInfo),
// passes that to the FetchRows service, and returns the results in the response.
//
// Expected JSON Body:
//
//	{
//	  "Connection": { ... },     // database connection info
//	  "Table": "table_name",     // name of the table to fetch from
//	  "start": 1000              // offset
//	}
//
// Response Format (on success):
//
//	{
//	  "columnNames": ["col1", "col2", ...],
//	  "rows": [ [...], [...], ... ],
//	  "items": <length of rows>
//	}
func GetRows(c *gin.Context) {
	var rowInfo models.RowInfo

	// Bind the incoming JSON to rowInfo struct
	err := c.BindJSON(&rowInfo)
	if err != nil {
		c.Error(err)
		return
	}

	// Call the service layer to fetch rows
	columnNames, fetchedRows, items, err := services.FetchRows(rowInfo)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items":       items,
		"columnNames": columnNames,
		"rows":        fetchedRows,
	})
}
