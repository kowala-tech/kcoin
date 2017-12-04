package state_test

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/accounts/abi/bind/backends"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/crypto"
	"github.com/kowala-tech/kUSD/params"
)

//go:generate solc --bin --abi --overwrite --out test_contracts test_contracts/SmallInts.sol
//go:generate solc --bin --abi --overwrite --out test_contracts test_contracts/BigInts.sol
//go:generate solc --bin --abi --overwrite --out test_contracts test_contracts/Strings.sol
//go:generate solc --bin --abi --overwrite --out test_contracts test_contracts/Arrays.sol
//go:generate solc --bin --abi --overwrite --out test_contracts test_contracts/Mappings.sol

func parseFromTest(name string, v interface{}) error {
	// read contract bytecode
	hexBytecode, err := ioutil.ReadFile(path.Join("test_contracts", name+".bin"))
	if err != nil {
		return err
	}
	contractBytecode := make([]byte, len(hexBytecode)/2)
	if _, err = hex.Decode(contractBytecode, hexBytecode); err != nil {
		return err
	}
	// read contract ABI
	f, err := os.Open(path.Join("test_contracts", name+".abi"))
	if err != nil {
		return err
	}
	defer f.Close()
	contractABI, err := abi.JSON(f)
	if err != nil {
		return err
	}
	// generate new private key
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	// create a new simulated backend
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		crypto.PubkeyToAddress(privKey.PublicKey): core.GenesisAccount{
			Balance: new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(params.Ether)),
		},
	})
	sim.Commit()
	// deploy contract
	contractAddr, _, _, err := bind.DeployContract(
		bind.NewKeyedTransactor(privKey),
		contractABI,
		contractBytecode,
		sim,
	)
	if err != nil {
		return err
	}
	sim.Commit()
	// get stateDB
	stateDb, err := sim.BlockChain.State()
	if err != nil {
		return err
	}
	// parse state
	err = stateDb.UnmarshalState(
		&contractAddr,
		common.Hash{},
		v,
	)
	if err != nil {
		return err
	}
	return nil
}

type dataSmallInts struct {
	I1 int8
	I2 int16
	I3 int32 `solSize:"3"`
	I4 int32
	I5 int64 `solSize:"5"`
	I6 int64 `solSize:"6"`
	I7 int64 `solSize:"7"`
	I8 int64

	U1 uint8
	U2 uint16
	U3 uint32 `solSize:"3"`
	U4 uint32
	U5 uint64 `solSize:"5"`
	U6 uint64 `solSize:"6"`
	U7 uint64 `solSize:"7"`
	U8 uint64

	BoolByte bool
}

func TestContractStorageParserSmallInts(t *testing.T) {
	exp := dataSmallInts{I1: -127, U1: 129, BoolByte: true}
	exp.I2 = int16(exp.I1-1)*256 + 1
	exp.I3 = int32(exp.I2-1)*256 + 1
	exp.I4 = int32(exp.I3-1)*256 + 1
	exp.I5 = int64(exp.I4-1)*256 + 1
	exp.I6 = int64(exp.I5-1)*256 + 1
	exp.I7 = int64(exp.I6-1)*256 + 1
	exp.I8 = int64(exp.I7-1)*256 + 1
	exp.U2 = uint16(exp.U1-1)*256 + 1
	exp.U3 = uint32(exp.U2-1)*256 + 1
	exp.U4 = uint32(exp.U3-1)*256 + 1
	exp.U5 = uint64(exp.U4-1)*256 + 1
	exp.U6 = uint64(exp.U5-1)*256 + 1
	exp.U7 = uint64(exp.U6-1)*256 + 1
	exp.U8 = uint64(exp.U7-1)*256 + 1

	got := dataSmallInts{}

	if err := parseFromTest("SmallInts", &got); err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(&exp, &got) {
		t.Errorf("expected: %v\ngot: %v\n", exp, got)
		return
	}
}

