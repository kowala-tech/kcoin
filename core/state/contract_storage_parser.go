package state

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/crypto/sha3"
)

var (
	// ErrNeedValidStructPointer is returned by the reflection code
	ErrNeedValidStructPointer = fmt.Errorf("r needs to be a valid pointer to a struct type")
	// ErrNilPointer is returned when finding a nil pointer
	ErrNilPointer = fmt.Errorf("can't marshal a nil pointer")
)

// UnmarshalState is  used to parse a contract's local storage, starting
// at the 256 bit word indicated by first into the struct passed
// as r (must be a pointer).
//
// The number of bytes to be parsed from the contract's storage is
// determined by the type of the fields defined in r.
//
// If the data type implements StateUnmarshaler, the UnmarshalStateBytes method
// is executed to parse the slice of bytes (a solSize tag is required).
//
// A solSize tag can be used to set the number of bytes to parse.
//
// When parsing uints/ints to *big.Int, it's possible to define the
// signal by using the solSign tag to set the value as signed or unsigned.
//
// A solStartWord tag can also be provided (with any value) to indicate
// that the parser should start at the begining of a 256 bit word.
//
// UnmarshalState maps Solidity to Go type in the following way:
//
// Ints and uints up to 64 bits can be parsed to the corresponding Go type
// or the closest match in size (a Solidity uint24 can be parsed into
// a uint32 or uint64). A *big.Int is used to parse ints and uints bigger
// than 64 bits. If there is no solSize or solSign present, defaults
// to uint256 (32 bytes in size, unsigned).
//
// Solidity bools and strings have a one to one match with Go.
//
// The Solidity type address can be parsed as a common.Address.
//
// Fixed size arrays map to a Go array and dynamic arrays map to a slice.
//
// Solidity mappings need to be parsed by the provided type Mapping.
//
// A field can be ignored by using the tag solIgnore (any value).
func (s *StateDB) UnmarshalState(addr common.Address, r interface{}) error {
	return s.GetOrNewStateObject(addr).UnmarshalState(r)
}

func (so *stateObject) UnmarshalState(r interface{}) error {
	// needs to be a pointer
	val := reflect.ValueOf(r)
	if val.Kind() != reflect.Ptr {
		return ErrNeedValidStructPointer
	}
	// can't be nil
	if val.IsNil() {
		return ErrNeedValidStructPointer
	}
	val = val.Elem()
	// must point to a struct
	if val.Kind() != reflect.Struct {
		return ErrNeedValidStructPointer
	}
	// parse
	return newStateReader(so, common.Hash{}).readValue(&val, 0, false)
}

var (
	bigMinusOne = big.NewInt(-1)
	big64k      = big.NewInt(65535)
	sizeOfInt   = int(unsafe.Sizeof(int(42)))
)

// import a slice of bytes in twos complement
func fromSignedBytes(b []byte) *big.Int {
	v := new(big.Int).SetBytes(b)
	if b[0]&0x80 == 0 || v.Cmp(common.Big0) == 0 {
		return v
	}
	bb := make([]byte, len(b)+1)
	bb[0] = 1
	bv := new(big.Int).SetBytes(bb)
	return bv.Mul(bv.Sub(bv, v), bigMinusOne)
}

// save a *big.Int as twos complement
func intoSignedBytes(v *big.Int, size int) []byte {
	bv := v.Bytes()
	if len(bv) > size {
		panic(fmt.Sprintf("bad size: %v, %v %q", len(bv), size, bv))
	}
	if v.Sign() >= 0 {
		return append(make([]byte, size-len(bv), size), bv...)
	}
	bv2 := make([]byte, size+1)
	bv2[0] = 1
	v2 := new(big.Int).SetBytes(bv2)
	bv2 = v2.Add(v2, v).Bytes()
	r := make([]byte, 0, size)
	for i := 0; i < size-len(bv2); i++ {
		r = append(r, 0xff)
	}
	return append(r, bv2...)
}

// StateUnmarshaler is implemented by data types that need to parse the contract raw bytes
type StateUnmarshaler interface {
	UnmarshalStateBytes([]byte) error
}

type stateIo struct {
	so     *stateObject
	curKey *big.Int
}

func newStateIo(so *stateObject, firstKey common.Hash) *stateIo {
	return &stateIo{
		so:     so,
		curKey: new(big.Int).SetBytes(firstKey.Bytes()),
	}
}

func (sio *stateIo) keyHash() common.Hash { return common.BytesToHash(sio.curKey.Bytes()) }
func (sio *stateIo) keyBytes() []byte     { return sio.keyHash().Bytes() }

// state parser
type stateReader struct {
	*stateIo
	curBytes []byte
}

