package ssl

import (
	"testing"
)

// go test -v -run="TestGenRsaKey"
func TestGenRsaKey(t *testing.T) {
	GenRsaKey(2048, "uidpw")
}
