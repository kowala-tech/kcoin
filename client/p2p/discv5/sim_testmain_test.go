// +build go1.4,nacl,faketime_simulation

package discv5

import (
	"os"
	"runtime"
	"testing"
	"unsafe"
)

// Enable fake time mode in the runtime, like on the go playground.
// There is a slight chance that this won't work because some go code
// might have executed before the variable is set.

//go:linkname faketime runtime.faketime
var faketime = 1

func TestMain(m *testing.M) {
	// We need to use unsafe somehow in order to get access to go:linkname.
	_ = unsafe.Sizeof(0)

	// Run the actual test. runWithPlaygroundTime ensures that the only test
	// that runs is the one calling it.
	runtime.GOMAXPROCS(8)
	os.Exit(m.Run())
}
