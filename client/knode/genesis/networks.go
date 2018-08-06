package genesis

import "github.com/kowala-tech/kcoin/client/knode"

var Networks = map[string]map[string]Options{
	knode.KUSD: {
		MainNetwork: Options{
			Network:     MainNetwork,
			BlockNumber: 0,
			ExtraData:   "Kowala's first block",
			SystemVars: &SystemVarsOpts{
				InitialPrice: 1,
			},
			Governance: &GovernanceOpts{
				Origin: "0x259be75d96876f2ada3d202722523e9cd4dd917d",
				Governors: []string{
					"0xa1e8587ed7f915d5bbbf283b21af4813232069f7",
					"0xbfAdCF85554F139F978DE5442aacFBe085c754f7",
					"0xF358eb1020375800746ccd5c6638DA36C5a6bec9",
				},
				NumConfirmations: 2,
			},
			Consensus: &ConsensusOpts{
				Engine:           KonsensusConsensus,
				MaxNumValidators: 500,
				FreezePeriod:     1,
				BaseDeposit:      1000000,
				SuperNodeAmount:  6000000,
				Validators: []Validator{
					{
						Address: "0xd6e579085c82329c89fca7a9f012be59028ed53f",
						Deposit: 1000000,
					},
				},
				MiningToken: &MiningTokenOpts{
					Name:     "mUSD",
					Symbol:   "mUSD",
					Cap:      1073741824,
					Decimals: 18,
					Holders: []TokenHolder{
						{
							Address:   "0xd6e579085c82329c89fca7a9f012be59028ed53f",
							NumTokens: 3000000,
						},
					},
				},
			},
			StabilityContract: &StabilityContractOpts{
				MinDeposit: 50,
			},
			DataFeedSystem: &DataFeedSystemOpts{
				MaxNumOracles: 50,
				Price: PriceOpts{
					SyncFrequency: 600,
					UpdatePeriod:  30,
				},
			},
			PrefundedAccounts: []PrefundedAccount{
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
		TestNetwork: Options{
			Network:     TestNetwork,
			BlockNumber: 0,
			ExtraData:   "Kowala's first block",
			SystemVars: &SystemVarsOpts{
				InitialPrice: 1,
			},
			Governance: &GovernanceOpts{
				Origin: "0x259be75d96876f2ada3d202722523e9cd4dd917d",
				Governors: []string{
					"0xa1e8587ed7f915d5bbbf283b21af4813232069f7",
					"0xbfAdCF85554F139F978DE5442aacFBe085c754f7",
					"0xF358eb1020375800746ccd5c6638DA36C5a6bec9",
				},
				NumConfirmations: 2,
			},
			Consensus: &ConsensusOpts{
				Engine:           KonsensusConsensus,
				MaxNumValidators: 500,
				FreezePeriod:     1,
				BaseDeposit:      1000000,
				SuperNodeAmount:  6000000,
				Validators: []Validator{
					{
						Address: "0xd6e579085c82329c89fca7a9f012be59028ed53f",
						Deposit: 6000000,
					},
				},
				MiningToken: &MiningTokenOpts{
					Name:     "mUSD",
					Symbol:   "mUSD",
					Cap:      1073741824,
					Decimals: 18,
					Holders: []TokenHolder{
						{
							Address:   "0xd6e579085c82329c89fca7a9f012be59028ed53f",
							NumTokens: 10000000,
						},
					},
				},
			},
			StabilityContract: &StabilityContractOpts{
				MinDeposit: 50,
			},
			DataFeedSystem: &DataFeedSystemOpts{
				MaxNumOracles: 50,
				Price: PriceOpts{
					SyncFrequency: 600,
					UpdatePeriod:  30,
				},
			},
			PrefundedAccounts: []PrefundedAccount{
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