type dataBigInts struct {
	I9  *big.Int `solSize:"9" solSign:"signed"`
	I10 *big.Int `solSize:"10" solSign:"signed"`
	I11 *big.Int `solSize:"11" solSign:"signed"`
	I12 *big.Int `solSize:"12" solSign:"signed"`
	I13 *big.Int `solSize:"13" solSign:"signed"`
	I14 *big.Int `solSize:"14" solSign:"signed"`
	I15 *big.Int `solSize:"15" solSign:"signed"`
	I16 *big.Int `solSize:"16" solSign:"signed"`
	I17 *big.Int `solSize:"17" solSign:"signed"`
	I18 *big.Int `solSize:"18" solSign:"signed"`
	I19 *big.Int `solSize:"19" solSign:"signed"`
	I20 *big.Int `solSize:"20" solSign:"signed"`
	I21 *big.Int `solSize:"21" solSign:"signed"`
	I22 *big.Int `solSize:"22" solSign:"signed"`
	I23 *big.Int `solSize:"23" solSign:"signed"`
	I24 *big.Int `solSize:"24" solSign:"signed"`
	I25 *big.Int `solSize:"25" solSign:"signed"`
	I26 *big.Int `solSize:"26" solSign:"signed"`
	I27 *big.Int `solSize:"27" solSign:"signed"`
	I28 *big.Int `solSize:"28" solSign:"signed"`
	I29 *big.Int `solSize:"29" solSign:"signed"`
	I30 *big.Int `solSize:"30" solSign:"signed"`
	I31 *big.Int `solSize:"31" solSign:"signed"`
	I32 *big.Int `solSize:"32" solSign:"signed"`

	U9  *big.Int `solSize:"9" solSign:"unsigned"`
	U10 *big.Int `solSize:"10"`
	U11 *big.Int `solSize:"11"`
	U12 *big.Int `solSize:"12"`
	U13 *big.Int `solSize:"13"`
	U14 *big.Int `solSize:"14"`
	U15 *big.Int `solSize:"15"`
	U16 *big.Int `solSize:"16"`
	U17 *big.Int `solSize:"17"`
	U18 *big.Int `solSize:"18"`
	U19 *big.Int `solSize:"19"`
	U20 *big.Int `solSize:"20"`
	U21 *big.Int `solSize:"21"`
	U22 *big.Int `solSize:"22"`
	U23 *big.Int `solSize:"23"`
	U24 *big.Int `solSize:"24"`
	U25 *big.Int `solSize:"25"`
	U26 *big.Int `solSize:"26"`
	U27 *big.Int `solSize:"27"`
	U28 *big.Int `solSize:"28"`
	U29 *big.Int `solSize:"29"`
	U30 *big.Int `solSize:"30"`
	U31 *big.Int `solSize:"31"`
	U32 *big.Int `solSize:"32"`

	I32_2 *big.Int `solSign:"signed"`
	U32_2 *big.Int

	AddrU20 common.Address
}

var (
	big1   = big.NewInt(1)
	big256 = big.NewInt(256)
)

func mulShift(v *big.Int) *big.Int {
	r := new(big.Int).Sub(v, big1)
	r.Mul(r, big256)
	return r.Add(r, big1)
}

func TestContractStorageParserBigInts(t *testing.T) {
	t1, _ := new(big.Int).SetString("-549755813887", 10)
	t2, _ := new(big.Int).SetString("2147483649", 10)
	exp := dataBigInts{I9: t1, U9: t2}
	exp.I10 = mulShift(exp.I9)
	exp.I11 = mulShift(exp.I10)
	exp.I12 = mulShift(exp.I11)
	exp.I13 = mulShift(exp.I12)
	exp.I14 = mulShift(exp.I13)
	exp.I15 = mulShift(exp.I14)
	exp.I16 = mulShift(exp.I15)
	exp.I17 = mulShift(exp.I16)
	exp.I18 = mulShift(exp.I17)
	exp.I19 = mulShift(exp.I18)
	exp.I20 = mulShift(exp.I19)
	exp.I21 = mulShift(exp.I20)
	exp.I22 = mulShift(exp.I21)
	exp.I23 = mulShift(exp.I22)
	exp.I24 = mulShift(exp.I23)
	exp.I25 = mulShift(exp.I24)
	exp.I26 = mulShift(exp.I25)
	exp.I27 = mulShift(exp.I26)
	exp.I28 = mulShift(exp.I27)
	exp.I29 = mulShift(exp.I28)
	exp.I30 = mulShift(exp.I29)
	exp.I31 = mulShift(exp.I30)
	exp.I32 = mulShift(exp.I31)
	exp.U10 = mulShift(exp.U9)
	exp.U11 = mulShift(exp.U10)
	exp.U12 = mulShift(exp.U11)
	exp.U13 = mulShift(exp.U12)
	exp.U14 = mulShift(exp.U13)
	exp.U15 = mulShift(exp.U14)
	exp.U16 = mulShift(exp.U15)
	exp.U17 = mulShift(exp.U16)
	exp.U18 = mulShift(exp.U17)
	exp.U19 = mulShift(exp.U18)
	exp.U20 = mulShift(exp.U19)
	exp.U21 = mulShift(exp.U20)
	exp.U22 = mulShift(exp.U21)
	exp.U23 = mulShift(exp.U22)
	exp.U24 = mulShift(exp.U23)
	exp.U25 = mulShift(exp.U24)
	exp.U26 = mulShift(exp.U25)
	exp.U27 = mulShift(exp.U26)
	exp.U28 = mulShift(exp.U27)
	exp.U29 = mulShift(exp.U28)
	exp.U30 = mulShift(exp.U29)
	exp.U31 = mulShift(exp.U30)
	exp.U32 = mulShift(exp.U31)
	exp.I32_2 = exp.I32
	exp.U32_2 = exp.U32
	exp.AddrU20 = common.HexToAddress("ddea7d9bdc0a21b1e88788de4ce1fc89fcd17fd7")

	got := dataBigInts{}

	if err := parseFromTest("BigInts", &got); err != nil {
		t.Error(err)
		return
	}

	vExp := reflect.ValueOf(exp)
	vGot := reflect.ValueOf(got)
	for i := 0; i < 50; i++ {
		bExp := vExp.Field(i).Interface().(*big.Int)
		bGot := vGot.Field(i).Interface().(*big.Int)
		if bExp.Cmp(bGot) != 0 {
			t.Errorf("field %s has a different value than the expected.\nexpected: %v\ngot: %v\n", vExp.Type().Field(i).Name, bExp, bGot)
			return
		}
	}
	if !reflect.DeepEqual(exp.AddrU20, got.AddrU20) {
		t.Errorf("expected: %v\ngot: %v\n", exp.AddrU20, got.AddrU20)
		return
	}
}

