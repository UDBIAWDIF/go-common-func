package funcs

import "reflect"

type SliceType interface{ comparable }

func SliceFilter[T SliceType](sliceToFilter []T, filterFunc func(item T) bool) []T {
	sliceAfterFilter := []T{}
	for _, eachItem := range sliceToFilter {
		if filterFunc(eachItem) {
			sliceAfterFilter = append(sliceAfterFilter, eachItem)
		}
	}

	return sliceAfterFilter
}

func SliceGetEnd[T any](list []T) T {
	return list[len(list)-1]
}

func SliceContains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, e) {
			return true
		}
	}
	return false
}

func SliceRemoveDuplicateElement[T SliceType](sliceToRemoveDuplicate []T) []T {
	result := make([]T, 0, len(sliceToRemoveDuplicate))
	set := map[T]struct{}{}
	for _, item := range sliceToRemoveDuplicate {
		if _, ok := set[item]; !ok {
			set[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func SliceCovertToInterface[T any](sliceToCovert []T) []interface{} {
	result := make([]interface{}, 0, len(sliceToCovert))
	for _, item := range sliceToCovert {
		result = append(result, item)
	}
	return result
}

// 切片 转换成 map
// 一般用于快速找到特定数据
// 比如查询出来的数据列表, 用ID做 map 的 key
// 使用者可以很方便的找到 ID 对应的数据
// 不需要去遍历整个列表
func SliceCovertToMap[KEY_TYPE comparable, DATA_TYPE any](sliceToCovert []DATA_TYPE, keyName string) map[KEY_TYPE]DATA_TYPE {
	result := map[KEY_TYPE]DATA_TYPE{}
	for _, curItem := range sliceToCovert {
		curItemValue := reflect.ValueOf(curItem)
		keyField := curItemValue.FieldByName(keyName)
		if keyField.Kind() == reflect.Ptr { //类型为指针 需要取elem
			keyField = keyField.Elem()
		}
		result[keyField.Interface().(KEY_TYPE)] = curItem
	}
	return result
}
