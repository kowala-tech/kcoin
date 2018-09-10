// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"io"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/math"
	"github.com/kowala-tech/kcoin/client/core/vm"
)

type JSONLogger struct {
	encoder *json.Encoder
	cfg     *vm.LogConfig
}

// NewJSONLogger creates a new VM tracer that prints execution steps as JSON objects
// into the provided stream.
func NewJSONLogger(cfg *vm.LogConfig, writer io.Writer) *JSONLogger {
	return &JSONLogger{json.NewEncoder(writer), cfg}
}

func (l *JSONLogger) CaptureStart(from common.Address, to common.Address, create bool, input []byte, resourceUsage uint64, value *big.Int) error {
	return nil
}

// CaptureState outputs state information on the logger.
func (l *JSONLogger) CaptureState(env *vm.VM, pc uint64, op vm.OpCode, resourceUsage, price uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	log := vm.StructLog{
		Pc:                  pc,
		Op:                  op,
		ComputationalEffort: resourceUsage,
		ComputeUnitPrice:    price,
		MemorySize:          memory.Len(),
		Storage:             nil,
		Depth:               depth,
		Err:                 err,
	}
	if !l.cfg.DisableMemory {
		log.Memory = memory.Data()
	}
	if !l.cfg.DisableStack {
		log.Stack = stack.Data()
	}
	return l.encoder.Encode(log)
}

// CaptureFault outputs state information on the logger.
func (l *JSONLogger) CaptureFault(env *vm.VM, pc uint64, op vm.OpCode, resourceUsage, price uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	return nil
}

// CaptureEnd is triggered at end of execution.
func (l *JSONLogger) CaptureEnd(output []byte, resourceUsage uint64, t time.Duration, err error) error {
	type endLog struct {
		Output        string              `json:"output"`
		ResourceUsage math.HexOrDecimal64 `json:"resourceUsage"`
		Time          time.Duration       `json:"time"`
		Err           string              `json:"error,omitempty"`
	}
	if err != nil {
		return l.encoder.Encode(endLog{common.Bytes2Hex(output), math.HexOrDecimal64(resourceUsage), t, err.Error()})
	}
	return l.encoder.Encode(endLog{common.Bytes2Hex(output), math.HexOrDecimal64(resourceUsage), t, ""})
}
