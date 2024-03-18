package utils

import (
	"testing"
)

var spinLock = &SpinLock{}

// go test -v -run="TestLock"
func TestLock(t *testing.T) {
	for {
		if !spinLock.TryLock() {
			continue
		}

		t.Log("Got lock!")

		spinLock.Unlock()
		break
	}
}
