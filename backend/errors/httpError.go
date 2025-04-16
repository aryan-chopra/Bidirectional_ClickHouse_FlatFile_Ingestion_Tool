package errors

// HttpError defines a custom interface for HTTP-related errors.
// It allows for standardized error handling in web applications
// by combining an error message with an associated HTTP status code.
//
// Any struct implementing this interface must provide:
// - Error(): string        - a descriptive error message
// - StatusCode(): int      - the appropriate HTTP status code to return
//
// Example implementation:
//
//	type NotFoundError struct {
//	    Msg string
//	}
//
//	func (e NotFoundError) Error() string {
//	    return e.Msg
//	}
//
//	func (e NotFoundError) StatusCode() int {
//	    return http.StatusNotFound
//	}
type HttpError interface {
	Error() string
	StatusCode() int
}
