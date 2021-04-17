package tools

import (
	"reflect"
)

// Deprecated: use InSlice instead
func InArray(target interface{}, slice []interface{}) bool {
	for _, item := range slice {
		if reflect.DeepEqual(target, item) {
			return true
		}
	}
	return false
}

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
