package controllers

import (
	"net/http"
	"zeotap/models"
	"zeotap/services"

	"github.com/gin-gonic/gin"
)

// Connect handles POST requests to establish a connection to a database.
// It expects a JSON payload containing connection parameters such as host,
// port, database name, username, and password.
//
// Example JSON payload:
//
//	{
//	  "Host": "localhost",
//	  "Port": 9440,
//	  "Database": "default",
//	  "Username": "default",
//	  "Password": "password"
//	}
//
// If the connection is successful, it returns a JSON response with status "ok".
// If there is an error in binding the JSON or establishing the connection,
// it returns the corresponding error.
func Connect(c *gin.Context) {

	// Bind the incoming JSON payload to the ConnectionInfo struct
	var connectionInfo models.ConnectionInfo

	err := c.BindJSON(&connectionInfo)

	if err != nil {
		c.Error(err)
		return
	}

	// Attempt to establish a connection using the provided info
	_, err = services.Connect(connectionInfo)

	if err != nil {
		c.Error(err)
		return
	}

	// Respond with status OK if connection succeeds
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
