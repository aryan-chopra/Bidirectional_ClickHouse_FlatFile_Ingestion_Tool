package controllers

import (
	"net/http"

	"zeotap/models"
	"zeotap/services"

	"github.com/gin-gonic/gin"
)

// SetData handles POST requests to insert a batch of data rows into the database.
// It expects a JSON body containing batch data in the structure defined by models.Batch.
// The function delegates the actual write operation to the services layer.
//
// Expected JSON Body:
//
//	{
//	  "connection": { ... },  // Connection details to the DB
//	  "table": "your_table_name",
//	  "columns": ["col1", "col2", ...],
//	  "rows": [
//	    ["value1", "value2", ...],
//	    ["value1", "value2", ...],
//	    ...
//	  ]
//	}
//
// Response Format (on success):
//
//	{
//	  "count": <number_of_rows_successfully_written>
//	}
func SetData(c *gin.Context) {
	var batch models.Batch

	if err := c.BindJSON(&batch); err != nil {
		c.Error(err)
		return
	}

	// Write the batch to the database using the service layer
	count, err := services.WriteBatch(batch)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