// create a new state parser
func newStateReader(so *stateObject, firstKey common.Hash) *stateReader {
	return &stateReader{
		stateIo:  newStateIo(so, firstKey),
		curBytes: so.GetState(so.db.db, firstKey).Bytes(),
	}
}

func (sr *stateReader) nextFullWord() {
	if len(sr.curBytes) == 32 {
		return
	}
	sr.curKey.Add(sr.curKey, common.Big1)
	sr.curBytes = sr.so.GetState(sr.so.db.db, sr.keyHash()).Bytes()
}

func (sr *stateReader) readN(n int) []byte {
	if n > len(sr.curBytes) {
		sr.nextFullWord()
	}
	r := make([]byte, 0, n)
	for len(r) < n {
		if len(sr.curBytes) == 0 {
			sr.nextFullWord()
		}
		sz := n - len(r)
		if bsz := len(sr.curBytes); sz > bsz {
			sz = bsz
		}
		s := len(sr.curBytes) - sz
		r = append(r, sr.curBytes[s:]...)
		sr.curBytes = sr.curBytes[:s]
	}
	return r
}

func (sr *stateReader) readInt(size int, signed bool) *big.Int {
	b := sr.readN(size)
	if signed {
		return fromSignedBytes(b)
	}
	return new(big.Int).SetBytes(b)
}

func (sr *stateReader) readValue(v *reflect.Value, size int, signed bool) error {
	switch vv := v.Interface().(type) {
	case StateUnmarshaler:
		return vv.UnmarshalStateBytes(sr.readN(size))
	case *big.Int:
		sz := 32
		if size > 32 {
			return fmt.Errorf("can only read up to 32 bytes (%v requested)", size)
		}
		if size > 0 {
			sz = size
		}
		v.Set(reflect.ValueOf(sr.readInt(sz, signed)))
		return nil
	case common.Address:
		v.Set(reflect.ValueOf(common.BytesToAddress(sr.readN(20))))
		return nil
	case *Mapping:
		if vv == nil {
			v.Set(reflect.New(v.Type().Elem()))
			vv = v.Interface().(*Mapping)
		}
		vv.SetRoot(sr.so, sr.keyHash())
		if len(sr.curBytes) == 32 {
			sr.readN(1)
		}
		sr.nextFullWord()
		return nil
	}

	switch k := v.Kind(); k {
	case reflect.Int, reflect.Uint,
		reflect.Int64, reflect.Uint64,
		reflect.Int32, reflect.Uint32,
		reflect.Int16, reflect.Uint16,
		reflect.Int8, reflect.Uint8:
		var (
			isInt      bool
			sz, typeSz int
		)
		switch k {
		case reflect.Int:
			sz, isInt = sizeOfInt, true
		case reflect.Uint:
			sz, isInt = sizeOfInt, false
		case reflect.Int64:
			sz, isInt = 8, true
		case reflect.Uint64:
			sz, isInt = 8, false
		case reflect.Int32:
			sz, isInt = 4, true
		case reflect.Uint32:
			sz, isInt = 4, false
		case reflect.Int16:
			sz, isInt = 2, true
		case reflect.Uint16:
			sz, isInt = 2, false
		case reflect.Int8:
			sz, isInt = 1, true
		case reflect.Uint8:
			sz, isInt = 1, false
		}
		typeSz = sz
		if size > 0 {
			if size > sz {
				return fmt.Errorf("can only read %v bytes (%v requested)", sz, size)
			}
			sz = size
		}
		b := intoSignedBytes(sr.readInt(sz, isInt), typeSz)
		switch k {
		case reflect.Int:
			if sizeOfInt == 8 {
				v.Set(reflect.ValueOf(int(binary.BigEndian.Uint64(b))))
			} else {
				v.Set(reflect.ValueOf(int(binary.BigEndian.Uint32(b))))
			}
		case reflect.Uint:
			if sizeOfInt == 8 {
				v.Set(reflect.ValueOf(uint(binary.BigEndian.Uint64(b))))
			} else {
				v.Set(reflect.ValueOf(uint(binary.BigEndian.Uint32(b))))
			}
		case reflect.Int64:
			v.Set(reflect.ValueOf(int64(binary.BigEndian.Uint64(b))))
		case reflect.Uint64:
			v.Set(reflect.ValueOf(uint64(binary.BigEndian.Uint64(b))))
		case reflect.Int32:
			v.Set(reflect.ValueOf(int32(binary.BigEndian.Uint32(b))))
		case reflect.Uint32:
			v.Set(reflect.ValueOf(uint32(binary.BigEndian.Uint32(b))))
		case reflect.Int16:
			v.Set(reflect.ValueOf(int16(binary.BigEndian.Uint16(b))))
		case reflect.Uint16:
			v.Set(reflect.ValueOf(uint16(binary.BigEndian.Uint16(b))))
		case reflect.Int8:
			v.Set(reflect.ValueOf(int8(b[0])))
		case reflect.Uint8:
			v.Set(reflect.ValueOf(uint8(b[0])))
		}
	case reflect.Bool:
		var vb bool
		if sr.readN(1)[0] != 0 {
			vb = true
		}
		v.Set(reflect.ValueOf(vb))
		return nil
	case reflect.String:
		sr.nextFullWord()
		if sr.curBytes[31]%2 == 0 {
			fullWord := sr.readN(32)
			fullWord = fullWord[:fullWord[31]/2]
			v.Set(reflect.ValueOf(string(fullWord)))
			sr.nextFullWord()
		} else {
			sz := sr.readInt(31, false)
			sz.Div(sz, common.Big2)
			h := keccak256Sum(sr.keyBytes())
			nWords, rem := new(big.Int).DivMod(sz, common.Big32, new(big.Int))
			if rem.Cmp(common.Big0) != 0 {
				nWords.Add(nWords, common.Big1)
			}
			if nWords.Cmp(big64k) > 0 {
				panic("string to big")
			}
			strSr := newStateReader(sr.so, common.BytesToHash(h))
			nw := nWords.Uint64()
			s := strSr.readN(int(nw * 32))
			v.Set(reflect.ValueOf(string(s[:sz.Int64()])))
		}
		return nil
	case reflect.Struct:
		sr.nextFullWord()
		rt := v.Type()
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			sf := rt.Field(i)
			var (
				size   *int
				signed *bool
			)
			sign := false
			sz := 0
			if _, ok := sf.Tag.Lookup("solIgnore"); ok {
				continue
			}
			if strSize, ok := sf.Tag.Lookup("solSize"); ok {
				t, err := strconv.Atoi(strSize)
				if err != nil {
					return err
				}
				size = &t
			}
			if strSign, ok := sf.Tag.Lookup("solSign"); ok {
				var t bool
				switch strSign {
				case "signed":
					t = true
				case "unsigned":
					t = false
				default:
					return fmt.Errorf("unknown sign: %s", strSign)
				}
				signed = &t
			}
			if size == nil {
				size = &sz
			}
			if signed == nil {
				signed = &sign
			}
			if t, ok := sf.Tag.Lookup("solStartWord"); ok && t != "" {
				sr.nextFullWord()
			}
			if err := sr.readValue(&fv, *size, *signed); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array, reflect.Slice:
		var (
			newSr *stateReader
			sz    int
		)
		if k == reflect.Array {
			newSr = sr
			sz = v.Len()
		} else {
			sz = int(sr.readInt(31, false).Int64())
			newSr = newStateReader(sr.so, common.BytesToHash(keccak256Sum(sr.keyBytes())))
			sr.nextFullWord()
		}
		newS := reflect.New(reflect.SliceOf(v.Type().Elem())).Elem()
		for i := 0; i < sz; i++ {
			var newV reflect.Value
			if k == reflect.Array {
				newV = reflect.New(v.Type().Elem()).Elem()
			} else {
				newV = reflect.New(v.Type().Elem()).Elem()
			}
			if err := newSr.readValue(&newV, 0, false); err != nil {
				return err
			}
			newS = reflect.Append(newS, newV)
		}
		if k == reflect.Array {
			reflect.Copy(*v, newS)
		} else {
			v.Set(newS)
		}
	case reflect.Ptr:
		if v.IsNil() {
			nv := reflect.New(v.Type().Elem())
			v.Set(nv)
		}
		nv := v.Elem()
		return sr.readValue(&nv, size, signed)
	}
	return nil
}

