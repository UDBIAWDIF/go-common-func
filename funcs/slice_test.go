package funcs

import (
	"testing"
)

// type EmployeeForTest struct {
// 	Name      string
// 	Birthday  *time.Time
// 	NickName  *string
// 	LoginName *string
// 	Age       int64
// 	FakeAge   int
// 	EmployeID int64
// 	DoubleAge int32
// 	SuperRule string
// }

// go test -v -run="TestSliceCovertToMap"
func TestSliceCovertToMap(t *testing.T) {
	emList := []EmployeeForTest{}

	nickName := "付晓南老公"
	loginName := "admin"
	birthday, _ := StrToLocalTime("1979-01-01")
	emList = append(emList, EmployeeForTest{
		Name:      "uid",
		NickName:  &nickName,
		LoginName: &loginName,
		Age:       18,
		FakeAge:   20,
		Birthday:  &birthday,
	})

	nickName1 := "郑洪耕老婆"
	loginName1 := "beauty"
	birthday, _ = StrToLocalTime("1988-11-24")
	emList = append(emList, EmployeeForTest{
		Name:      "fxn",
		NickName:  &nickName1,
		LoginName: &loginName1,
		Age:       16,
		FakeAge:   15,
		Birthday:  &birthday,
	})

	nickName2 := "不知道谁"
	loginName2 := "unknown"
	birthday, _ = StrToLocalTime("2000-10-16")
	emList = append(emList, EmployeeForTest{
		Name:      "bb",
		NickName:  &nickName2,
		LoginName: &loginName2,
		Age:       45,
		FakeAge:   46,
		Birthday:  &birthday,
	})

	toMap := SliceCovertToMap[string](emList, "Name")
	PrintAsJson(toMap)
	t.Log(toMap)

	toMap = SliceCovertToMap[string](emList, "NickName")
	PrintAsJson(toMap)
	t.Log(toMap)
}
