package middlewares

import (
	"fmt"
	"net/http"
	"zeotap/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			fmt.Print("!ERROR!:")
			fmt.Println(err)
			switch e := err.Err.(type) {
				case errors.HttpError:
					c.AbortWithStatusJSON(e.StatusCode(), gin.H{
						"message": e.Error(),
					})
				default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Some Error Occured"})
			}
		}
	}
}
