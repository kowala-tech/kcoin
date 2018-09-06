package blockchain

import "math/big"

//Transaction represents a transaction inside the domain of the wallet backend.
type Transaction struct {
	Hash        string   `json:"hash"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Amount      *big.Int `json:"amount"`
	Timestamp   *big.Int `json:"timestamp"`
	BlockHeight *big.Int `json:"block_height"`
	GasUsed     *big.Int `json:"gas_used"`
}
