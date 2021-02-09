package e

import (
	"errors"
	"fmt"
	"io"
	"runtime/debug"
)

var (
	NotFoundError        = errors.New("resource not found")
	UnknownFileTypeError = errors.New("unknown file type")
	InvalidFileError     = errors.New("无效的文件")
	ConflictError        = errors.New("资源冲突")
	BadGetaway           = errors.New("网关错误")
)

type Error struct {
	message string
	Stack   []byte
	err     error
}

func (e *Error) Format(s fmt.State, r rune) {
	switch r {
	case 'w':
		io.WriteString(s, e.message)
	case 'v':
		io.WriteString(s, fmt.Sprintf("%s\n", e.message))
		io.WriteString(s, fmt.Sprintf("\t%s\n", e.Stack))
	}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Unwrap() error {
	return e.err
}

func New(message string) error {
	return &Error{
		err:     nil,
		message: message,
		Stack:   debug.Stack(),
	}
}

func WithStack(err error) error {
	if err == nil {
		return nil
	}
	return &Error{
		err:     err,
		message: err.Error(),
		Stack:   debug.Stack(),
	}
}

func WithMessage(err error, message string) error {
	if e, ok := err.(*Error); ok {
		return &Error{
			err:     err,
			message: message,
			Stack:   e.Stack,
		}
	} else {
		e := Wrap(err, message)
		return e
	}

}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &Error{
		err:     err,
		message: message,
		Stack:   debug.Stack(),
	}
}
