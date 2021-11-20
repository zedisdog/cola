// Package errx is a package which enhance package errors
package errx

import (
	"errors"
	"fmt"
	"github.com/zedisdog/cola/i18n"
	"io"
	"runtime"
	"runtime/debug"
)

var (
	// Deprecated: use customize instead
	NotFoundError = errors.New("resource not found")
	// Deprecated: use customize instead
	UnknownFileTypeError = errors.New("unknown file type")
	// Deprecated: use customize instead
	InvalidFileError = errors.New("invalid file")
	// Deprecated: use customize instead
	ConflictError = errors.New("conflict")
	// Deprecated: use customize instead
	BadGetaway = errors.New("bad getaway")
)

type HasStack interface {
	Stack() []byte
	Format(s fmt.State, r rune)
}

type Error struct {
	file    string
	line    int
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
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'w':
		_, _ = io.WriteString(s, e.Error())
	case 'v':
		_, _ = io.WriteString(s, fmt.Sprintf("%s\n", e.Error()))
		_, _ = io.WriteString(s, fmt.Sprintf("%s\n", string(e.Stack())))
	}
}

//Error return error string translate by i18n
func (e Error) Error() string {
	return fmt.Sprintf("%s:%d %s\n%s",
		e.file,
		e.line,
		i18n.Trans(e.message),
		e.Unwrap(),
	)
}

//RawError return raw error string
func (e Error) RawError() string {
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
	var e Error
	if e, ok := err.(HasStack); ok {
		e = &Error{
			stack: e.Stack(),
		}
	} else {
		e = &Error{
			stack: debug.Stack(),
		}
	}

	e.message = message
	e.err = err
	_, e.file, e.line, _ = runtime.Caller(1)
	return e
}

func WrapOrNew(err error, message string) error {
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
