package errors

import "net/http"

type AuthError struct {
	decription string
	statusCode  int
}

func MakeAuthError(description string) AuthError {
	return AuthError{description, http.StatusUnauthorized}
}

func (err AuthError) Error() string {
	return err.decription
}

func (err AuthError) StatusCode() int {
	return err.statusCode
}