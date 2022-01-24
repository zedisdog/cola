package tools

import (
	"reflect"
)

func InSlice(target interface{}, slice interface{}) bool {
	ts := reflect.TypeOf(slice)
	if ts.Kind() != reflect.Slice {
		panic("slice required")
	}
	vs := reflect.ValueOf(slice)
	for i := 0; i < vs.Len(); i++ {
		item := vs.Index(i)
		if target == item.Interface() {
			return true
		}
	}
	return false
}
