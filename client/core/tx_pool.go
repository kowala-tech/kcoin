package core

type TxPool struct {
	byNonce map[common.Address]*txList
}

// AddLocal enqueues a single transaction into the pool if it is valid, making the sender 
// as local if the distinction between local/remote transactions is enabled.
func (pool *TxPool) AddLocal(tx *types.Transaction) error {
	return pool.addTx(tx, !pool.config.NoLocals)
}

// AddRemote enqueues a single remote transaction into the pool if it is valid.
func (pool *TxPool) AddRemote(tx *types.Transaction) error {
	return pool.addTx(tx, false)
}

// AddLocals enqueues a batch of transactions into the pool if they are valid,
// marking the senders as a local ones if the distinction between local/remote 
// transactions is enabled.
func (pool *TxPool) AddLocals(txs []*types.Transaction) []error {
	return pool.addTxs(txs, !pool.config.NoLocals)
}

// AddRemotes enqueues a batch of remote transactions into the pool if they are valid.
func (pool *TxPool) AddRemotes(txs []*types.Transaction) []error {
	return pool.addTxs(txs, false)
}

// addTxs attempts to queue a batch of transactions if they are valid.
func (pool *TxPool) addTxs(txs []*types.Transaction, local bool) []error {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	return pool.addTxsLocked(txs, local)
}

// addTx enqueues a single transaction into the pool if it is valid.
func (pool *TxPool) addTx(tx *types.Transaction, local bool) error {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	// Try to inject the transaction and update any state
	replace, err := pool.add(tx, local)
	if err != nil {
		return err
	}
	// If we added a new transaction, run promotion checks and return
	if !replace {
		from, _ := types.TxSender(pool.signer, tx) // already validated
		pool.promoteExecutables([]common.Address{from})
	}
	return nil
}

// addTxsLocked attempts to queue a batch of transactions if they are valid,
// whilst assuming the transaction pool lock is already held.
func (pool *TxPool) addTxsLocked(txs []*types.Transaction, local bool) []error {
	// Add the batch of transaction, tracking the accepted ones
	dirty := make(map[common.Address]struct{})
	errs := make([]error, len(txs))

	for i, tx := range txs {
		var replace bool
		if replace, errs[i] = pool.add(tx, local); errs[i] == nil && !replace {
			from, _ := types.TxSender(pool.signer, tx) // already validated
			dirty[from] = struct{}{}
		}
	}
	// Only reprocess the internal state if something was actually added
	if len(dirty) > 0 {
		addrs := make([]common.Address, 0, len(dirty))
		for addr := range dirty {
			addrs = append(addrs, addr)
		}
		pool.promoteExecutables(addrs)
	}
	return errs
}

// add validates a transaction and inserts it into the pending queue for
// later promotion and execution.
func (pool *TxPool) add(tx *types.Transaction, local bool) (bool, error) {
	// If the transaction is already known, discard it
	hash := tx.Hash()
	if pool.all.Get(hash) != nil {
		log.Trace("Discarding already known transaction", "hash", hash)
		return false, fmt.Errorf("known transaction: %x", hash)
	}
	// If the transaction fails basic validation, discard it
	if err := pool.validateTx(tx, local); err != nil {
		log.Trace("Discarding invalid transaction", "hash", hash, "err", err)
		invalidTxCounter.Inc(1)
		return false, err
	}

	// If the transaction overlaps an existing one (repeated nonce)
	from, _ := types.TxSender(pool.signer, tx)
	if list := pool.byNonce[from]; list != nil && list.Overlaps(tx) {
		log.Trace("Discarding transaction that overlaps an existing one", "nonce", tx.Nonce(), "err", err)
		invalidTxCounter.Inc(1)
		return false, fmt.Errorf("known transaction: %x", hash)
	}

	replace, err := pool.enqueueTx(hash, tx)
	if err != nil {
		return false, err
	}
	// Mark local addresses and journal local transactions
	if local {
		pool.locals.add(from)
	}
	pool.journalTx(from, tx)

	log.Trace("Pooled new future transaction", "hash", hash, "from", from, "to", tx.To())
	return replace, nil
}

// validateTx checks whether a transaction is valid according to the consensus
// rules and adheres to some heuristic limits (ex: size).
func (pool *TxPool) validateTx(tx *types.Transaction, local bool) error {
	// Heuristic limit, reject transactions over the maximum size to prevent DOS attacks
	if tx.Size() > params.MaxTxSize {
		return ErrOversizedData
	}
	// Transactions can't be negative. This may never happen using RLP decoded
	// transactions but may occur if you create a transaction using the RPC.
	if tx.Value().Sign() < 0 {
		return ErrNegativeValue
	}
	// Ensure the transaction doesn't exceed the block compute capacity.
	if params.ComputeCapacity < tx.ComputeLimit() {
		return ErrComputeCapacity
	}
	// Make sure the transaction is signed properly
	from, err := types.TxSender(pool.signer, tx)
	if err != nil {
		return ErrInvalidSender
	}

	// Ensure the transaction has a valid nonce
	if pool.currentState.GetNonce(from) > tx.Nonce() {
		return ErrNonceTooLow
	}

	// Sender should have enough funds to cover the costs
	// cost == value + computational effort (compute units) * compute unit price
	if pool.currentState.GetBalance(from).Cmp(tx.Cost()) < 0 {
		return ErrInsufficientFunds
	}

	intrCompEffort, err := IntrinsicCompEffort(tx.Data(), tx.To() == nil, true)
	if err != nil {
		return err
	}
	if tx.ComputeLimit() < intrCompEffort {
		return ErrIntrinsicCompEffort
	}
	return nil
}

