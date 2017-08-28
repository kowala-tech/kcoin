package state

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/crypto/sha3"
)

var (
	// // ErrNotEnoughData is returned when there is not enough data to parse
	// ErrNotEnoughData = errors.New("no data to parse")
	// ErrNeedValidStructPointer is returned by the reflection code
	ErrNeedValidStructPointer = errors.New("r needs to be a valid pointer to a struct type")

// // ErrSizeRequired is returned when the size field is set to a value < 1
// ErrSizeRequired = errors.New("size is required")
// // ErrToBig is returned when the size field is set to a value > 32
// ErrToBig = errors.New("size can't be bigger than 32")
)

// ParseState is used to parse a contract's local storage, starting
// at the 256 bit word indicated by "first" into the struct passed
// as r (must be a pointer).
//
// The amount of bytes to be parsed from the contract's storage is
// determined by the type of the fields defined in r.
//
// If the data type implements Parser, the ParseStateBytes
// method is executed to parse the slice of bytes.
//
// A "solSize" tag can be used to set the number of bytes to parse.
//
// When parsing uints/ints to *big.Int, it's possible to define the
// signal by using the "solSign" tag set to the value "signed" or "unsigned".
//
// A "solStartWord" tag can also be provided (with any value) to indicate
// that the parser should start at the begining of a 256 bit word.
//
// ParseState maps Solidity to Go type in the following way:
//
// Ints and uints up to 64 bits can be parsed to the corresponding Go type
// or the closest match in size (a Solidity uint24 can be parsed into
// a uin32 or uint64). A *big.Int is used to parse ints and uints bigger
// than 64 bits. If there is no "solSize" or "solSign" present, defaults
// to uint256 (32 bytes in size, unsigned).
//
// Solidity bools and strings have a one to one match with Go.
//
// The Solidity type address can be parsed as a common.Address.
//
// Fixed size arrays map to a Go array and dynamic arrays map to a slice.
//
// Mappings need to be parsed by the provided type Mapping.
func (s *StateDB) ParseState(addr *common.Address, first common.Hash, r interface{}) error {
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
	return newStateParser(s, *addr, first).readValue(&val, 0, false)
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

// Parser is implemented by data types that need to parse the contract's raw bytes
type Parser interface {
	ParseStateBytes([]byte) error
}

// state parser
type stateParser struct {
	sdb      *StateDB
	addr     common.Address
	curKey   *big.Int
	curBytes []byte
}

// create a new state parser
func newStateParser(sdb *StateDB, addr common.Address, firstKey common.Hash) *stateParser {
	return &stateParser{
		sdb:      sdb,
		addr:     addr,
		curKey:   new(big.Int).SetBytes(firstKey.Bytes()),
		curBytes: sdb.GetState(addr, firstKey).Bytes(),
	}
}

func (sp *stateParser) nextFullWord() {
	if len(sp.curBytes) == 32 {
		return
	}
	sp.curKey.Add(sp.curKey, common.Big1)
	sp.curBytes = sp.sdb.GetState(sp.addr, common.BytesToHash(sp.curKey.Bytes())).Bytes()
}

func (sp *stateParser) readN(n int) []byte {
	if n > len(sp.curBytes) {
		sp.nextFullWord()
	}
	r := make([]byte, 0, n)
	for len(r) < n {
		if len(sp.curBytes) == 0 {
			sp.nextFullWord()
		}
		sz := n - len(r)
		if bsz := len(sp.curBytes); sz > bsz {
			sz = bsz
		}
		s := len(sp.curBytes) - sz
		r = append(r, sp.curBytes[s:]...)
		sp.curBytes = sp.curBytes[:s]
	}
	return r
}

func (sp *stateParser) readInt(size int, signed bool) *big.Int {
	b := sp.readN(size)
	if signed {
		return fromSignedBytes(b)
	}
	return new(big.Int).SetBytes(b)
}

func (sp *stateParser) readValue(v *reflect.Value, size int, signed bool) error {
	switch vv := v.Interface().(type) {
	case Parser:
		return vv.ParseStateBytes(sp.readN(size))
	case *big.Int:
		sz := 32
		if size > 32 {
			return fmt.Errorf("can only read up to 32 bytes (%v requested)", size)
		}
		if size > 0 {
			sz = size
		}
		v.Set(reflect.ValueOf(sp.readInt(sz, signed)))
		return nil
	case common.Address:
		v.Set(reflect.ValueOf(common.BytesToAddress(sp.readN(20))))
		return nil
	case Mapping:
		// parse mappings here
		panic("mapping parsing not implemented")

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
			if size < sz {
				return fmt.Errorf("can only read %v bytes (%v requested)", sz, size)
			}
			sz = size
		}
		b := intoSignedBytes(sp.readInt(sz, isInt), typeSz)
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
		if sp.readN(1)[0] != 0 {
			vb = true
		}
		v.Set(reflect.ValueOf(vb))
		return nil
	case reflect.String:
		sp.nextFullWord()
		if sp.curBytes[31]%2 == 0 {
			fullWord := sp.readN(32)
			fullWord = fullWord[:fullWord[0]/2]
			v.Set(reflect.ValueOf(string(fullWord)))
		} else {
			sz := sp.readInt(32, false)
			h := keccak256Sum(sp.curKey.Bytes())

			fmt.Println(">>>", hex.EncodeToString(h))

			strSp := newStateParser(sp.sdb, sp.addr, common.BytesToHash(h))
			_ = strSp

			nWords := new(big.Int).Set(sz)
			rem := new(big.Int)
			nWords.DivMod(nWords, common.Big32, rem)
			if rem.Cmp(common.Big0) != 0 {
				nWords.Add(nWords, common.Big1)
			}
			if nWords.Cmp(big64k) > 0 {
				panic("string to big")
			}
			nw := nWords.Uint64()
			s := strSp.readN(int(nw * 32))
			v.Set(reflect.ValueOf(string(s[sz.Int64():])))
		}
		return nil
	case reflect.Struct:
		sp.nextFullWord()
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
				sp.nextFullWord()
			}
			if err := sp.readValue(&fv, *size, *signed); err != nil {
				return err
			}
		}
		return nil

	case reflect.Array:
		// parse arrays here
		panic("array parsing not implemented")
	case reflect.Slice:
		// parse slices here
		panic("slice parsing not implemented")
	case reflect.Ptr:
		if v.IsNil() {
			nv := reflect.New(v.Type().Elem())
			v.Set(nv)
		}
		nv := v.Elem()
		return sp.readValue(&nv, size, signed)
	}
	return nil
}

func keccak256Sum(b []byte) []byte {
	h := sha3.NewKeccak256()
	h.Write(b)
	return h.Sum(nil)
}

type Mapping struct {
	sdb    *StateDB
	addr   common.Address
	mapKey common.Hash
}

func NewMapping(sdb *StateDB, addr common.Address, key common.Hash) *Mapping {
	return &Mapping{sdb: sdb, addr: addr, mapKey: key}
}

// func (m *Mapping) FirstKey(key common.Hash) common.Hash {
// 	b := make([]byte, 0, len(key)*2)
// 	b = append(b, m.mapKey[:]...)
// 	b = append(b, key[:]...)
// 	return common.BytesToHash(keccak256Sum(b))
// }

// func (m *Mapping) Get(key common.Hash, r interface{}) error {
// 	v, err := valueOfStructPointer(r)
// 	if err != nil {
// 		return err
// 	}
// 	h := sha3.NewKeccak256()
// 	h.Write(m.mapKey[:])
// 	h.Write(key[:])
// 	k := h.Sum(nil)
// 	sp := newStateParser(m.sdb, m.addr, common.BytesToHash(k))
// 	return parseState(sp, v, nil)
// }
