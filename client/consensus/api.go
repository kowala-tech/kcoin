package consensus

// PrivateValidatorAPI provides private RPC methods to control the validator.
// These methods can be abused by external users and must be considered insecure for use by untrusted users.
type PrivateValidatorAPI struct {
	service *MiningService
}

// NewPrivateValidatorAPI create a new RPC service which controls the validator of this node.
func NewPrivateValidatorAPI(service *MiningService) *PrivateValidatorAPI {
	return &PrivateValidatorAPI{service: service}
}

// Start the validator.
func (api *PrivateValidatorAPI) Start(deposit *hexutil.Big) error {
	// Start the validator and return
	if !api.service.IsValidating() {

		// @TODO (rgeraldes)
		/*
		// Propagate the initial price point to the transaction pool
		api.service.lock.RLock()
		price := api.service.gasPrice
		api.service.lock.RUnlock()
		*/

		bigint := deposit.ToInt()
		if bigint.Cmp(big.NewInt(0)) != 0 {
			err := api.service.SetDeposit(bigint)
			if err != nil && err != validator.ErrIsNotRunning {
				return err
			}
		}

		// @TODO (rgeraldes)
		//api.service.txPool.SetGasPrice(price)
		return api.service.StartValidating()
	}
	return nil
}

// Stop the validator
func (api *PrivateValidatorAPI) Stop() bool {
	api.service.StopValidating()
	return true
}

// SetExtra sets the extra data string that is included when this validator proposes a block.
func (api *PrivateValidatorAPI) SetExtra(extra string) (bool, error) {
	if err := api.service.Validator().SetExtra([]byte(extra)); err != nil {
		return false, err
	}
	return true, nil
}

// SetGasPrice sets the minimum accepted gas price for the validator.
func (api *PrivateValidatorAPI) SetGasPrice(gasPrice hexutil.Big) bool {
	api.service.lock.Lock()
	api.service.gasPrice = (*big.Int)(&gasPrice)
	api.service.lock.Unlock()

	api.service.txPool.SetGasPrice((*big.Int)(&gasPrice))
	return true
}

// SetCoinbase sets the coinbase of the validator
func (api *PrivateValidatorAPI) SetCoinbase(coinbase common.Address) bool {
	api.service.SetCoinbase(coinbase)
	return true
}

// GetMinimumDeposit gets the minimum deposit required to take a slot as a validator
func (api *PrivateValidatorAPI) GetMinimumDeposit() (*big.Int, error) {
	return api.service.GetMinimumDeposit()
}

// GetDepositsResult is the result of a validator_getDeposits API call.
type GetDepositsResult struct {
	Deposits []depositEntry `json:"deposits"`
}

type depositEntry struct {
	Amount      *big.Int `json:"value"`
	AvailableAt string   `json:",omitempty"`
}

// GetDeposits returns the validator deposits
func (api *PrivateValidatorAPI) GetDeposits(address *common.Address) (GetDepositsResult, error) {
	rawDeposits, err := api.service.Validator().Deposits(address)
	if err != nil {
		return GetDepositsResult{}, err
	}

	return depositsToResponse(rawDeposits), nil
}

func depositsToResponse(rawDeposits []*types.Deposit) GetDepositsResult {
	deposits := make([]depositEntry, len(rawDeposits))

	for i, deposit := range rawDeposits {
		// @NOTE (rgeraldes) - zero values are not shown for this field
		var availableAt string

		if deposit.AvailableAtTimeUnix() != 0 {
			availableAt = time.Unix(deposit.AvailableAtTimeUnix(), 0).String()
		}

		deposits[i] = depositEntry{
			Amount:      deposit.Amount(),
			AvailableAt: availableAt,
		}
	}

	return GetDepositsResult{Deposits: deposits}
}

// IsValidating returns the validator is currently validating
func (api *PrivateValidatorAPI) IsValidating() bool {
	return api.service.IsValidating()
}

// IsValidating returns the validator is currently running
func (api *PrivateValidatorAPI) IsRunning() bool {
	return api.service.IsRunning()
}

// RedeemDeposits requests a transfer of the unlocked deposits back
// to the validator account
func (api *PrivateValidatorAPI) RedeemDeposits() error {
	return api.service.Validator().RedeemDeposits()
}