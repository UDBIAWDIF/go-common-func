package funcs

import (
	"errors"
	"reflect"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// 不建议使用的方法（即将过时）
// Deprecated method (out of date)
func StrToInt(err error, index string) int {
	result, err := strconv.Atoi(index)
	if err != nil {
		HasError(err, "string to int error"+err.Error(), -1)
	}
	return result
}

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Assert 条件断言
// 当断言条件为 假 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
func Assert(condition bool, msg string, code ...int) {
	if !condition {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func HasError(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		if msg == "" {
			msg = err.Error()
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// 判断某个元素是否在 slice, array, map 中
func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

// 从 slice, array 中找到指定元素的索引
// obj: 要查找的元素
// target: 要查找的slice, array
// 成功时返回 索引, nil
// 失败时返回 -1, 错误
func IndexOf(obj interface{}, target interface{}) (int, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for idx := 0; idx < targetValue.Len(); idx++ {
			if targetValue.Index(idx).Interface() == obj {
				return idx, nil
			}
		}
	}

	return -1, errors.New("not in array")
}
