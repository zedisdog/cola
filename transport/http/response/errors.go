package response

import "net/http"

type HttpError struct {
	StatusCode int
	err        error
}

func (he HttpError) Error() string {
	return he.err.Error()
}

func (he HttpError) Unwrap() error {
	return he.err
}

func NewHttpErrorConflict(err error) error {
	return &HttpError{
		StatusCode: http.StatusConflict,
		err:        err,
	}
}

func NewHttpErrorUnprocessableEntity(err error) error {
	return &HttpError{
		StatusCode: http.StatusUnprocessableEntity,
		err:        err,
	}
}
