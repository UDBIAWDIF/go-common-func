package funcs

import "reflect"

// 从 查询条件 生成查询 map
// 一般用于生成查询条件, 给GORM的 Where() 用
// query结构体每个属性都设置成指针类型
// 属性值不为空的就说明此属性的值要做为查询条件, 会被放到生成的 map 里
// WARNING: 会忽略不是指针类型的字段, 如果要做到更灵活, 以后再扩展
func QueryFetchConditionMap(query any, conditionMap map[string]interface{}) {
	dataType := reflect.TypeOf(query)
	dataValues := reflect.ValueOf(query)
	numField := dataType.NumField()
	for fieldIdx := 0; fieldIdx < numField; fieldIdx++ {
		elem := dataValues.Field(fieldIdx)

		if elem.Kind() == reflect.Struct {
			QueryFetchConditionMap(elem.Interface(), conditionMap)
			continue
		}

		if elem.IsZero() || elem.Kind() != reflect.Ptr {
			continue
		}

		conditionMap[dataType.Field(fieldIdx).Name] = elem.Elem().Interface()
	}
}
