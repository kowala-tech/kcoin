package tendermint

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/require"
)

func TestMintedAmount(t *testing.T) {
	testCases := []struct {
		blockNumber      *big.Int
		currentPrice     *big.Int
		prevPrice        *big.Int
		prevSupply       *big.Int
		prevMintedAmount *big.Int
		mintedAmount     *big.Int
	}{
		{
			blockNumber:      common.Big1,
			currentPrice:     new(big.Int),
			prevPrice:        new(big.Int),
			prevSupply:       new(big.Int),
			prevMintedAmount: new(big.Int),
			mintedAmount:     initialMintedAmount,
		},
		{
			blockNumber:      common.Big2,
			currentPrice:     new(big.Int).Mul(common.Big3, big.NewInt(params.Kcoin)),
			prevPrice:        new(big.Int).Mul(common.Big2, big.NewInt(params.Kcoin)),
			prevSupply:       new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			prevMintedAmount: new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			mintedAmount:     big.NewInt(10001e14),
		},
		{
			blockNumber:      common.Big2,
			currentPrice:     new(big.Int).Mul(common.Big3, big.NewInt(params.Kcoin)),
			prevPrice:        new(big.Int).Mul(common.Big2, big.NewInt(params.Kcoin)),
			prevSupply:       new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			prevMintedAmount: new(big.Int).Mul(new(big.Int).SetUint64(83), big.NewInt(params.Kcoin)),
			mintedAmount:     new(big.Int).Mul(new(big.Int).SetUint64(82), big.NewInt(params.Kcoin)),
		},
		{
			blockNumber:      common.Big2,
			currentPrice:     new(big.Int).Mul(common.Big3, big.NewInt(params.Kcoin)),
			prevPrice:        new(big.Int).Mul(common.Big2, big.NewInt(params.Kcoin)),
			prevSupply:       new(big.Int).Mul(new(big.Int).SetUint64(1000000), big.NewInt(params.Kcoin)),
			prevMintedAmount: new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			mintedAmount:     big.NewInt(10001e14),
		},
		{
			blockNumber:      common.Big2,
			currentPrice:     new(big.Int).Mul(common.Big3, big.NewInt(params.Kcoin)),
			prevPrice:        new(big.Int).Mul(common.Big2, big.NewInt(params.Kcoin)),
			prevSupply:       new(big.Int).Mul(new(big.Int).SetUint64(1000000), big.NewInt(params.Kcoin)),
			prevMintedAmount: new(big.Int).Mul(new(big.Int).SetUint64(101), big.NewInt(params.Kcoin)),
			mintedAmount:     new(big.Int).Mul(new(big.Int).SetUint64(100), big.NewInt(params.Kcoin)),
		},
		{
			blockNumber:      common.Big2,
			currentPrice:     new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			prevPrice:        new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			prevSupply:       new(big.Int),
			prevMintedAmount: big.NewInt(1e11),
			mintedAmount:     big.NewInt(1e12),
		},
		{
			blockNumber:      common.Big2,
			currentPrice:     new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			prevPrice:        new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			prevSupply:       new(big.Int),
			prevMintedAmount: big.NewInt(1e14),
			mintedAmount:     new(big.Int).Mul(new(big.Int).SetUint64(9999), big.NewInt(1e10)),
		},
		{
			blockNumber:      common.Big2,
			currentPrice:     new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			prevPrice:        new(big.Int).Mul(common.Big2, big.NewInt(params.Kcoin)),
			prevSupply:       new(big.Int),
			prevMintedAmount: big.NewInt(1e11),
			mintedAmount:     big.NewInt(1e12),
		},
		{
			blockNumber:      common.Big2,
			currentPrice:     new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin)),
			prevPrice:        new(big.Int).Mul(common.Big2, big.NewInt(params.Kcoin)),
			prevSupply:       new(big.Int),
			prevMintedAmount: big.NewInt(1e14),
			mintedAmount:     new(big.Int).Mul(new(big.Int).SetUint64(9999), big.NewInt(1e10)),
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("block number %d current price %d previous price %d previous supply %d previous minted amount %d", tc.blockNumber.Uint64(), tc.currentPrice.Uint64(), tc.prevPrice.Uint64(), tc.prevSupply.Uint64(), tc.prevMintedAmount.Uint64()), func(t *testing.T) {
			require.Equal(t, tc.mintedAmount, mintedAmount(tc.blockNumber, tc.currentPrice, tc.prevPrice, tc.prevSupply, tc.prevMintedAmount))
		})
	}

}

func TestCap(t *testing.T) {
	testCases := []struct {
		blockNumber *big.Int
		prevSupply  *big.Int
		cap         *big.Int
	}{
		{
			blockNumber: common.Big1,
			prevSupply:  new(big.Int),
			cap:         new(big.Int).Mul(new(big.Int).SetUint64(82), big.NewInt(params.Kcoin)),
		},
		{
			blockNumber: common.Big2,
			prevSupply:  new(big.Int).Mul(new(big.Int).SetUint64(999999), big.NewInt(params.Kcoin)),
			cap:         new(big.Int).Mul(new(big.Int).SetUint64(82), big.NewInt(params.Kcoin)),
		},
		{
			blockNumber: common.Big2,
			prevSupply:  new(big.Int).Mul(new(big.Int).SetUint64(1000000), big.NewInt(params.Kcoin)),
			cap:         new(big.Int).Mul(new(big.Int).SetUint64(100), big.NewInt(params.Kcoin)),
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("block number %d previous supply %d", tc.blockNumber.Uint64(), tc.prevSupply.Uint64()), func(t *testing.T) {
			require.Equal(t, tc.cap, cap(tc.blockNumber, tc.prevSupply))
		})
	}
}

func TestHasLowSupply(t *testing.T) {
	testCases := []struct {
		supply       *big.Int
		hasLowSupply bool
	}{
		{
			supply:       new(big.Int).Mul(new(big.Int).SetUint64(1000000), big.NewInt(params.Kcoin)),
			hasLowSupply: false,
		},
		{
			supply:       new(big.Int).Mul(new(big.Int).SetUint64(999999), big.NewInt(params.Kcoin)),
			hasLowSupply: true,
		},
		{
			supply:       new(big.Int).Mul(new(big.Int).SetUint64(1000001), big.NewInt(params.Kcoin)),
			hasLowSupply: false,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("supply %d", tc.supply.Uint64()), func(t *testing.T) {
			require.Equal(t, tc.hasLowSupply, hasLowSupply(tc.supply))
		})
	}
}
