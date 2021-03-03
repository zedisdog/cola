package tools

import (
	"reflect"
)

func InArray(target interface{}, slice []interface{}) bool {
	for _, item := range slice {
		if reflect.DeepEqual(target, item) {
			return true
		}
	}
	return false
}