func keccak256Sum(b []byte) []byte {
	h := sha3.NewKeccak256()
	h.Write(b)
	return h.Sum(nil)
}

// Mapping is the type used to access solidity's mapping type.
type Mapping struct {
	so     *stateObject
	mapKey common.Hash
}

// NewEmptyMapping returns a new (and without a root set) *Mapping.
func NewEmptyMapping() *Mapping { return &Mapping{} }

// NewMapping returns a new *Mapping.
func NewMapping(so *stateObject, key common.Hash) *Mapping {
	return &Mapping{so: so, mapKey: key}
}

// SetRoot is used to set the StateDB and address fields.
func (m *Mapping) SetRoot(so *stateObject, key common.Hash) {
	*m = Mapping{so: so, mapKey: key}
}

// Get unmarshals the value of v (must be a non nil pointer)
// using k as the key (only solidity primitive types allowed).
func (m *Mapping) Get(k, v interface{}) error {
	kb, err := marshalMappingKey(k)
	if err != nil {
		return err
	}
	switch k.(type) {
	case string:
	default:
		kb = append(make([]byte, 32-len(kb), 32), kb...)
	}
	return m.GetWithKeyBytes(kb, v)
}

// GetWithKeyBytes unmarshals the value of v (must be a non nil pointer)
// using k as the key (already marshaled data).
func (m *Mapping) GetWithKeyBytes(k []byte, v interface{}) error {
	vv := reflect.ValueOf(v)
	return newStateReader(m.so, m.getFirstKey(k)).readValue(&vv, 0, false)
}

