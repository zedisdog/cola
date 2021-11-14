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

func NewHttpError(code int, msg string, data ...interface{}) *HttpError {
	err := &HttpError{
		StatusCode: code,
		err:        New(msg),
	}
	if data != nil && len(data) > 0 {
		err.Data = data[0]
	}
	return err
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

func NewHttpErrorUnprocessableEntityWithDetail(msg string, data map[string]string) error {
	return NewHttpError(
		http.StatusUnprocessableEntity,
		msg,
		data,
	)
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

func NewHttpErrorTeapot(msg string, data ...interface{}) error {
	err := NewHttpError(http.StatusTeapot, msg)
	if len(data) > 0 {
		err.Data = data[0]
	}
	return err
}

func NewHttpErrorUnauthorized(msg string, data ...interface{}) error {
	err := NewHttpError(http.StatusUnauthorized, msg)
	if len(data) > 0 {
		err.Data = data[0]
	}
	return err
}

func NewHttpErrorNotFound(msg string) error {
	return NewHttpError(http.StatusNotFound, msg)
}

func WrapByHttpError(code int, err error) *HttpError {
	return WarpByHttpError(code, err)
}

//Deprecated: use WrapByHttpError instead
func WarpByHttpError(code int, err error) *HttpError {
	return &HttpError{
		StatusCode: code,
		err:        err,
	}
}

func NewHttpErrorBadGateway(msg string) error {
	return NewHttpError(http.StatusBadGateway, msg)
}

func NewHttpErrorBadRequest(msg string) error {
	return NewHttpError(http.StatusBadRequest, msg)
}

func NewHttpErrorInternalServerError(msg string) error {
	return NewHttpError(http.StatusInternalServerError, msg)
}
