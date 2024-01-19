package funcs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// 请注意它不能完全替代三元表达式, trueVal 和 falseVal 必须计算出来, 这意味着  If(len(list) > 0, list[0], 0) 这样使用时会因为 list 为空而产生异常
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckRunnable(err error, runnable *bool) {
	if err != nil {
		fmt.Println("CheckRunnable err:", err.Error())
		*runnable = false
	}
}

// 本软件是否64位
func AppIs64Bit() bool {
	return uint64(^uintptr(0)) == ^uint64(0)
}
