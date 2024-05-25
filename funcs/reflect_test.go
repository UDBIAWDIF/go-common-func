package funcs

import (
	"testing"
	"time"
)

type UserForTest struct {
	Name      string
	Birthday  *time.Time
	LoginName *string
	Nickname  string
	Role      string
	Age       int32
	FakeAge   *int32
}

type EmployeeForTest struct {
	Name      string
	Birthday  *time.Time
	NickName  *string
	LoginName *string
	Age       int64
	FakeAge   int
	EmployeID int64
	DoubleAge int32
	SuperRule string
}

// go test -v -run="TestReflectCopyProperties"
func TestReflectCopyProperties(t *testing.T) {
	nickName := "付晓南老公"
	loginName := "admin"
	birthday, _ := StrToLocalTime("1979-01-01")
	em := EmployeeForTest{
		Name:      "uid",
		NickName:  &nickName,
		LoginName: &loginName,
		Age:       18,
		FakeAge:   20,
		Birthday:  &birthday,
	}
	user := UserForTest{}

	err := ReflectCopyProperties(&user, em)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(user)
}

// go test -v -run="TestReflectSetFieldValue"
func TestReflectSetFieldValue(t *testing.T) {
	nickName := "付晓南老公"
	birthday, _ := StrToLocalTime("1979-01-01")
	em := EmployeeForTest{
		// Name:     "uid",
		// NickName: &nickName,
		Age:      18,
		FakeAge:  20,
		Birthday: &birthday,
	}

	err := ReflectSetFieldValue(&em, "Name", "uid")
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(em)

	err = ReflectSetFieldValue(&em, "NickName", &nickName)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(em)
}