type dataStrings struct {
	SmallString string
	BigString   string
}

func TestContractStorageParserStrings(t *testing.T) {
	exp := dataStrings{
		SmallString: "small string",
		BigString:   "this string is longer than 31 bytes",
	}

	got := dataStrings{}

	if err := parseFromTest("Strings", &got); err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(&exp, &got) {
		t.Errorf("expected: %v\ngot: %v\n", exp, got)
		return
	}
}

type smallStruct struct {
	Id    uint64
	Nonce uint32
}

type mediumStruct struct {
	Id   uint64
	Addr common.Address
}

type bigStruct struct {
	Id    *big.Int `solSize:"16"`
	Addr  common.Address
	Nonce *big.Int `solSize:"16"`
}

func (b bigStruct) Cmp(x bigStruct) bool {
	return b.Id.Cmp(x.Id) == 0 && b.Addr == x.Addr && b.Nonce.Cmp(x.Nonce) == 0
}

type dataArrays struct {
	Owners        [3]common.Address
	Votes         []common.Address
	SmallFixed    [3]smallStruct
	SmallDynamic  []smallStruct
	MediumFixed   [3]mediumStruct
	MediumDynamic []mediumStruct
	BigFixed      [3]bigStruct
	BigDynamic    []bigStruct
}

func TestContractStorageParserArrays(t *testing.T) {
	exp := dataArrays{
		Owners: [3]common.Address{
			common.HexToAddress("0xe92a2a4e3f4c378495145619f2975ce8c60819c2"),
			common.HexToAddress("0x14dd8d9c759a6827aacbf726085ef13a357989ec"),
			common.HexToAddress("0xa1f0a100522350ee2a044fe69831cf469c0f7123"),
		},
	}
	exp.Votes = exp.Owners[:]
	for i := 0; i < 3; i++ {
		s := smallStruct{uint64(i), uint32(i + 1)}
		exp.SmallFixed[i] = s
		exp.SmallDynamic = append(exp.SmallDynamic, s)
		m := mediumStruct{uint64(i), exp.Owners[i]}
		exp.MediumFixed[i] = m
		exp.MediumDynamic = append(exp.MediumDynamic, m)
		b := bigStruct{big.NewInt(int64(i)), exp.Owners[i], big.NewInt(int64(i) * 256)}
		exp.BigFixed[i] = b
		exp.BigDynamic = append(exp.BigDynamic, b)
	}
	got := dataArrays{}

	if err := parseFromTest("Arrays", &got); err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(exp.Owners, got.Owners) {
		t.Errorf("expected: %v\ngot: %v\n", exp, got)
		return
	}
	if !reflect.DeepEqual(exp.Votes, got.Votes) {
		t.Errorf("expected: %v\ngot: %v\n", exp, got)
		return
	}
	if !reflect.DeepEqual(exp.SmallFixed, got.SmallFixed) {
		t.Errorf("expected: %v\ngot: %v\n", exp, got)
		return
	}
	if !reflect.DeepEqual(exp.SmallDynamic, got.SmallDynamic) {
		t.Errorf("expected: %v\ngot: %v\n", exp, got)
		return
	}
	if !reflect.DeepEqual(exp.MediumFixed, got.MediumFixed) {
		t.Errorf("expected: %v\ngot: %v\n", exp, got)
		return
	}
	if !reflect.DeepEqual(exp.MediumDynamic, got.MediumDynamic) {
		t.Errorf("expected: %v\ngot: %v\n", exp, got)
		return
	}
	cmpBigStruct := func(bs1, bs2 *bigStruct) bool {
		if !reflect.DeepEqual(bs1.Addr, bs2.Addr) {
			return false
		}
		if bs1.Id.Cmp(bs2.Id) != 0 {
			return false
		}
		if bs1.Nonce.Cmp(bs2.Nonce) != 0 {
			return false
		}
		return true
	}
	for i := 0; i < 3; i++ {
		if !cmpBigStruct(&exp.BigFixed[i], &got.BigFixed[i]) {
			t.Errorf("expected: %v\ngot: %v\n", exp.BigFixed[i], got.BigFixed[i])
			return
		}
		if !cmpBigStruct(&exp.BigDynamic[i], &got.BigDynamic[i]) {
			t.Errorf("expected: %v\ngot: %v\n", exp.BigDynamic[i], got.BigDynamic[i])
			return
		}
	}
}

