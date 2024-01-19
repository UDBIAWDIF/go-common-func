package funcs

import "reflect"

// TODO: 验证会不会比 IsEmpty() 好用
func ValidateIsEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

func isNotEmpty(a interface{}) bool {
	return !ValidateIsEmpty(a)
}
