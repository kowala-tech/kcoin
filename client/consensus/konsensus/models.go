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

		// reset oracle mgr
		oracleMgr := &OracleMgr{
			AveragePrice: common.Big0
			
		}

		// reset hasSubmittedPrice per author
		keccak := sha3.NewKeccak256()
		participants, err := sys.provider.Submissions()
		if err != nil {
			return err
		}
		for _, oracle := range participants {
			// oracle contract key (oracleRegistry)
			keccak.Write(oracle.Bytes())
			keccak.Write(hasSubmittedPriceIdx.Bytes())
			key := common.BytesToHash(keccak.Sum(nil))
			keccak.Reset()

			// reset hasSubmittedPrice per submission
			sys.SetState(sys.provider.Address(), key, common.BytesToHash([]byte{0}))

			// reset submissions entry
			sys.SetState(sys.provider.Address(), key, common.BytesToHash([]byte{}))
		}
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
