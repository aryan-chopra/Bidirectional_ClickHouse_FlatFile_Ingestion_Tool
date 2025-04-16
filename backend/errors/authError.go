package errors

import "net/http"

// AuthError represents an authentication-related error.
// It implements the HttpError interface by providing
// both an error message and an HTTP status code (401 Unauthorized)
type AuthError struct {
	decription string
	statusCode int
}

// MakeAuthError creates a new AuthError with the given description.
// The status code is automatically set to 401 Unauthorized.
func MakeAuthError(description string) AuthError {
	return AuthError{description, http.StatusUnauthorized}
}

// Error returns the error description.
func (err AuthError) Error() string {
	return err.decription
}

// StatusCode returns the HTTP status code for the error.
func (err AuthError) StatusCode() int {
	return err.statusCode
}
