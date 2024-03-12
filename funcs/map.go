package funcs

import (
	"sort"

	"github.com/iancoleman/orderedmap"
	"github.com/mitchellh/mapstructure"
)

// 普通的MAP转换成有序MAP(顺手按KEY升序排序)
func MapToOrderedMap(input map[string]interface{}) *orderedmap.OrderedMap {
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
