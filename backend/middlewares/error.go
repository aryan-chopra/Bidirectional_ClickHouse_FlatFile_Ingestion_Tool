package middlewares

import (
	"fmt"
	"net/http"
	"zeotap/errors"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a Gin middleware function that handles errors that occur during the
// request processing cycle. It checks if the error is a custom HttpError and
// returns the appropriate HTTP status code and error message. If the error is of
// a different type or no custom error is provided, it returns a generic internal
// server error (500).
//
// It intercepts the request, processes the context, and checks for errors that may
// have been added during the request processing. If an error of type `HttpError` is
// found, the middleware responds with the custom status code and error message defined
// in that error type. For other error types, a generic 500 Internal Server Error response
// is returned with a generic message.
//
// Example usage:
//
//	r := gin.Default()
//	r.Use(middlewares.ErrorHandler()) // This will be used to catch and handle errors globally for all routes
//	r.GET("/", func(c *gin.Context) {
//		c.Error(errors.MakeConnectionError("Connection Failed"))
//	})
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Proceed with the request, allowing handlers to process it
		c.Next()

		// Iterate over all the errors that were accumulated during the request processing
		for _, err := range c.Errors {
			fmt.Print("!ERROR!:")
			fmt.Println(err)
			switch e := err.Err.(type) {
			// If the error is of type HttpError, respond with the associated status code and message
			case errors.HttpError:
				c.AbortWithStatusJSON(e.StatusCode(), gin.H{
					"message": e.Error(),
				})
			// If the error is not an HttpError, respond with a generic 500 Internal Server Error
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Some Error Occured"})
			}
		}
	}
}
