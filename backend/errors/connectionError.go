package errors

import "net/http"

type ConnectionError struct {
	decription string
	statusCode  int
}

func MakeConnectionError(description string) ConnectionError {
	return ConnectionError{description, http.StatusBadRequest}
}

func (err ConnectionError) Error() string {
	return err.decription
}

func (err ConnectionError) StatusCode() int {
	return err.statusCode
}
