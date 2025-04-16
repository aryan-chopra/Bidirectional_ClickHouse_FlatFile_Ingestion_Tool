package errors

import "net/http"

// ConnectionError represents an error that occurs during a connection attempt.
// It implements the HttpError interface and returns a status code of 400 (Bad Request)
type ConnectionError struct {
	decription string
	statusCode int
}

// MakeConnectionError creates a new ConnectionError with the given description.
// The status code is automatically set to 400 Bad Request.
func MakeConnectionError(description string) ConnectionError {
	return ConnectionError{description, http.StatusBadRequest}
}

// Error returns the error description.
func (err ConnectionError) Error() string {
	return err.decription
}

// StatusCode returns the HTTP status code for the error.
func (err ConnectionError) StatusCode() int {
	return err.statusCode
}