func (m *Mapping) getFirstKey(b []byte) common.Hash {
	kb := make([]byte, 0, len(b)+32)
	kb = append(kb, b...)
	return common.BytesToHash(keccak256Sum(append(kb, m.mapKey.Bytes()...)))
}

func marshalMappingKey(k interface{}) (kb []byte, err error) {
	v := reflect.ValueOf(k)
	switch k.(type) {
	case int, uint, int64, uint64,
		int32, uint32, int16, uint16,
		int8, uint8, bool, *big.Int,
		common.Address, *common.Address:
		return valueAsBytes(v, 32)
	case string:
		return valueAsBytes(v, 0)
	}
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil, ErrNilPointer
		}
		return marshalMappingKey(v.Elem().Interface())
	}
	return nil, fmt.Errorf("\"%s\" does not match with a solidity primitive type", v.Type())
}

func whichIntSize(k reflect.Kind) int {
	switch k {
	case reflect.Int, reflect.Uint:
		return sizeOfInt
	case reflect.Int64, reflect.Uint64:
		return 8
	case reflect.Int32, reflect.Uint32:
		return 4
	case reflect.Int16, reflect.Uint16:
		return 2
	case reflect.Int8, reflect.Uint8:
		return 1
	}
	panic("only for ints")
}

func valueAsBytes(v reflect.Value, size int) ([]byte, error) {
	switch vv := v.Interface().(type) {
	case StateMarshaler:
		return vv.MarshalStateBytes()
	case *big.Int:
		b := vv.Bytes()
		sz := len(b)
		if vv.Cmp(common.Big0) < 0 && b[0]&0x80 != 0 {
			sz++
		}
		if sz < size {
			sz = size
		}
		return intoSignedBytes(vv, sz), nil
	case common.Address:
		return vv.Bytes(), nil
	case *common.Address:
		return vv.Bytes(), nil
	case string:
		return []byte(vv), nil
	case bool:
		if vv {
			return []byte{1}, nil
		}
		return []byte{0}, nil
	}
	switch k := v.Kind(); k {
	case reflect.Ptr:
		if v.IsNil() {
			return nil, nil
		}
		return valueAsBytes(v.Elem(), size)
	case reflect.Struct:
		vt := v.Type()
		bb := make([][]byte, 0, vt.NumField())
		bSz := 0
		for i := 0; i < vt.NumField(); i++ {
			sf := vt.Field(i)
			if _, ok := sf.Tag.Lookup("solIgnore"); ok {
				continue
			}
			var (
				fldSz int
				err   error
			)
			if szStr, ok := sf.Tag.Lookup("solSize"); ok {
				if fldSz, err = strconv.Atoi(szStr); err != nil {
					panic(fmt.Sprintf("bad size: %s", szStr))
				}
			}
			b, err := valueAsBytes(v.Field(i), fldSz)
			bb = append(bb, b)
			bSz += len(b)
		}
		return bytes.Join(packBytesIntoWords(bb), []byte{}), nil
	case reflect.Array, reflect.Slice:
		if v.Len() == 0 {
			return nil, nil
		}
		r := make([]byte, 0, 1024)
		lastWord := 0
		for i := 0; i < v.Len(); i++ {
			vv := v.Index(i)
			b, err := valueAsBytes(vv, size)
			if err != nil {
				return nil, err
			}
			if (len(r)+len(b))/32 != lastWord {
				r = append(r, make([]byte, (lastWord+1)*32-len(r))...)
			}
			r = append(r, b...)
		}
		return r, nil
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return intoSignedBytes(big.NewInt(v.Int()), whichIntSize(k)), nil
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return intoSignedBytes(big.NewInt(int64(v.Uint())), whichIntSize(k)), nil
	}
	return nil, nil
}

func packBytesIntoWords(chunks [][]byte) [][]byte {
	r := make([][]byte, 0, len(chunks))
	for len(chunks) > 0 {
		nChunks := 0
		remBytes := 0
		wsz := 0
		for nChunks = 0; nChunks < len(chunks); nChunks++ {
			if remBytes = wsz + len(chunks[nChunks]); remBytes > 32 {
				break
			}
		}
		b := make([]byte, remBytes, 32)
		for i := nChunks - 1; i < 0; i-- {
			b = append(b, chunks[i]...)
		}
		chunks = chunks[nChunks:]
		r = append(r, b)
	}
	return r
}

// StateMarshaler is implemented by types that need to customize the data layout
type StateMarshaler interface {
	MarshalStateBytes() ([]byte, error)
}
