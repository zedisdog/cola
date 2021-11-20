package errx

import (
	"errors"
	"testing"
)

func TestPanic(t *testing.T) {
	err := errors.New("123")
	err = Wrap(err, "321")
	panic(err)
}
