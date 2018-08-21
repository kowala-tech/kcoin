package genesis

import "github.com/kowala-tech/kcoin/client/knode"

var Networks = map[string]map[string]Options{
	knode.KUSD: {
		MainNetwork: Options{
			Network:   MainNetwork,
			ExtraData: "Kowala's first block",
			Governance: &GovernanceOpts{
				Origin: "0x259be75d96876f2ada3d202722523e9cd4dd917d",
				Governors: []string{
					"0x6D5E05684c737D42F313d5B82A88090136e831F8",
					"0x049ec8777b4806eff0Bb6039551690D8f650B25a",
					"0x902f069aF381a650B7F18Ff28ffdAd0f11eb425b",
				},
				NumConfirmations: 2,
			},
			Consensus: &ConsensusOpts{
				Engine:           KonsensusConsensus,
				MaxNumValidators: 100,
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
			DataFeedSystem: &DataFeedSystemOpts{
				MaxNumOracles: 1000,
				FreezePeriod:  1,
				BaseDeposit:   10,
				Price: PriceOpts{
					InitialPrice:  1,
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
			Network:   TestNetwork,
			ExtraData: "Kowala's first block",
			Governance: &GovernanceOpts{
				Origin: "0x259be75d96876f2ada3d202722523e9cd4dd917d",
				Governors: []string{
					"0xf861e10641952a42f9c527a43ab77c3030ee2c8f",
					"0x7dd43075b89c129bcd2cca1e2d680a6f3f30b5d9",
					"0xa1d4755112491db5ddf0e10b9253b5a0f6783759",
				},
				NumConfirmations: 2,
			},
			Consensus: &ConsensusOpts{
				Engine:           KonsensusConsensus,
				MaxNumValidators: 100,
				FreezePeriod:     1,
				BaseDeposit:      1000000,
				SuperNodeAmount:  6000000,
				Validators: []Validator{
					{
						Address: "0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0",
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
							Address:   "0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0",
							NumTokens: 10000000,
						},
					},
				},
			},
			DataFeedSystem: &DataFeedSystemOpts{
				MaxNumOracles: 1000,
				FreezePeriod:  1,
				BaseDeposit:   10,
				Price: PriceOpts{
					InitialPrice:  1,
					SyncFrequency: 600,
					UpdatePeriod:  30,
				},
			},
			PrefundedAccounts: []PrefundedAccount{
				{
					Address: "0xf861e10641952a42f9c527a43ab77c3030ee2c8f",
					Balance: 50,
				},
				{
					Address: "0x7dd43075b89c129bcd2cca1e2d680a6f3f30b5d9",
					Balance: 50,
				},
				{
					Address: "0xa1d4755112491db5ddf0e10b9253b5a0f6783759",
					Balance: 50,
				},
				{
					Address: "0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0",
					Balance: 1000000,
				},
				{
					Address: "0x45880e0ab20b1ca0391e8fe871fa035e58edada9",
					Balance: 1000000,
				},
				{
					Address: "0xdac38f0e18ef8bd32aaae695f82e37e14a75a74b",
					Balance: 1000000,
				},
			},
		},
	},
}
