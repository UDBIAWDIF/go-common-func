package funcs

import (
	"testing"
)

type QueryForTest struct {
	ApplicableTo *string `json:"applicableTo" form:"applicableTo" search:"-"`
	Id           *int64  `json:"id" form:"id" search:"type:exact;column:id"`
	DeviceName   *string `json:"deviceName" form:"deviceName" search:"type:contains;column:device_name"`
	EquipmentNum *string `json:"equipmentNum" form:"equipmentNum" search:"type:exact;column:equipment_num"`
}

// go test -v -run="TestQueryFetchConditionMap"
func TestQueryFetchConditionMap(t *testing.T) {
	applicableTo := "car"
	deviceName := "fxnlg"
	query := QueryForTest{
		ApplicableTo: &applicableTo,
		DeviceName:   &deviceName,
	}
	toMap := map[string]interface{}{}
	QueryFetchConditionMap(query, toMap)
	PrintAsJson(toMap)
	t.Log(toMap)
}
