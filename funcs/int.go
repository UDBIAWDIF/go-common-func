package funcs

import "strconv"

type number interface {
	int8 | int | int32 | int64 | float64 | float32
}

func IntToString(e int) string {
	return strconv.Itoa(e)
}

func IntAbs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}

// 数字比较, 相等返回 0; arg1大于arg2返回 1; arg1小于arg2返回 -1;
func NumberCompare[T number](arg1, arg2 T) int {
	result := 0
	if arg1 > arg2 {
		result = 1
	}
	if arg1 < arg2 {
		result = -1
	}
	return result
}

// 数字比较, arg1是否大于arg2
func NumberEQ[T number](arg1, arg2 T) bool {
	return NumberCompare(arg1, arg2) == 0
}

// 数字比较, arg1是否大于arg2
func NumberGT[T number](arg1, arg2 T) bool {
	return NumberCompare(arg1, arg2) == 1
}

// 数字比较, arg1是否小于arg2
func NumberLT[T number](arg1, arg2 T) bool {
	return NumberCompare(arg1, arg2) == -1
}

// 数字比较, arg1是否不大于arg2
func NumberNGT[T number](arg1, arg2 T) bool {
	return !NumberGT(arg1, arg2)
}

// 数字比较, arg1是否不小于arg2
func NumberNLT[T number](arg1, arg2 T) bool {
	return !NumberLT(arg1, arg2)
}
