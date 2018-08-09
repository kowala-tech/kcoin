// +build linux,386

package scraper

/*
#cgo CFLAGS: -I./libsgx
#cgo CFLAGS: -I./libsgx/app/
#cgo LDFLAGS: -L./libsgx -loracle
#include "App.h"
*/
import (
	"C"
)

func Init() error {
	C.initSGX()
}

func Free() error {
	C.destroySGX()
}

func GetPrice() ([]byte, error) {
	/*
	var txSize uintptr
	txBuffer := make([]byte, 2048)
	txBufferUnsafe := C.CBytes(txBuffer)
	C.assemblePriceTx((*C.uint8_t)(txBufferUnsafe), (*C.size_t)(unsafe.Pointer(&y)))
	return C.GoBytes(unsafe.Pointer(txBufferC), *txSize)
	*/
	return []byte{}
}
