package main

import genesisgen "github.com/kowala-tech/kcoin/knode/genesis"

var genesisOptions = map[string]map[string]genesisgen.Options{
	"kusd": {
		"main": genesisgen.Options{
			Network:   genesisgen.MainNetwork,
			ExtraData: "Kowala's first block",
			Governance: &genesisgen.GovernanceOpts{
				Origin: "0x259be75d96876f2ada3d202722523e9cd4dd917d",
				Governors: []string{
					"0xa1e8587ed7f915d5bbbf283b21af4813232069f7",
					"0xbfAdCF85554F139F978DE5442aacFBe085c754f7",
					"0xF358eb1020375800746ccd5c6638DA36C5a6bec9",
				},
				NumConfirmations: 2,
			},
			Consensus: &genesisgen.ConsensusOpts{
				Engine:           genesisgen.TendermintConsensus,
				MaxNumValidators: 100,
				FreezePeriod:     1,
				BaseDeposit:      1000000,
				Validators: []genesisgen.Validator{
					{
						Address: "0xd6e579085c82329c89fca7a9f012be59028ed53f",
						Deposit: 1000000,
					},
				},
				MiningToken: &genesisgen.MiningTokenOpts{
					Name:     "mUSD",
					Symbol:   "mUSD",
					Cap:      1073741824,
					Decimals: 18,
					Holders: []genesisgen.TokenHolder{
						{
							Address:   "0xd6e579085c82329c89fca7a9f012be59028ed53f",
							NumTokens: 3000000,
						},
					},
				},
			},
			DataFeedSystem: &genesisgen.DataFeedSystemOpts{
				MaxNumOracles: 1000,
				FreezePeriod:  1,
				BaseDeposit:   10,
			},
			PrefundedAccounts: []genesisgen.PrefundedAccount{
				{
					Address: "0xa1e8587ed7f915d5bbbf283b21af4813232069f7",
					Balance: 50,
				},
				{
					Address: "0xbfAdCF85554F139F978DE5442aacFBe085c754f7",
					Balance: 50,
				},
				{
					Address: "0xF358eb1020375800746ccd5c6638DA36C5a6bec9",
					Balance: 50,
				},
				{
					Address: "0xd6e579085c82329c89fca7a9f012be59028ed53f",
					Balance: 1000000,
				},
			},
		},
		"testnet": genesisgen.Options{
			Network:   genesisgen.TestNetwork,
			ExtraData: "Kowala's first block",
			Governance: &genesisgen.GovernanceOpts{
				Origin: "0x259be75d96876f2ada3d202722523e9cd4dd917d",
				Governors: []string{
					"0xa1e8587ed7f915d5bbbf283b21af4813232069f7",
					"0xbfAdCF85554F139F978DE5442aacFBe085c754f7",
					"0xF358eb1020375800746ccd5c6638DA36C5a6bec9",
				},
				NumConfirmations: 2,
			},
			Consensus: &genesisgen.ConsensusOpts{
				Engine:           genesisgen.TendermintConsensus,
				MaxNumValidators: 100,
				FreezePeriod:     1,
				BaseDeposit:      1000000,
				Validators: []genesisgen.Validator{
					{
						Address: "0xd6e579085c82329c89fca7a9f012be59028ed53f",
						Deposit: 1000000,
					},
				},
				MiningToken: &genesisgen.MiningTokenOpts{
					Name:     "mUSD",
					Symbol:   "mUSD",
					Cap:      1073741824,
					Decimals: 18,
					Holders: []genesisgen.TokenHolder{
						{
							Address:   "0xd6e579085c82329c89fca7a9f012be59028ed53f",
							NumTokens: 3000000,
						},
					},
				},
			},
			DataFeedSystem: &genesisgen.DataFeedSystemOpts{
				MaxNumOracles: 1000,
				FreezePeriod:  1,
				BaseDeposit:   10,
			},
			PrefundedAccounts: []genesisgen.PrefundedAccount{
				{
					Address: "0xa1e8587ed7f915d5bbbf283b21af4813232069f7",
					Balance: 50,
				},
				{
					Address: "0xbfAdCF85554F139F978DE5442aacFBe085c754f7",
					Balance: 50,
				},
				{
					Address: "0xF358eb1020375800746ccd5c6638DA36C5a6bec9",
					Balance: 50,
				},
				{
					Address: "0xd6e579085c82329c89fca7a9f012be59028ed53f",
					Balance: 1000000,
				},
			},
		},
	},
}
