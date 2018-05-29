package genesis

import (
	"sync"
	"time"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/vm"
)

type vmTracer struct {
	data map[common.Address]map[common.Hash]common.Hash
	sync.RWMutex
}

func newVmTracer() *vmTracer {
	return &vmTracer{
		data: make(map[common.Address]map[common.Hash]common.Hash, 1024),
	}
}

func (vmt *vmTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	if err != nil {
		return err
	}
	if op == vm.SSTORE {
		s := stack.Data()

		addrStorage, ok := vmt.getAddrStorage(contract.Address())
		if !ok {
			addrStorage = make(map[common.Hash]common.Hash, 1024)
			vmt.setAddrStorage(contract.Address(), addrStorage)
		}
		addrStorage[common.BigToHash(s[len(s)-1])] = common.BigToHash(s[len(s)-2])
	}
	return nil
}

func (vmt *vmTracer) getAddrStorage(contractAddress common.Address) (map[common.Hash]common.Hash, bool) {
	vmt.RLock()
	defer vmt.RUnlock()

	addrStorage, ok := vmt.data[contractAddress]
	return addrStorage, ok
}

func (vmt *vmTracer) setAddrStorage(contractAddress common.Address, addrStorage map[common.Hash]common.Hash) {
	vmt.Lock()
	defer vmt.Unlock()

	vmt.data[contractAddress] = addrStorage
	return
}

func (vmt *vmTracer) setAddrStorageData(key, value common.Hash, addrStorage map[common.Hash]common.Hash) {
	vmt.Lock()
	defer vmt.Unlock()

	addrStorage[key] = value
	return
}

func (vmt *vmTracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	return nil
}
