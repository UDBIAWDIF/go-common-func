package utils

import (
	"testing"
)

// go test -v -run="TestRunner"
func TestRunner(t *testing.T) {
	runner := NewRunner()
	errMsg := ""
	runner.Exec(func() bool {
		errMsg = "Step 1 error"
		return true
	}).Exec(func() bool {
		errMsg = "Step 2 error"
		return true
	}).Exec(func() bool {
		errMsg = "Step 3 error"
		return false
	}).Exec(func() bool {
		errMsg = "Step 4 error"
		return true
	}).Exec(func() bool {
		errMsg = "Step 5 error"
		return true
	}).Success(func() {
		t.Log("All success")
	}).Failed(func() {
		t.Log("Failed! reason: {}", errMsg)
	})

	runner = NewRunner()
	errMsg = ""
	runner.Exec(func() bool {
		errMsg = "Step 1 error"
		return true
	}).Exec(func() bool {
		errMsg = "Step 2 error"
		return true
	}).Exec(func() bool {
		errMsg = "Step 3 error"
		return true
	}).Exec(func() bool {
		errMsg = "Step 4 error"
		return true
	}).Exec(func() bool {
		errMsg = "Step 5 error"
		return true
	}).Success(func() {
		t.Log("All success")
	}).Failed(func() {
		t.Log("Failed! reason: {}", errMsg)
	})
}
