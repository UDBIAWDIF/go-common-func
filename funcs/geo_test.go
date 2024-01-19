package funcs

import (
	"encoding/json"
	"testing"
)

// go test -v -run="TestCheckGISPointInArea"
func TestCheckGISPointInArea(t *testing.T) {
	regions := ""

	err := json.Unmarshal([]byte(""), &regions)
	t.Log(err.Error())

	json.Unmarshal([]byte("[[119.306423,26.097579],[119.306448,26.097167],[119.306714,26.097377]]"), &regions)
	t.Log(regions)
}
