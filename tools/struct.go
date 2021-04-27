package tools

import (
	"github.com/zedisdog/cola/e"
	"reflect"
)

func CopyFields(src interface{}, dest interface{}) error {
	var (
		sv reflect.Value
		dv reflect.Value
		dt reflect.Type
	)

	dt = reflect.TypeOf(dest)
	if dt.Kind() != reflect.Ptr {
		return e.New("need dest ptr")
	} else {
		dt = dt.Elem()
	}

	sv = reflect.ValueOf(src)
	if sv.Kind() == reflect.Ptr {
		sv = sv.Elem()
	}

	dv = reflect.ValueOf(dest).Elem()
	for i := 0; i < dt.NumField(); i++ {
		dtf := dt.Field(i)
		svf := sv.FieldByName(dtf.Name)
		if !svf.IsValid() || svf.IsZero() {
			continue
		}
		dvf := dv.Field(i)
		dvf.Set(svf)
	}

	return nil
}
