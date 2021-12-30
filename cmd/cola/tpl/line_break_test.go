package tpl

import (
	"fmt"
	"testing"
)

func TestGenLineBreak(t *testing.T) {
	s := GenLineBreak(0)
	fmt.Print(string(s))
}
