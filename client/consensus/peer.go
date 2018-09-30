package consensus

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/knode/protocol"
	"github.com/kowala-tech/kcoin/client/p2p"
	"gopkg.in/fatih/set.v0"
)

var (
	errClosed            = errors.New("peer set is closed")
	errAlreadyRegistered = errors.New("peer is already registered")
	errNotRegistered     = errors.New("peer is not registered")
)

const (
	maxKnownVotes          = 2048 // Maximum vote hashes to keep in the known list (prevent DOS)
	maxKnownProposals      = 2048 // Maximum proposal hashes to keep in the known list (prevent DOS)
	maxKnownBlockFragments = 2048 // Maximum fragment hashes to keep in the known list (prevent DOS)

	// maxQueuedProps is the maximum number of block propagations to queue up before
	// dropping broadcasts. There's not much point in queueing stale blocks, so a few
	// that might cover uncles should be enough.
	maxQueuedProps = 4

	// maxQueuedAnns is the maximum number of block announcements to queue up before
	// dropping broadcasts. Similarly to block propagations, there's no point to queue
	// above some healthy uncle limit, so use that.
	maxQueuedAnns = 4

	handshakeTimeout = 5 * time.Second
)

// PeerInfo represents a short summary of the Kowala sub-protocol metadata known
// about a connected peer.
type PeerInfo struct {
	Version     int      `json:"version"` // Ethereum protocol version negotiated
	BlockNumber *big.Int `json:"number"`  // Block number of the peer's blockchain
	Head        string   `json:"head"`    // SHA3 hash of the peer's best owned block
}

// propEvent is a block propagation, waiting for its turn in the broadcast queue.
type propEvent struct {
	block *types.Block
}

type peer struct {
	id string

	*p2p.Peer
	rw p2p.MsgReadWriter

	version int // Constants version negotiated

	blockNumber *big.Int
	head        common.Hash
	lock        sync.RWMutex

	knownProposals      *set.Set // Set of proposal hashes known to be known by this peer
	knownVotes          *set.Set // Set of vote hashes known to be known by this peer
	knownBlockFragments *set.Set // Set of fragment hashes known to be known by this peer

	queuedProps chan *propEvent   // Queue of blocks to broadcast to the peer
	queuedAnns  chan *types.Block // Queue of blocks to announce to the peer
	term        chan struct{}     // Termination channel to stop the broadcaster
}

func newPeer(version int, p *p2p.Peer, rw p2p.MsgReadWriter) *peer {
	return &peer{
		Peer:                p,
		rw:                  rw,
		version:             version,
		id:                  fmt.Sprintf("%x", p.ID().Bytes()[:8]),
		knownProposals:      set.New(),
		knownVotes:          set.New(),
		knownBlockFragments: set.New(),
		queuedProps:         make(chan *propEvent, maxQueuedProps),
		queuedAnns:          make(chan *types.Block, maxQueuedAnns),
		term:                make(chan struct{}),
	}
}

// broadcast is a write loop that multiplexes block propagations, announcements
// and transaction broadcasts into the remote peer. The goal is to have an async
// writer that does not lock up node internals.
func (p *peer) broadcast() {
	for {
		select {
		case txs := <-p.queuedTxs:
			if err := p.SendTransactions(txs); err != nil {
				return
			}
			p.Log().Trace("Broadcast transactions", "count", len(txs))

		case prop := <-p.queuedProps:
			if err := p.SendNewBlock(prop.block); err != nil {
				return
			}
			p.Log().Trace("Propagated block", "number", prop.block.Number(), "hash", prop.block.Hash())

		case block := <-p.queuedAnns:
			if err := p.SendNewBlockHashes([]common.Hash{block.Hash()}, []uint64{block.NumberU64()}); err != nil {
				return
			}
			p.Log().Trace("Announced block", "number", block.Number(), "hash", block.Hash())

		case <-p.term:
			return
		}
	}
}

// close signals the broadcast goroutine to terminate.
func (p *peer) close() {
	close(p.term)
}

