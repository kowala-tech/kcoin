package state_test

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"io/ioutil"
	"math/big"
	"path"
	"strings"
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

func newKeyAddr() (*ecdsa.PrivateKey, common.Address, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, err
	}
	return key, crypto.PubkeyToAddress(key.PublicKey), nil
}

func parseAddress(a string) common.Address {
	a = strings.ToLower(strings.TrimSpace(a))
	if strings.HasPrefix(a, "0x") {
		a = strings.TrimPrefix(a, "0x")
	}
	b, err := hex.DecodeString(a)
	if err != nil {
		panic(err)
	}
	return common.BytesToAddress(b)
}

func TestContractStorageParser(t *testing.T) {
	// read contract's bytecode (hex)
	contractHexBytecode, err := ioutil.ReadFile(path.Join("test_contract", "LocalStorageTest.bin"))
	if err != nil {
		t.Error(err)
		return
	}
	// convert
	contractBytecode := make([]byte, len(contractHexBytecode)/2)
	if _, err = hex.Decode(contractBytecode, contractHexBytecode); err != nil {
		t.Error(err)
		return
	}
	// read contract ABI
	b, err := ioutil.ReadFile(path.Join("test_contract", "LocalStorageTest.abi"))
	if err != nil {
		t.Error(err)
		return
	}
	// read json abi
	contractABI, err := abi.JSON(bytes.NewBuffer(b))
	if err != nil {
		t.Error(err)
		return
	}
	// generate a new key
	privKey, privAddr, err := newKeyAddr()
	if err != nil {
		t.Error(err)
		return
	}
	// create a new simulated backend
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		privAddr: core.GenesisAccount{
			Balance: new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(params.Ether)),
		},
	})
	// deploy contract
	contractAddr, contractTx, _, err := bind.DeployContract(
		bind.NewKeyedTransactor(privKey),
		contractABI,
		contractBytecode,
		sim,
	)
	if err != nil {
		t.Error(err)
		return
	}
	sim.Commit()
	// parse the contract storage
	type SmallStruct struct {
		id    uint64
		nonce uint32
	}
	type MediumStruct struct {
		id   uint64
		addr common.Address
	}
	type BigStruct struct {
		id    *big.Int `solSize:"16"`
		addr  common.Address
		nonce *big.Int `solSize:"16"`
	}
	type ContractData struct {
		AA            int8
		AB            uint8
		AC            int16
		AD            uint16
		AE            int32  `solSize:"3"`
		AF            uint32 `solSize:"3"`
		AG            int32
		AH            uint32
		AI            int64  `solSize:"5"`
		AJ            uint64 `solSize:"5"`
		AK            int64  `solSize:"6"`
		AL            uint64 `solSize:"6"`
		AM            int64  `solSize:"7"`
		AN            uint64 `solSize:"7"`
		AO            int64
		AP            uint64
		AQ            *big.Int `solSize:"9" solSign:"signed"`
		AR            *big.Int `solSize:"9" solSign:"unsigned"`
		AT            *big.Int `solSize:"10" solSign:"signed"`
		AU            *big.Int `solSize:"10"`
		AV            *big.Int `solSize:"11" solSign:"signed"`
		AW            *big.Int `solSize:"11"`
		AX            *big.Int `solSize:"12" solSign:"signed"`
		AY            *big.Int `solSize:"12"`
		AZ            *big.Int `solSize:"13" solSign:"signed"`
		BA            *big.Int `solSize:"13"`
		BB            *big.Int `solSize:"14" solSign:"signed"`
		BC            *big.Int `solSize:"14"`
		BD            *big.Int `solSize:"15" solSign:"signed"`
		BE            *big.Int `solSize:"15"`
		BF            *big.Int `solSize:"16" solSign:"signed"`
		BG            *big.Int `solSize:"16"`
		BH            *big.Int `solSize:"17" solSign:"signed"`
		BI            *big.Int `solSize:"17"`
		BJ            *big.Int `solSize:"18" solSign:"signed"`
		BK            *big.Int `solSize:"18"`
		BL            *big.Int `solSize:"19" solSign:"signed"`
		BM            *big.Int `solSize:"19"`
		BN            *big.Int `solSize:"20" solSign:"signed"`
		BO            *big.Int `solSize:"20"`
		BP            *big.Int `solSize:"21" solSign:"signed"`
		BQ            *big.Int `solSize:"21"`
		BR            *big.Int `solSize:"22" solSign:"signed"`
		BS            *big.Int `solSize:"22"`
		BT            *big.Int `solSize:"23" solSign:"signed"`
		BU            *big.Int `solSize:"23"`
		BV            *big.Int `solSize:"24" solSign:"signed"`
		BW            *big.Int `solSize:"24"`
		BX            *big.Int `solSize:"25" solSign:"signed"`
		BY            *big.Int `solSize:"25"`
		BZ            *big.Int `solSize:"26" solSign:"signed"`
		CA            *big.Int `solSize:"26"`
		CB            *big.Int `solSize:"27" solSign:"signed"`
		CC            *big.Int `solSize:"27"`
		CD            *big.Int `solSize:"28" solSign:"signed"`
		CE            *big.Int `solSize:"28"`
		CF            *big.Int `solSize:"29" solSign:"signed"`
		CG            *big.Int `solSize:"29"`
		CH            *big.Int `solSize:"30" solSign:"signed"`
		CI            *big.Int `solSize:"30"`
		CJ            *big.Int `solSize:"31" solSign:"signed"`
		CK            *big.Int `solSize:"31"`
		CL            *big.Int `solSize:"32" solSign:"signed"`
		CM            *big.Int `solSize:"32"`
		CN            *big.Int `solSign:"signed"`
		CO            *big.Int
		CP            bool
		CQ            *common.Address
		MediumFixed   [2]MediumStruct
		MediumDynamic []MediumStruct
		SmallString   string
		BigString     string
		MappingSimple *state.Mapping
		BytesFixed    [50]byte
		BytesDynamic  []byte
		SmallFixed    [4]SmallStruct
		SmallDynamic  []SmallStruct
		BigFixed      [2]BigStruct
		BigDynamic    []BigStruct
		MappingBig    *state.Mapping
		FinalBit      bool
	}
	bigMinusOne := big.NewInt(-1)
	a1 := parseAddress("0xddea7d9bdc0a21b1e88788de4ce1fc89fcd17fd7")
	expectedContractData := ContractData{
		AA:          -100,
		AB:          100,
		CP:          true,
		CQ:          &a1,
		SmallString: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		BigString:   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		FinalBit:    true,
		SmallFixed: [4]SmallStruct{
			{1, 1 + 1},
			{2, 2 + 1},
			{3, 3 + 1},
			{4, 4 + 1},
		},
		SmallDynamic: []SmallStruct{
			{5, 5 + 1},
			{6, 6 + 1},
		},
		MediumFixed: [2]MediumStruct{
			{1, parseAddress("0xe92a2a4e3f4c378495145619f2975ce8c60819c2")},
			{2, parseAddress("0x14dd8d9c759a6827aacbf726085ef13a357989ec")},
		},
		MediumDynamic: []MediumStruct{
			{5, parseAddress("0x97e5f97782770d049cfd2e8dff61393ef090656c")},
			{6, parseAddress("0xd289e02286e7dd7d6ac464e72f01b87c9678a8a5")},
			{7, parseAddress("0x73789ad2d8db1d51730ab0eba9d8de102780c069")},
		},
		BigFixed: [2]BigStruct{
			{big.NewInt(1), parseAddress("0xa1f0a100522350ee2a044fe69831cf469c0f7123"), big.NewInt(1 + 1)},
			{big.NewInt(2), parseAddress("0xba41414996cf0127641366eca386318a5be3fd7b"), big.NewInt(2 + 1)},
		},
		BigDynamic: []BigStruct{
			{big.NewInt(5), parseAddress("0xd3f7a7d9d33e53cb1c4591d644005300156ccf2a"), big.NewInt(5 + 1)},
			{big.NewInt(6), parseAddress("0x58e10587afbade73778e33f43ef28f5f48df17cd"), big.NewInt(6 + 1)},
			{big.NewInt(7), parseAddress("0xc6985e85b1791e84bbb8323b6af94e18d7d502b2"), big.NewInt(7 + 1)},
			{big.NewInt(8), parseAddress("0x538aecfa3c8c06f2944a57221bbc8d3fb9695a32"), big.NewInt(8 + 1)},
		},
	}
	expectedContractData.AC = int16(expectedContractData.AA) * 256
	expectedContractData.AD = uint16(expectedContractData.AC * -1)
	expectedContractData.AE = int32(expectedContractData.AC * 256)
	expectedContractData.AF = uint32(expectedContractData.AE * -1)
	expectedContractData.AG = int32(expectedContractData.AE * 256)
	expectedContractData.AH = uint32(expectedContractData.AG * -1)
	expectedContractData.AI = int64(expectedContractData.AG * 256)
	expectedContractData.AJ = uint64(expectedContractData.AI * -1)
	expectedContractData.AK = int64(expectedContractData.AI * 256)
	expectedContractData.AL = uint64(expectedContractData.AK * -1)
	expectedContractData.AM = int64(expectedContractData.AK * 256)
	expectedContractData.AN = uint64(expectedContractData.AM * -1)
	expectedContractData.AO = int64(expectedContractData.AM * 256)
	expectedContractData.AP = uint64(expectedContractData.AO * -1)
	expectedContractData.AQ = new(big.Int).SetInt64(expectedContractData.AO)
	expectedContractData.AQ.Mul(expectedContractData.AQ, common.Big256)
	expectedContractData.AR = new(big.Int).Set(expectedContractData.AQ)
	expectedContractData.AR.Mul(expectedContractData.AR, bigMinusOne)
	expectedContractData.AT = new(big.Int).Set(expectedContractData.AQ)
	expectedContractData.AT.Mul(expectedContractData.AT, common.Big256)
	expectedContractData.AU = new(big.Int).Set(expectedContractData.AT)
	expectedContractData.AU.Mul(expectedContractData.AU, bigMinusOne)
	expectedContractData.AV = new(big.Int).Set(expectedContractData.AT)
	expectedContractData.AV.Mul(expectedContractData.AV, common.Big256)
	expectedContractData.AW = new(big.Int).Set(expectedContractData.AV)
	expectedContractData.AW.Mul(expectedContractData.AW, bigMinusOne)
	expectedContractData.AX = new(big.Int).Set(expectedContractData.AV)
	expectedContractData.AX.Mul(expectedContractData.AX, common.Big256)
	expectedContractData.AY = new(big.Int).Set(expectedContractData.AX)
	expectedContractData.AY.Mul(expectedContractData.AY, bigMinusOne)
	expectedContractData.AZ = new(big.Int).Set(expectedContractData.AX)
	expectedContractData.AZ.Mul(expectedContractData.AZ, common.Big256)
	expectedContractData.BA = new(big.Int).Set(expectedContractData.AZ)
	expectedContractData.BA.Mul(expectedContractData.BA, bigMinusOne)
	expectedContractData.BB = new(big.Int).Set(expectedContractData.AZ)
	expectedContractData.BB.Mul(expectedContractData.BB, common.Big256)
	expectedContractData.BC = new(big.Int).Set(expectedContractData.BB)
	expectedContractData.BC.Mul(expectedContractData.BC, bigMinusOne)
	expectedContractData.BD = new(big.Int).Set(expectedContractData.BB)
	expectedContractData.BD.Mul(expectedContractData.BD, common.Big256)
	expectedContractData.BE = new(big.Int).Set(expectedContractData.BD)
	expectedContractData.BE.Mul(expectedContractData.BE, bigMinusOne)
	expectedContractData.BF = new(big.Int).Set(expectedContractData.BD)
	expectedContractData.BF.Mul(expectedContractData.BF, common.Big256)
	expectedContractData.BG = new(big.Int).Set(expectedContractData.BF)
	expectedContractData.BG.Mul(expectedContractData.BG, bigMinusOne)
	expectedContractData.BH = new(big.Int).Set(expectedContractData.BF)
	expectedContractData.BH.Mul(expectedContractData.BH, common.Big256)
	expectedContractData.BI = new(big.Int).Set(expectedContractData.BH)
	expectedContractData.BI.Mul(expectedContractData.BI, bigMinusOne)
	expectedContractData.BJ = new(big.Int).Set(expectedContractData.BH)
	expectedContractData.BJ.Mul(expectedContractData.BJ, common.Big256)
	expectedContractData.BK = new(big.Int).Set(expectedContractData.BJ)
	expectedContractData.BK.Mul(expectedContractData.BK, bigMinusOne)
	expectedContractData.BL = new(big.Int).Set(expectedContractData.BJ)
	expectedContractData.BL.Mul(expectedContractData.BL, common.Big256)
	expectedContractData.BM = new(big.Int).Set(expectedContractData.BL)
	expectedContractData.BM.Mul(expectedContractData.BM, bigMinusOne)
	expectedContractData.BN = new(big.Int).Set(expectedContractData.BL)
	expectedContractData.BN.Mul(expectedContractData.BN, common.Big256)
	expectedContractData.BO = new(big.Int).Set(expectedContractData.BN)
	expectedContractData.BO.Mul(expectedContractData.BO, bigMinusOne)
	expectedContractData.BP = new(big.Int).Set(expectedContractData.BN)
	expectedContractData.BP.Mul(expectedContractData.BP, common.Big256)
	expectedContractData.BQ = new(big.Int).Set(expectedContractData.BP)
	expectedContractData.BQ.Mul(expectedContractData.BQ, bigMinusOne)
	expectedContractData.BR = new(big.Int).Set(expectedContractData.BP)
	expectedContractData.BR.Mul(expectedContractData.BR, common.Big256)
	expectedContractData.BS = new(big.Int).Set(expectedContractData.BR)
	expectedContractData.BS.Mul(expectedContractData.BS, bigMinusOne)
	expectedContractData.BT = new(big.Int).Set(expectedContractData.BR)
	expectedContractData.BT.Mul(expectedContractData.BT, common.Big256)
	expectedContractData.BU = new(big.Int).Set(expectedContractData.BT)
	expectedContractData.BU.Mul(expectedContractData.BU, bigMinusOne)
	expectedContractData.BV = new(big.Int).Set(expectedContractData.BT)
	expectedContractData.BV.Mul(expectedContractData.BV, common.Big256)
	expectedContractData.BW = new(big.Int).Set(expectedContractData.BV)
	expectedContractData.BW.Mul(expectedContractData.BW, bigMinusOne)
	expectedContractData.BX = new(big.Int).Set(expectedContractData.BV)
	expectedContractData.BX.Mul(expectedContractData.BX, common.Big256)
	expectedContractData.BY = new(big.Int).Set(expectedContractData.BX)
	expectedContractData.BY.Mul(expectedContractData.BY, bigMinusOne)
	expectedContractData.BZ = new(big.Int).Set(expectedContractData.BX)
	expectedContractData.BZ.Mul(expectedContractData.BZ, common.Big256)
	expectedContractData.CA = new(big.Int).Set(expectedContractData.BZ)
	expectedContractData.CA.Mul(expectedContractData.CA, bigMinusOne)
	expectedContractData.CB = new(big.Int).Set(expectedContractData.BZ)
	expectedContractData.CB.Mul(expectedContractData.CB, common.Big256)
	expectedContractData.CC = new(big.Int).Set(expectedContractData.CB)
	expectedContractData.CC.Mul(expectedContractData.CC, bigMinusOne)
	expectedContractData.CD = new(big.Int).Set(expectedContractData.CB)
	expectedContractData.CD.Mul(expectedContractData.CD, common.Big256)
	expectedContractData.CE = new(big.Int).Set(expectedContractData.CD)
	expectedContractData.CE.Mul(expectedContractData.CE, bigMinusOne)
	expectedContractData.CF = new(big.Int).Set(expectedContractData.CD)
	expectedContractData.CF.Mul(expectedContractData.CF, common.Big256)
	expectedContractData.CG = new(big.Int).Set(expectedContractData.CF)
	expectedContractData.CG.Mul(expectedContractData.CG, bigMinusOne)
	expectedContractData.CH = new(big.Int).Set(expectedContractData.CF)
	expectedContractData.CH.Mul(expectedContractData.CH, common.Big256)
	expectedContractData.CI = new(big.Int).Set(expectedContractData.CH)
	expectedContractData.CI.Mul(expectedContractData.CI, bigMinusOne)
	expectedContractData.CJ = new(big.Int).Set(expectedContractData.CH)
	expectedContractData.CJ.Mul(expectedContractData.CJ, common.Big256)
	expectedContractData.CK = new(big.Int).Set(expectedContractData.CJ)
	expectedContractData.CK.Mul(expectedContractData.CK, bigMinusOne)
	expectedContractData.CL = new(big.Int).Set(expectedContractData.CJ)
	expectedContractData.CL.Mul(expectedContractData.CL, common.Big256)
	expectedContractData.CM = new(big.Int).Set(expectedContractData.CL)
	expectedContractData.CM.Mul(expectedContractData.CM, bigMinusOne)
	expectedContractData.CN = new(big.Int).Set(expectedContractData.CL)
	expectedContractData.CN.Mul(expectedContractData.CN, common.Big256)
	expectedContractData.CO = new(big.Int).Set(expectedContractData.CN)
	expectedContractData.CO.Mul(expectedContractData.CO, bigMinusOne)

	_, _ = contractAddr, contractTx
}
