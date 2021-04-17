package tools

import (
	"reflect"
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

func TestNormal(t *testing.T) {
	var a interface{} = &[]string{"test"}
	//va := reflect.ValueOf(a)
	ta := reflect.TypeOf(a)
	println(ta.Kind().String())
	_, ok := a.([]interface{})
	if !ok {
		t.Fatal("error")
	}
}
