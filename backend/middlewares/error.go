package middlewares

import (
	"net/http"
	"zeotap/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
				case errors.HttpError:
					c.AbortWithStatusJSON(e.StatusCode(), e)
				default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Some Error Occured"})
			}
		}
	}
}
