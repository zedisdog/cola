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
	r := InSlice(target, slice)
	if !r {
		t.Fatal("fatal")
	}

	target1 := test{
		id: 2,
	}
	slice1 := []interface{}{target1}
	r = InSlice(target1, slice1)
	if !r {
		t.Fatal("fatal")
	}
}

func TestInSlice(t *testing.T) {
	target := 4
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	if !InSlice(target, slice) {
		t.Fatal("error")
	}

	td := "test"
	s := []interface{}{1, 2, 3, 4, 5, "test", 6, 7, 8, 9, 0}
	if !InSlice(td, s) {
		t.Fatal("error2")
	}
}
