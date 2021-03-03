package tools

import (
	"testing"
)

type test struct {
	id int
}

func TestInArray(t *testing.T) {
	target := 1
	slice := []interface{}{1, "222"}
	r := InArray(target, slice)
	if !r {
		t.Fatal("fatal")
	}

	target1 := test{
		id: 2,
	}
	slice1 := []interface{}{target1}
	r = InArray(target1, slice1)
	if !r {
		t.Fatal("fatal")
	}
}
