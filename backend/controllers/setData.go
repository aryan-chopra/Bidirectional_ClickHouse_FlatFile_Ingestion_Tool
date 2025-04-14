package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ConnectionInfo struct {
	Host string
	Port int
	Database string
	User string
	Password string
}

func setData (c *gin.Context) error {
	var connectionInfo ConnectionInfo
	
	if err := c.BindJSON(&connectionInfo); err != nil {
		return err
	}
	
	fmt.Println(connectionInfo)
	return nil
}
