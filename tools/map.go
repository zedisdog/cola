package tools

import (
	"encoding/json"
	"reflect"
)

func Struct2Map(data interface{}) map[string]interface{} {
	v := reflect.ValueOf(data)
	if v.Kind().String() != "struct" && v.Elem().Kind().String() != "struct" {
		return nil
	}
	j, _ := json.Marshal(data)
	r := make(map[string]interface{})
	_ = json.Unmarshal(j, &r)
	return r
}
