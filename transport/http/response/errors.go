package response

import "net/http"

// Deprecated: use package errx instead
type HttpError struct {
	StatusCode int
	err        error
}

// Deprecated: use package errx instead
func (he HttpError) Error() string {
	return he.err.Error()
}

// Deprecated: use package errx instead
func (he HttpError) Unwrap() error {
	return he.err
}

// Deprecated: use package errx instead
func NewHttpErrorConflict(err error) error {
	return &HttpError{
		StatusCode: http.StatusConflict,
		err:        err,
	}
}

// Deprecated: use package errx instead
func NewHttpErrorUnprocessableEntity(err error) error {
	return &HttpError{
		StatusCode: http.StatusUnprocessableEntity,
		err:        err,
	}
}
