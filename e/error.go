// Package e
package e

import (
	"errors"
	"fmt"
	"io"
	"runtime/debug"
)

var (
	// Deprecated: use package errx instead
	NotFoundError = errors.New("resource not found")
	// Deprecated: use package errx instead
	UnknownFileTypeError = errors.New("unknown file type")
	// Deprecated: use package errx instead
	InvalidFileError = errors.New("无效的文件")
	// Deprecated: use package errx instead
	ConflictError = errors.New("资源冲突")
	// Deprecated: use package errx instead
	BadGetaway = errors.New("网关错误")
)

// Deprecated: use package errx instead
type Error struct {
	message string
	stack   []byte
	err     error
}

// Deprecated: use package errx instead
func (e Error) Stack() []byte {
	return e.stack
}

// Deprecated: use package errx instead
func (e Error) Format(s fmt.State, r rune) {
	switch r {
	case 'w':
		_, _ = io.WriteString(s, e.message)
	case 'v':
		_, _ = io.WriteString(s, fmt.Sprintf("%s\n", e.message))
		_, _ = io.WriteString(s, fmt.Sprintf("\t%s\n", e.Stack))
	}
}

// Deprecated: use package errx instead
func (e Error) Error() string {
	return e.message
}

// Deprecated: use package errx instead
func (e Error) Unwrap() error {
	return e.err
}

// Deprecated: use package errx instead
func New(message string) error {
	return &Error{
		err:     nil,
		message: message,
		stack:   debug.Stack(),
	}
}

// Deprecated: use package errx instead
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	return &Error{
		err:     err,
		message: err.Error(),
		stack:   debug.Stack(),
	}
}

// Deprecated: use package errx instead
func WithMessage(err error, message string) error {
	if e, ok := err.(*Error); ok {
		return &Error{
			err:     err,
			message: message,
			stack:   e.stack,
		}
	} else {
		e := Wrap(err, message)
		return e
	}

}

// Deprecated: use package errx instead
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &Error{
		err:     err,
		message: message,
		stack:   debug.Stack(),
	}
}
