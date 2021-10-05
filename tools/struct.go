package tools

import (
	"github.com/zedisdog/cola/errx"
	"reflect"
)

func CopyFields(src interface{}, dest interface{}) error {
	var (
		sValue reflect.Value
		dValue reflect.Value
		dType  reflect.Type
	)

	// dest必须为指针
	dType = reflect.TypeOf(dest)
	if dType.Kind() != reflect.Ptr {
		return errx.New("need dest ptr")
	} else {
		dType = dType.Elem()
	}

	// 取src的value, 如果是指针就避开指针
	sValue = reflect.ValueOf(src)
	if sValue.Kind() == reflect.Ptr {
		sValue = sValue.Elem()
	}

	dValue = reflect.ValueOf(dest).Elem()
	for i := 0; i < dType.NumField(); i++ {
		dTypeField := dType.Field(i)
		sValueField := sValue.FieldByName(dTypeField.Name)
		if !sValueField.IsValid() {
			continue
		}
		dValueField := dValue.Field(i)
		dValueField.Set(sValueField)
	}

	return nil
}
