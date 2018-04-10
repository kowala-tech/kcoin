package kcoinclient

import "github.com/kowala-tech/kcoin"

// Verify that Client implements the ethereum interfaces.
var (
	_ = kowala.ChainReader(&Client{})
	_ = kowala.TransactionReader(&Client{})
	_ = kowala.ChainStateReader(&Client{})
	_ = kowala.ChainSyncReader(&Client{})
	_ = kowala.ContractCaller(&Client{})
	_ = kowala.GasEstimator(&Client{})
	_ = kowala.GasPricer(&Client{})
	_ = kowala.LogFilterer(&Client{})
	_ = kowala.PendingStateReader(&Client{})
	// _ = ethereum.PendingStateEventer(&Client{})
	_ = kowala.PendingContractCaller(&Client{})
)
