package genesis

import (
	"sync"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/vm"
)

type vmTracer struct {
	sync.RWMutex
	data map[common.Address]map[common.Hash]common.Hash
}

const defaultSize = 1024

func newVmTracer() *vmTracer {
	return &vmTracer{
		data: make(map[common.Address]map[common.Hash]common.Hash, defaultSize),
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
			addrStorage = make(map[common.Hash]common.Hash, defaultSize)
			vmt.setAddrStorage(contract.Address(), addrStorage)
		}
		addrStorage[common.BigToHash(s[len(s)-1])] = common.BigToHash(s[len(s)-2])
	}
	return nil
}

func (vmt *vmTracer) getAddrStorage(contractAddress common.Address) (addrStorage map[common.Hash]common.Hash, ok bool) {
	vmt.RLock()
	addrStorage, ok = vmt.data[contractAddress]
	vmt.RUnlock()
	return
}

func (vmt *vmTracer) setAddrStorage(contractAddress common.Address, addrStorage map[common.Hash]common.Hash) {
	vmt.Lock()
	vmt.data[contractAddress] = addrStorage
	vmt.Unlock()
}

func (vmt *vmTracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	return nil
}
