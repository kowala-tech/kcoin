package oracle

import (
	"C"
)

/*

#include "App.h"
#cgo LDFLAGS: -I./liboracle/App -L. -l./liboracle/oracle

*/

type SecurePriceProvider interface {
	Init()
	Free()
	GetPrice() []byte
}

// sgx represents a SGX implementation of a price provider
type sgx struct{}

// Init creates a new SGX enclave
func (s *sgx) Init() {
	//C.initSGX()
}

// Free destroys the active SGX enclave
func (s *sgx) Free() {
	//C.destroySGX()
}

// GetPrice returns a raw transaction containing the latest average price
func (s *sgx) GetPrice() []byte {
	var txSize uintptr
	txBuffer := make([]byte, 2048)
	txBufferUnsafe := C.CBytes(txBuffer)
	//C.assemblePriceTx((*C.uint8_t)(txBufferUnsafe), (*C.size_t)(unsafe.Pointer(&y)))
	//return C.GoBytes(unsafe.Pointer(txBufferC), *txSize)
	return []byte{}
}
