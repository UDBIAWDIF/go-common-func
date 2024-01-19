package funcs

import "strconv"

func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

func Int64Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
