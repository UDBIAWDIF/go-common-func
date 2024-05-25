package funcs

import (
	"errors"
	"fmt"
	"reflect"
)

// 获取指定属性的值
func ReflectGetFieldValue[T any](data interface{}, fieldName string) T {
	dataValue := reflect.ValueOf(data)
	return dataValue.FieldByName(fieldName).Interface().(T)
}

// 设置指定属性的值
func ReflectSetFieldValue[T any](dst interface{}, fieldName string, value T) error {
	runnable := true
	var err error = nil
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		err = errors.New("dst type should be a struct pointer")
		runnable = false
	}

	if runnable {
		dstValue = dstValue.Elem()
		fieldToSet := dstValue.FieldByName(fieldName)
		if fieldToSet.CanSet() {
			fieldToSet.Set(reflect.ValueOf(value))
		}
	}

	return err
}

// 属性复制
func ReflectCopyProperties(dst, src interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	fieldCount := dstType.NumField()
	for fieldIdx := 0; fieldIdx < fieldCount; fieldIdx++ {
		// 属性
		curField := dstType.Field(fieldIdx)
		// 待填充属性值
		curValue := srcValue.FieldByName(curField.Name)
		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !curValue.IsValid() || curField.Type != curValue.Type() {
			continue
		}

		if dstValue.Field(fieldIdx).CanSet() {
			dstValue.Field(fieldIdx).Set(curValue)
		}
	}

	return nil
}