// Info gathers and returns a collection of metadata known about a peer.
func (p *peer) Info() *PeerInfo {
	hash, blockNumber := p.Head()

	return &PeerInfo{
		Version:     p.version,
		BlockNumber: blockNumber,
		Head:        hash.Hex(),
	}
}

// Head retrieves a copy of the current head hash and block number of the
// peer.
func (p *peer) Head() (hash common.Hash, blockNumber *big.Int) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	copy(hash[:], p.head[:])
	return hash, new(big.Int).Set(p.blockNumber)
}

// SetHead updates the head hash and block number of the peer.
func (p *peer) SetHead(hash common.Hash, blockNumber *big.Int) {
	p.lock.Lock()
	defer p.lock.Unlock()

	copy(p.head[:], hash[:])
	p.blockNumber.Set(blockNumber)
}

// MarkVote marks a vote as known for the peer, ensuring that the block will
// never be propagated to this particular peer.
func (p *peer) MarkVote(hash common.Hash) {
	// If we reached the memory allowance, drop a previously known vote hash
	for p.knownVotes.Size() >= maxKnownVotes {
		p.knownVotes.Pop()
	}
	p.knownVotes.Add(hash)
}

// MarkBlockFragment marks a block fragment as known for the peer, ensuring that the
// fragment will never be propagated to this particular peer.
func (p *peer) MarkBlockFragment(hash common.Hash) {
	// If we reached the memory allowance, drop a previously known fragment hash
	for p.knownBlockFragments.Size() >= maxKnownBlockFragments {
		p.knownBlockFragments.Pop()
	}
	p.knownBlockFragments.Add(hash)
}

// SendNewBlock propagates a proposal to a remote peer.
func (p *peer) SendNewProposal(proposal *types.Proposal) error {
	p.knownProposals.Add(proposal.Hash())
	return p2p.Send(p.rw, ProposalMsg, proposal)
}

// AsyncSendNewBlock queues an entire block for propagation to a remote peer. If
// the peer's broadcast queue is full, the event is silently dropped.
func (p *peer) AsyncSendNewBlock(block *types.Block) {
	select {
	case p.queuedProps <- &propEvent{block: block}:
		p.knownBlocks.Add(block.Hash())
	default:
		p.Log().Debug("Dropping block propagation", "number", block.NumberU64(), "hash", block.Hash())
	}
}

// SendNewBlock propagates a vote to a remote peer.
func (p *peer) SendVote(vote *types.Vote) error {
	p.knownVotes.Add(vote.Hash())
	return p2p.Send(p.rw, VoteMsg, vote)
}

// SendBlockFragment propagates a block fragment to a remote peer.
func (p *peer) SendBlockFragment(blockNumber *big.Int, round uint64, data *types.BlockFragment) error {
	p.knownBlockFragments.Add(data.Proof)
	return p2p.Send(p.rw, BlockFragmentMsg, blockFragmentData{blockNumber, round, data})
}

// Handshake executes the eth protocol handshake, negotiating version number,
// network IDs, head and genesis blocks.
func (p *peer) Handshake(network uint64, blockNumber *big.Int, head common.Hash, genesis common.Hash) error {
	// Send out own handshake in a new thread
	errc := make(chan error, 2)
	var status statusData // safe to read after two values have been received from errc

	go func() {
		errc <- p2p.Send(p.rw, StatusMsg, &statusData{
			ProtocolVersion: uint32(p.version),
			NetworkId:       network,
			BlockNumber:     blockNumber,
			CurrentBlock:    head,
			GenesisBlock:    genesis,
		})
	}()
	go func() {
		errc <- p.readStatus(network, &status, genesis)
	}()
	timeout := time.NewTimer(handshakeTimeout)
	defer timeout.Stop()
	for i := 0; i < 2; i++ {
		select {
		case err := <-errc:
			if err != nil {
				return err
			}
		case <-timeout.C:
			return p2p.DiscReadTimeout
		}
	}
	p.blockNumber, p.head = status.BlockNumber, status.CurrentBlock
	return nil
}

