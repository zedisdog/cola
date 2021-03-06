package errx

import (
	"fmt"
	"io"
	"net/http"
)

type HttpError struct {
	StatusCode int
	err        error
	Data       interface{}
}

func NewHttpError(code int, msg string) *HttpError {
	return &HttpError{
		StatusCode: code,
		err:        New(msg),
	}
}

func (he HttpError) Error() string {
	return he.err.Error()
}

func (he HttpError) Stack() []byte {
	if e, ok := he.err.(*Error); ok {
		return e.Stack()
	}
	return nil
}

func (he HttpError) Format(s fmt.State, r rune) {
	if e, ok := he.err.(*Error); ok {
		e.Format(s, r)
	}
	_, _ = io.WriteString(s, he.err.Error())
}

func (he HttpError) Unwrap() error {
	return he.err
}

func NewHttpErrorUnprocessableEntity(msg string) error {
	return NewHttpError(
		http.StatusUnprocessableEntity,
		msg,
	)
}

func NewHttpErrorConflict(msg string) error {
	return NewHttpError(
		http.StatusConflict,
		msg,
	)
}

//Deprecated: use NewHttpErrorForbidden instead
func NewHttpForbidden(msg string) error {
	return NewHttpError(
		http.StatusForbidden,
		msg,
	)
}

func NewHttpErrorForbidden(msg string) error {
	return NewHttpError(
		http.StatusForbidden,
		msg,
	)
}

func NewHttpErrorTeapot(msg string, data interface{}) error {
	err := NewHttpError(http.StatusTeapot, msg)
	if data != nil {
		err.Data = data
	}
	return err
}

func NewHttpErrorUnauthorized(msg string, data interface{}) error {
	err := NewHttpError(http.StatusUnauthorized, msg)
	if data != nil {
		err.Data = data
	}
	return err
}

func WarpByHttpError(code int, err error) *HttpError {
	return &HttpError{
		StatusCode: code,
		err:        err,
	}
}
