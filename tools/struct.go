package tools

import (
	"github.com/zedisdog/cola/errx"
	"reflect"
)

//CopyFields copy fields from src to dest, note: dest mast be point
//	params:
//		src       source object
//		dest      point of dest object
//		copyZero  if copy zero field too
func CopyFields(src interface{}, dest interface{}, copyZero ...bool) error {
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
		if !sValueField.IsValid() && (copyZero[0] && sValueField.IsZero()) {
			continue
		}
		dValueField := dValue.Field(i)
		dValueField.Set(sValueField)
	}

	return nil
}
