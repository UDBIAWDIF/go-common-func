package funcs

import (
	"testing"

	"github.com/mitchellh/mapstructure"
)

// go test -v -run="TestMapStruct"
func TestMapStruct(t *testing.T) {
	fromMap := map[string]int{
		"1664154087282737153": 100,
		"1664154087282737154": 101,
		"1664154087282737155": 102,
	}
	toMap := map[int64]int{}
	mapstructure.Decode(fromMap, &toMap)
	// 失败, key不能转类型

	PrintAsJson(toMap)
	t.Log(toMap)
}
