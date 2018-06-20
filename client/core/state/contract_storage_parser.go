package state

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/crypto/sha3"
)

var (
	// ErrNeedValidStruct is returned by the reflection code
	ErrNeedValidStruct = fmt.Errorf("r needs to be a struct type (or pointer to)")
	// ErrNilPointer is returned when finding a nil pointer
	ErrNilPointer  = fmt.Errorf("can't marshal/unmarshal a nil pointer")
	ErrInvalidSize = fmt.Errorf("invalid size")
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
		return ErrNeedValidStruct
	}
	// can't be nil
	if val.IsNil() {
		return ErrNeedValidStruct
	}
	val = val.Elem()
	// must point to a struct
	if val.Kind() != reflect.Struct {
		return ErrNeedValidStruct
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
func (sio *stateIo) nextKey()             { sio.curKey.Add(sio.curKey, common.Big1) }

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

func (sr *stateReader) forceNextKey() {
	sr.stateIo.nextKey()
	sr.curBytes = sr.so.GetState(sr.so.db.db, sr.keyHash()).Bytes()
}

func (sr *stateReader) nextKey() {
	if len(sr.curBytes) == 32 {
		return
	}
	sr.forceNextKey()
}

func (sr *stateReader) readN(n int) []byte {
	if n > len(sr.curBytes) {
		sr.nextKey()
	}
	r := make([]byte, 0, n)
	for len(r) < n {
		if len(sr.curBytes) == 0 {
			sr.nextKey()
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
		if len(sr.curBytes) != 32 {
			sr.nextKey()
		}
		if vv == nil {
			v.Set(reflect.New(v.Type().Elem()))
			vv = v.Interface().(*Mapping)
		}
		vv.SetRoot(sr.so, sr.keyHash())
		if len(sr.curBytes) == 32 {
			sr.curBytes = []byte{}
		}
		sr.nextKey()
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
		sr.nextKey()
		if sr.curBytes[31]%2 == 0 {
			fullWord := sr.readN(32)
			fullWord = fullWord[:fullWord[31]/2]
			v.Set(reflect.ValueOf(string(fullWord)))
			sr.nextKey()
		} else {
			sz := sr.readInt(32, false)
			sz.Div(sz, common.Big2)
			h := keccak256Sum(sr.keyBytes())
			nWords, rem := new(big.Int).DivMod(sz, common.Big32, new(big.Int))
			if rem.Cmp(common.Big0) != 0 {
				nWords.Add(nWords, common.Big1)
			}
			if nWords.Cmp(big64k) > 0 {
				return fmt.Errorf("string to big %s", nWords)
			}
			strSr := newStateReader(sr.so, common.BytesToHash(h))
			nw := nWords.Uint64()
			s := strSr.readN(int(nw * 32))
			v.Set(reflect.ValueOf(string(s[:sz.Int64()])))
		}
		return nil
	case reflect.Struct:
		sr.nextKey()
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
				sr.nextKey()
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
			sr.nextKey()
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
	return m.GetWithKeyBytes(kb, v)
}

// GetWithKeyBytes unmarshals the value of v (must be a non nil pointer)
// using k as the key (already marshaled data).
func (m *Mapping) GetWithKeyBytes(k []byte, v interface{}) error {
	vv := reflect.ValueOf(v)
	return newStateReader(m.so, mappingFirstKey(k, m.mapKey)).readValue(&vv, 0, false)
}

// Set marshals the value of v, using k as the key.
func (m *Mapping) Set(k, v interface{}) error {
	kb, err := marshalMappingKey(k)
	if err != nil {
		return err
	}
	return m.SetWithKeyBytes(kb, v)
}

// Set marshals the value of v, using k as the already marshaled key.
func (m *Mapping) SetWithKeyBytes(k []byte, v interface{}) error {
	sw := newStateWriter(m.so, mappingFirstKey(k, m.mapKey))
	if err := sw.writeValue(reflect.ValueOf(v), 0, false); err != nil {
		return err
	}
	return sw.flush(true)
}

func mappingFirstKey(b []byte, mapKey common.Hash) common.Hash {
	kb := make([]byte, 0, len(b)+32)
	kb = append(kb, b...)
	return common.BytesToHash(keccak256Sum(append(kb, mapKey.Bytes()...)))
}

func marshalMappingKey(k interface{}) (kb []byte, err error) {
	v := reflect.ValueOf(k)
	switch k.(type) {
	case int, uint, int64, uint64,
		int32, uint32, int16, uint16,
		int8, uint8, bool, *big.Int,
		common.Address, *common.Address:
		b, err := valueAsBytes(v, 32, nil)
		if err != nil {
			return nil, err
		}
		return append(make([]byte, 32-len(b), 32), b...), nil
	case string:
		return valueAsBytes(v, 0, nil)
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

func valueAsBytes(v reflect.Value, size int, signed *bool) ([]byte, error) {
	switch vv := v.Interface().(type) {
	case StateMarshaler:
		return vv.MarshalStateBytes()
	case *big.Int:
		var sign bool
		if signed != nil && *signed == true {
			sign = true
		}
		b := vv.Bytes()
		sz := len(b)
		if vv.Cmp(common.Big0) < 0 && b[0]&0x80 != 0 {
			sz++
		}
		if size < 1 {
			sz = 32
		} else if sz < size {
			sz = size
		}
		if sign {
			return intoSignedBytes(vv, sz), nil
		}
		return append(make([]byte, sz-len(b), sz), b...), nil
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
		return valueAsBytes(v.Elem(), size, nil)
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		var sz int
		if wSz := whichIntSize(k); size == 0 && size < wSz {
			sz = wSz
		} else {
			sz = size
		}
		b := intoSignedBytes(big.NewInt(v.Int()), sz)
		return b, nil
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		var sz int
		if wSz := whichIntSize(k); size == 0 && size < wSz {
			sz = wSz
		} else {
			sz = size
		}
		b := intoSignedBytes(big.NewInt(int64(v.Uint())), sz)
		return b, nil
	}
	return nil, nil
}

// StateMarshaler is implemented by types that need to customize the data layout
type StateMarshaler interface {
	MarshalStateBytes() ([]byte, error)
}

func (s *StateDB) MarshalState(addr common.Address, v interface{}) error {
	return s.GetOrNewStateObject(addr).MarshalState(v)
}

func (so *stateObject) MarshalState(v interface{}) error {
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return ErrNilPointer
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return ErrNeedValidStruct
	}
	sw := newStateWriter(so, common.Hash{})
	if err := sw.writeValue(val, 0, false); err != nil {
		return err
	}
	return sw.flush(true)
}

type stateWriter struct {
	*stateIo
	buf   [][]byte
	bufSz int
}

func newStateWriter(so *stateObject, firstKey common.Hash) *stateWriter {
	return &stateWriter{
		stateIo: newStateIo(so, firstKey),
	}
}

func (sw *stateWriter) completeWord() {
	sz := 32 - sw.bufSz%32
	if sz == 0 || sz == 32 {
		return
	}
	sw.buf = append(sw.buf, make([]byte, sz))
	sw.bufSz += sz
}

func (sw *stateWriter) isAtWordBoundary() bool {
	return sw.bufSz != 0 && sw.bufSz%32 == 0
}

func (sw *stateWriter) flush(all bool) error {
	if all {
		sw.completeWord()
	}
	for sw.bufSz > 31 {
		chunks := make([][]byte, 0, sw.bufSz/32)
		for sz := 0; sz < 32; {
			sz += len(sw.buf[0])
			chunks = append(chunks, sw.buf[0])
			sw.bufSz -= len(sw.buf[0])
			sw.buf = sw.buf[1:]
		}
		b := make([]byte, 0, 32)
		for i := len(chunks) - 1; i > -1; i-- {
			b = append(b, chunks[i]...)
		}
		sw.so.SetState(sw.so.db.db, common.BigToHash(sw.curKey), common.BytesToHash(b))
		sw.nextKey()
	}
	return nil
}

func (sw *stateWriter) writeBytes(b []byte) error {
	bSz := len(b)
	if bSz > 32 {
		if bSz%32 != 0 {
			return ErrInvalidSize
		}
		sw.completeWord()
	} else {
		remBytes := 32 - sw.bufSz%32
		if bSz > remBytes {
			sw.completeWord()
		}
	}
	for len(b) > 0 {
		var ssz int
		if len(b) > 31 {
			ssz = 32
		} else {
			ssz = len(b)
		}
		sw.buf = append(sw.buf, b[:ssz])
		sw.bufSz += ssz
		b = b[ssz:]
	}
	return sw.flush(false)
}

func (sw *stateWriter) writeStruct(v reflect.Value) error {
	if err := sw.flush(true); err != nil {
		return err
	}
	vt := v.Type()
	for i := 0; i < vt.NumField(); i++ {
		sf := vt.Field(i)
		fld := v.Field(i)
		vi := fld.Interface()
		if a, ok := vi.(common.Address); ok {
			if err := sw.writeBytes(a.Bytes()); err != nil {
				return err
			}
			continue
		} else if a, ok := vi.(*common.Address); ok {
			if err := sw.writeBytes(a.Bytes()); err != nil {
				return err
			}
			continue
		}
		switch fk := sf.Type.Kind(); fk {
		case reflect.String:
			if err := sw.writeString(fld.String()); err != nil {
				return err
			}
		case reflect.Array, reflect.Slice:
			if err := sw.flush(true); err != nil {
				return err
			}
			var nsw *stateWriter
			if fk == reflect.Slice {
				if err := sw.writeBytes(big.NewInt(int64(fld.Len())).Bytes()); err != nil {
					return err
				}
				nsw = newStateWriter(sw.so, common.BytesToHash(keccak256Sum(sw.keyBytes())))
				if err := sw.flush(true); err != nil {
					return err
				}
			} else {
				nsw = sw
			}
			for i := 0; i < fld.Len(); i++ {
				if err := nsw.writeValue(fld.Index(i), 0, true); err != nil {
					return err
				}
			}
			if err := nsw.flush(true); err != nil {
				return err
			}
		case reflect.Map:
			mapping := NewMapping(sw.so, sw.keyHash())
			for _, mk := range fld.MapKeys() {
				if err := mapping.Set(
					mk.Interface(),
					fld.MapIndex(mk).Interface(),
				); err != nil {
					return err
				}
			}
			if err := sw.writeBytes(make([]byte, 32)); err != nil {
				return err
			}
			if err := sw.flush(true); err != nil {
				return err
			}
		default:
			if _, ok := sf.Tag.Lookup("solIgnore"); ok {
				continue
			}
			var (
				fldSz int
				err   error
			)
			if szStr, ok := sf.Tag.Lookup("solSize"); ok {
				if fldSz, err = strconv.Atoi(szStr); err != nil {
					return fmt.Errorf("bad size: %s", szStr)
				}
			}
			var signed *bool
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
			bb, err := valueAsBytes(v.Field(i), fldSz, signed)
			if err != nil {
				return err
			}
			if err := sw.writeBytes(bb); err != nil {
				return err
			}
		}
	}
	return nil
}

func (sw *stateWriter) writeString(s string) error {
	b := []byte(s)
	bSz := len(b)
	if bSz < 32 {
		b = append(b, make([]byte, 32-len(b))...)
		b[31] = byte(bSz) * 2
		return sw.writeBytes(b)
	}
	b = append(b, make([]byte, 32-len(b)%32)...)
	if err := sw.flush(true); err != nil {
		return err
	}
	kb := sw.keyBytes()
	if err := sw.writeBytes(big.NewInt(int64(bSz*2 + 1)).Bytes()); err != nil {
		return err
	}
	ssw := newStateWriter(sw.so, common.BytesToHash(keccak256Sum(common.BytesToHash(kb).Bytes())))
	if err := ssw.writeBytes(b); err != nil {
		return err
	}
	if err := ssw.flush(true); err != nil {
		return err
	}
	return nil
}

func (sw *stateWriter) writeValue(v reflect.Value, size int, signed bool) error {
	switch kv := v.Kind(); kv {
	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return sw.writeValue(v.Elem(), size, signed)
	case reflect.String:
		if err := sw.writeString(v.String()); err != nil {
			return err
		}
	// case *Mapping:
	case reflect.Struct:
		if err := sw.writeStruct(v); err != nil {
			return err
		}
	// 	case reflect.Array, reflect.Slice:
	default:
		b, err := valueAsBytes(v, size, nil)
		if err != nil {
			return err
		}
		return sw.writeBytes(b)
	}
	return nil
}
