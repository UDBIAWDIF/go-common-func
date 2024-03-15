package utils

import (
	"testing"
)

// go test -v -run="TestRunner"
func TestRunner(t *testing.T) {
	runner := NewRunner()
	errMsg := ""
	runner.exec(func() bool {
		errMsg = "Step 1 error"
		return true
	}).exec(func() bool {
		errMsg = "Step 2 error"
		return true
	}).exec(func() bool {
		errMsg = "Step 3 error"
		return false
	}).exec(func() bool {
		errMsg = "Step 4 error"
		return true
	}).exec(func() bool {
		errMsg = "Step 5 error"
		return true
	}).success(func() {
		t.Log("All success")
	}).failed(func() {
		t.Log("Failed! reason: {}", errMsg)
	})

	runner = NewRunner()
	errMsg = ""
	runner.exec(func() bool {
		errMsg = "Step 1 error"
		return true
	}).exec(func() bool {
		errMsg = "Step 2 error"
		return true
	}).exec(func() bool {
		errMsg = "Step 3 error"
		return true
	}).exec(func() bool {
		errMsg = "Step 4 error"
		return true
	}).exec(func() bool {
		errMsg = "Step 5 error"
		return true
	}).success(func() {
		t.Log("All success")
	}).failed(func() {
		t.Log("Failed! reason: {}", errMsg)
	})
}
