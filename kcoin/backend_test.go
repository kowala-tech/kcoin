package kcoin

import (
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/kcoindb"
	"github.com/kowala-tech/kcoin/params"
)

func TestMipmapUpgrade(t *testing.T) {
	db, _ := kcoindb.NewMemDatabase()
	addr := common.BytesToAddress([]byte("jeff"))
	genesis := new(core.Genesis).MustCommit(db)

	chain, receipts := core.GenerateChain(params.TestChainConfig, genesis, db, 10, func(i int, gen *core.BlockGen) {
		switch i {
		case 1:
			receipt := types.NewReceipt(nil, new(big.Int))
			receipt.Logs = []*types.Log{{Address: addr}}
			gen.AddUncheckedReceipt(receipt)
		case 2:
			receipt := types.NewReceipt(nil, new(big.Int))
			receipt.Logs = []*types.Log{{Address: addr}}
			gen.AddUncheckedReceipt(receipt)
		}
	})
	for i, block := range chain {
		core.WriteBlock(db, block)
		if err := core.WriteCanonicalHash(db, block.Hash(), block.NumberU64()); err != nil {
			t.Fatalf("failed to insert block number: %v", err)
		}
		if err := core.WriteHeadBlockHash(db, block.Hash()); err != nil {
			t.Fatalf("failed to insert block number: %v", err)
		}
		if err := core.WriteBlockReceipts(db, block.Hash(), block.NumberU64(), receipts[i]); err != nil {
			t.Fatal("error writing block receipts:", err)
		}
	}

	err := addMipmapBloomBins(db)
	if err != nil {
		t.Fatal(err)
	}

	bloom := core.GetMipmapBloom(db, 1, core.MIPMapLevels[0])
	if (bloom == types.Bloom{}) {
		t.Error("got empty bloom filter")
	}

	data, _ := db.Get([]byte("setting-mipmap-version"))
	if len(data) == 0 {
		t.Error("setting-mipmap-version not written to database")
	}
}
