package errors

import "net/http"

// InternalServerError represents an error that occurs due to internal server issues.
// It implements the HttpError interface and returns a status code of 500 (Internal Server Error).
type InternalServerError struct {
	description string
	statusCode  int
}

// MakeInternalServerError creates a new InternalServerError with the given description.
// The status code is automatically set to 500 (Internal Server Error).
func MakeInternalServerError(description string) InternalServerError {
	return InternalServerError{description, http.StatusInternalServerError}
}

// Error returns the error description.
func (err InternalServerError) Error() string {
	return err.description
}

// StatusCode returns the HTTP status code for the error.
func (err InternalServerError) StatusCode() int {
	return err.statusCode
}
