package oracle

import (
	"C"
)

/*

#include "./App/App.h"
#cgo LDFLAGS: -I./App -L. -loracle

*/

type SGX struct{}

func (s *SGX) init() {
	//C.initSGX()
}

func (s *SGX) free() {
	//C.destroySGX()
}

func (s *SGX) assemblePriceTx() []byte {
	var txSize uintptr
	txBuffer := make([]byte, 2048)
	txBufferUnsafe := C.CBytes(txBuffer)
	//C.assemblePriceTx((*C.uint8_t)(txBufferUnsafe), (*C.size_t)(unsafe.Pointer(&y)))
	//return C.GoBytes(unsafe.Pointer(txBufferC), *txSize)
	return []byte{}
}
