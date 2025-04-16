package errors

import "net/http"

type InternalServerError struct {
	description string
	statusCode  int
}

func MakeInternalServerError(description string) InternalServerError {
	return InternalServerError{description, http.StatusInternalServerError}
}

func (err InternalServerError) Error() string {
	return err.description
}

func (err InternalServerError) StatusCode() int {
	return err.statusCode
}
