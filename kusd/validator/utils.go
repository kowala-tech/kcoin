package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	kowala "github.com/kowala-tech/kUSD"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

func getTransactionOpts(contractBackend bind.ContractBackend, from common.Address, value *big.Int, chainID *big.Int) (*bind.TransactOpts, error) {
	// estimate used gas for the transaction
	msg := kowala.CallMsg{From: from /*To: val.network.*/, Data: []byte("teste")}
	if value != nil {
		msg.Value = value
	}

	gasLimit, err := contractBackend.EstimateGas(context.TODO(), msg)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
	}

	// get a gas price suggestion
	gasPrice, err := contractBackend.SuggestGasPrice(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to confirm the gas price suggestion: %v", err)
	}

	opts := &bind.TransactOpts{
		From: from,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			// @NOTE (rgeraldes) - ignore the proposed signer as by default it will be a unprotected signer.
			if address != from {
				return nil, errors.New("not authorized to sign this account")
			}

			return val.wallet.SignTx(from, tx, val.config.ChainID)
		},
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}

	if value != nil {
		opts.Value = value
	}

}
