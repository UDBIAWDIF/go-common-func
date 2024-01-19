package funcs

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/axgle/mahonia"
)

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

// 字符串转float64
func StringToFloat64(e string) (float64, error) {
	return strconv.ParseFloat(e, 10)
}

// 字符串转float64, 不检测是否错误
func StringToFloat64NotError(e string) float64 {
	val, _ := StringToFloat64(e)
	return val
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func StringToBool(e string) (bool, error) {
	return strconv.ParseBool(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func FormatTimeStr(timeStr string) (string, error) {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", timeStr, loc)
	return theTime.Format("2006/01/02 15:04:05"), err
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func StringIsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// 从指定的内存地址取出字符串
// 过程: 一个个字节取出, 直到为 0 值结束
func UintptrToString(stringPtr uintptr) string {
	var buffBytes []byte
	stringPtrToMove := stringPtr
	for {
		eachByte := *((*byte)(unsafe.Pointer(stringPtrToMove)))
		if eachByte == 0 {
			break
		}
		buffBytes = append(buffBytes, eachByte)
		stringPtrToMove += 1
	}

	return string(buffBytes)
}

// 从指定的内存地址取出UTF8字符串
func UintptrToUTF8String(stringPtr uintptr) string {
	return ConvertGBKToUTF8(UintptrToString(stringPtr), "gbk", "utf-8")
}

func ConvertGBKToUTF8(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