type dataMappings struct {
	IdAddr       map[uint64]common.Address
	AddrsSmall   map[common.Address]smallStruct
	BigKeys      map[*big.Int]bigStruct
	StringMedium map[string]mediumStruct
}

func TestContractStorageParserMappings(t *testing.T) {
	exp := dataMappings{
		IdAddr: map[uint64]common.Address{
			0: common.HexToAddress("0xe92a2a4e3f4c378495145619f2975ce8c60819c2"),
			1: common.HexToAddress("0x14dd8d9c759a6827aacbf726085ef13a357989ec"),
			2: common.HexToAddress("0xa1f0a100522350ee2a044fe69831cf469c0f7123"),
		},
		AddrsSmall:   make(map[common.Address]smallStruct, 3),
		BigKeys:      make(map[*big.Int]bigStruct, 3),
		StringMedium: make(map[string]mediumStruct, 3),
	}
	for id, addr := range exp.IdAddr {
		exp.AddrsSmall[addr] = smallStruct{Id: id, Nonce: uint32(id) + 1}
		exp.BigKeys[big.NewInt(int64(id))] = bigStruct{Id: big.NewInt(int64(id)), Addr: exp.IdAddr[id], Nonce: big.NewInt(int64(id) * 256)}
		exp.StringMedium["small string"] = mediumStruct{Id: 0, Addr: exp.IdAddr[0]}
		exp.StringMedium["still a small string"] = mediumStruct{Id: 1, Addr: exp.IdAddr[1]}
		exp.StringMedium["a big string must be longer than 31 bytes"] = mediumStruct{Id: 2, Addr: exp.IdAddr[2]}
	}
	got := struct {
		IdAddr       *state.Mapping
		AddrsSmall   *state.Mapping
		BigKeys      *state.Mapping
		StringMedium *state.Mapping
	}{}
	if err := parseFromTest("Mappings", &got); err != nil {
		t.Error(err)
		return
	}
	if err := compareMaps(exp.IdAddr, got.IdAddr); err != nil {
		t.Error(err)
		return
	}
	if err := compareMaps(exp.AddrsSmall, got.AddrsSmall); err != nil {
		t.Error(err)
		return
	}
	if err := compareMaps(exp.StringMedium, got.StringMedium); err != nil {
		t.Error(err)
		return
	}
	for k, v := range exp.BigKeys {
		bs := &bigStruct{}
		if err := got.BigKeys.Get(k, bs); err != nil {
			t.Error(err)
			return
		}
		if !v.Cmp(*bs) {
			t.Errorf("expected: %#+v\ngot: %#+v", v, bs)
			return
		}
	}
}

func compareMaps(m interface{}, sm *state.Mapping) error {
	mv := reflect.ValueOf(m)
	if mv.Kind() != reflect.Map {
		return fmt.Errorf("need a map")
	}
	for _, k := range mv.MapKeys() {
		mVal := mv.MapIndex(k)
		gotV := reflect.New(mVal.Type())
		sm.Get(k.Interface(), gotV.Interface())
		if !reflect.DeepEqual(mVal.Interface(), gotV.Elem().Interface()) {
			return fmt.Errorf("got a mismatch")
		}
	}
	return nil
}
