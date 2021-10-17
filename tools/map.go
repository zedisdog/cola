package tools

import (
	"encoding/json"
	"reflect"
)

//Struct2Map covert struct to map use json package
//TODO: use reflect instead.
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
