// Package errx is a package which enhance package errors
package errx

import (
	"errors"
	"fmt"
	"io"
	"runtime/debug"
)

var (
	NotFoundError        = errors.New("resource not found")
	UnknownFileTypeError = errors.New("unknown file type")
	InvalidFileError     = errors.New("invalid file")
	ConflictError        = errors.New("conflict")
	BadGetaway           = errors.New("bad getaway")
)

type HasStack interface {
	Stack() []byte
	Format(s fmt.State, r rune)
}

type Error struct {
	message string
	stack   []byte
	err     error
}

func New(message string) error {
	return &Error{
		err:     nil,
		message: message,
		stack:   debug.Stack(),
	}
}

func (e Error) Stack() []byte {
	return e.stack
}

func (e Error) Format(s fmt.State, r rune) {
	switch r {
	case 'w':
		_, _ = io.WriteString(s, e.message)
	case 'v':
		_, _ = io.WriteString(s, fmt.Sprintf("%s\n", e.message))
		_, _ = io.WriteString(s, fmt.Sprintf("\t%s\n", e.Stack))
	}
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Unwrap() error {
	return e.err
}

func WithStack(err error) error {
	return Wrap(err, err.Error())
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(HasStack); ok {
		return &Error{
			err:     err,
			message: message,
			stack:   e.Stack(),
		}
	} else {
		return &Error{
			err:     err,
			message: message,
			stack:   debug.Stack(),
		}
	}
}
