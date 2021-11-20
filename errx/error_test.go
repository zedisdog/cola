package errx

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	err := errors.New("123")
	err = Wrap(err, "321")
	fmt.Printf("%v", err)
}
