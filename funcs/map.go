package funcs

import (
	"sort"

	"github.com/iancoleman/orderedmap"
	"github.com/mitchellh/mapstructure"
)

// 普通的MAP转换成有序MAP(顺手按KEY升序排序)
func MapToOrderedMap[VAL any](input map[string]VAL) *orderedmap.OrderedMap {
	output := orderedmap.New()
	output.SetEscapeHTML(false)
	if input != nil {
		for eachKey, eachVal := range input {
			output.Set(eachKey, eachVal)
		}
		output.SortKeys(sort.Strings)
	}
	return output
}

// 结构体转换成MAP
func StructToStringMap(input interface{}) map[string]interface{} {
	var output map[string]interface{}
	mapstructure.Decode(input, &output)
	return output
}

// 结构体转换成有序MAP(顺手按KEY升序排序)
func StructToOrderedMap(input interface{}) *orderedmap.OrderedMap {
	return MapToOrderedMap(StructToStringMap(input))
}

// map 里的key从驼峰转下划线
// 主要用于查询条件转成实际数据库字段名
func MapKeyToSnakeCase(fromMap map[string]any) map[string]any {
	toMap := map[string]any{}
	for curKey, curVal := range fromMap {
		toMap[StrToSnakeCase(curKey)] = curVal
	}
	return toMap
}

func MapToSlice[M_KEY comparable, VAL any](fromMap map[M_KEY]VAL) []VAL {
	slice := make([]VAL, 0, len(fromMap))
	for _, val := range fromMap {
		slice = append(slice, val)
	}

	return slice
}

func OrderMapToSlice[VAL any](fromMap *orderedmap.OrderedMap) []VAL {
	slice := make([]VAL, 0, len(fromMap.Keys()))
	for _, mapKey := range fromMap.Keys() {
		val, exists := fromMap.Get(mapKey)
		if exists {
			slice = append(slice, val.(VAL))
		}
	}

	return slice
}
