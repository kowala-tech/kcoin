package konsensus

import (
	"math/big"
	"reflect"
	"strings"

	"github.com/kowala-tech/kcoin/client/common"
)

type Storage interface {
	Domain() string
}

type BeforeUpdate interface {
	BeforeUpdate(fieldName string)
}

type SystemVars struct {
	PrevCurrencyPrice *big.Int
	CurrencyPrice     *big.Int
	MintedReward      *big.Int
	CurrencySupply    *big.Int
}

func (vars *SystemVars) BeforeUpdate(fieldName string) {
	switch field := strings.ToLower(fieldName); field {
	case "currencyprice":
		vars.PrevCurrencyPrice = vars.CurrencyPrice
	}
}

func (vars *SystemVars) Domain() string {
	return TypeNameLower(*vars)
}

func TypeNameLower(typ interface{}) string {
	return strings.ToLower(reflect.TypeOf(typ).Name())
}

type Oracle struct {
}

type OracleMgr struct {
	Paused         bool
	BaseDeposit    *big.Int
	MaxNumOracles  *big.Int
	FreezePeriod   *big.Int
	SyncFrequency  *big.Int
	UpdatePeriod   *big.Int
	AveragePrice   *big.Int
	OracleRegistry map[common.Address]Oracle
	oraclePool     []common.Address
}

func (mgr *OracleMgr) Domain() string {
	return TypeNameLower(*mgr)
}