func (p *peer) readStatus(network uint64, status *statusData, genesis common.Hash) (err error) {
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Code != StatusMsg {
		return errResp(ErrNoStatusMsg, "first msg has code %x (!= %x)", msg.Code, StatusMsg)
	}
	if msg.Size > protocol.Constants.MaxMsgSize {
		return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, protocol.Constants.MaxMsgSize)
	}
	// Decode the handshake and make sure everything matches
	if err := msg.Decode(&status); err != nil {
		return errResp(ErrDecode, "msg %v: %v", msg, err)
	}
	if status.GenesisBlock != genesis {
		return errResp(ErrGenesisBlockMismatch, "%x (!= %x)", status.GenesisBlock[:8], genesis[:8])
	}
	if status.NetworkId != network {
		return errResp(ErrNetworkIdMismatch, "%d (!= %d)", status.NetworkId, network)
	}
	if int(status.ProtocolVersion) != p.version {
		return errResp(ErrProtocolVersionMismatch, "%d (!= %d)", status.ProtocolVersion, p.version)
	}
	return nil
}

// String implements fmt.Stringer.
func (p *peer) String() string {
	return fmt.Sprintf("Peer %s [%s]", p.id,
		fmt.Sprintf("kcoin/%2d", p.version),
	)
}

// peerSet represents the collection of active peers currently participating in
// the Ethereum sub-protocol.
type peerSet struct {
	peers  map[string]*peer
	lock   sync.RWMutex
	closed bool
}

// newPeerSet creates a new peer set to track the active participants.
func newPeerSet() *peerSet {
	return &peerSet{
		peers: make(map[string]*peer),
	}
}

// Register injects a new peer into the working set, or returns an error if the
// peer is already known. If a new peer it registered, its broadcast loop is also
// started.
func (ps *peerSet) Register(p *peer) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if ps.closed {
		return errClosed
	}
	if _, ok := ps.peers[p.id]; ok {
		return errAlreadyRegistered
	}
	ps.peers[p.id] = p
	go p.broadcast()

	return nil
}

// Unregister removes a remote peer from the active set, disabling any further
// actions to/from that particular entity.
func (ps *peerSet) Unregister(id string) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	p, ok := ps.peers[id]
	if !ok {
		return errNotRegistered
	}
	delete(ps.peers, id)
	p.close()

	return nil
}

// Peer retrieves the registered peer with the given id.
func (ps *peerSet) Peer(id string) *peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	return ps.peers[id]
}

// Len returns if the current number of peers in the set.
func (ps *peerSet) Len() int {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	return len(ps.peers)
}

// PeersWithoutVote retrieves a list of peers that do not have a given vote
// in their set of known hashes.
func (ps *peerSet) PeersWithoutVote(hash common.Hash) []*peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	list := make([]*peer, 0, len(ps.peers))
	for _, p := range ps.peers {
		if !p.knownVotes.Has(hash) {
			list = append(list, p)
		}
	}
	return list
}

// PeersWithoutFragment retrieves a list of peers that do not have a given block fragment
// in their set of known hashes.
func (ps *peerSet) PeersWithoutFragment(hash common.Hash) []*peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	list := make([]*peer, 0, len(ps.peers))
	for _, p := range ps.peers {
		if !p.knownBlockFragments.Has(hash) {
			list = append(list, p)
		}
	}
	return list
}

func (ps *peerSet) Peers() []*peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	list := make([]*peer, 0, len(ps.peers))
	for _, p := range ps.peers {
		list = append(list, p)
	}
	return list
}

// BestPeer retrieves the known peer with the currently highest block number.
func (ps *peerSet) BestPeer() *peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	var (
		bestPeer        *peer
		bestBlockNumber *big.Int
	)
	for _, p := range ps.peers {
		if _, blockNumber := p.Head(); bestPeer == nil || blockNumber.Cmp(bestBlockNumber) > 0 {
			bestPeer, bestBlockNumber = p, blockNumber
		}
	}
	return bestPeer
}

// Close disconnects all peers.
// No new peers can be registered after Close has returned.
func (ps *peerSet) Close() {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	for _, p := range ps.peers {
		p.Disconnect(p2p.DiscQuitting)
	}
	ps.closed = true
}
