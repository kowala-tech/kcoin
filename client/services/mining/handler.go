package mining

import (
	"errors"
	"fmt"
	"sync"

	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/p2p"
	"github.com/kowala-tech/kcoin/client/p2p/discover"
	"github.com/kowala-tech/kcoin/client/services/mining/validator"
)

const (
	// voteChSize is the size of channel listening to NewVoteEvent.
	voteChSize = 4096
	// proposalChSize is the size of channel listening to NewVoteEvent.
	proposalChSize = 4096
	// blockFragmentChSize is the size of channel listening to NewBlockFragmentEvent.
	blockFragmentChSize = 4096
)

// errIncompatibleConfig is returned if the requested protocols and configs are
// not compatible (low protocol version restrictions and high requirements).
var errIncompatibleConfig = errors.New("incompatible configuration")

type ProtocolManager struct {
	SubProtocols []p2p.Protocol

	networkID uint64

	votingSystem *validator.VotingSystem
	validator    validator.Validator

	maxPeers  int
	peers     *peerSet
	newPeerCh chan *peer

	voteCh           chan core.NewVoteEvent
	voteSub          event.Subscription
	proposalCh       chan core.NewProposalEvent
	proposalSub      event.Subscription
	blockFragmentCh  chan core.NewBlockFragmentEvent
	blockFragmentSub event.Subscription

	chain *core.BlockChain

	quitSync chan struct{}

	// wait group is used for graceful shutdowns
	wg sync.WaitGroup

	log log.Logger
}

func NewProtocolManager(networkID uint64, validator validator.Validator, votingSystem *validator.VotingSystem, chain *core.BlockChain, log log.Logger) (*ProtocolManager, error) {
	manager := &ProtocolManager{
		networkID:    networkID,
		validator:    validator,
		votingSystem: votingSystem,
		peers:        newPeerSet(),
		newPeerCh:    make(chan *peer),
		chain:        chain,
		log:          log,
	}

	// Initiate a sub-protocol for every implemented version we can handle
	manager.SubProtocols = make([]p2p.Protocol, 0, len(Constants.Versions))
	for i, version := range Constants.Versions {
		// Compatible; initialise the sub-protocol
		version := version // Closure for the run
		manager.SubProtocols = append(manager.SubProtocols, p2p.Protocol{
			Name:    ProtocolName,
			Version: version,
			Length:  Constants.Lengths[i],
			Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
				peer := manager.newPeer(int(version), p, rw)
				select {
				case manager.newPeerCh <- peer:
					manager.wg.Add(1)
					defer manager.wg.Done()
					return manager.handle(peer)
				case <-manager.quitSync:
					return p2p.DiscQuitting
				}
			},
			NodeInfo: func() interface{} {
				return map[string]interface{}{
					"version": Constants.VersionStr,
				}
			},
			PeerInfo: func(id discover.NodeID) interface{} {
				if p := manager.peers.Peer(fmt.Sprintf("%x", id[:8])); p != nil {
					return p.Info()
				}
				return nil
			},
		})
	}
	if len(manager.SubProtocols) == 0 {
		return nil, errIncompatibleConfig
	}

	return manager, nil
}

func (pm *ProtocolManager) removePeer(id string) {
	// Short circuit if the peer was already removed
	peer := pm.peers.Peer(id)
	if peer == nil {
		return
	}
	pm.log.Debug("Removing consensus peer", "peer", id)

	if err := pm.peers.Unregister(id); err != nil {
		pm.log.Error("Peer removal failed", "peer", id, "err", err)
	}
	// Hard disconnect at the networking layer
	if peer != nil {
		peer.Peer.Disconnect(p2p.DiscUselessPeer)
	}
}

func (pm *ProtocolManager) Start(maxPeers int) {
	pm.maxPeers = maxPeers

	// broadcast proposals
	pm.proposalCh = make(chan core.NewProposalEvent, proposalChSize)
	pm.proposalSub = pm.validator.SubscribeNewProposalEvent(pm.proposalCh)
	go pm.proposalBroadcastLoop()

	// broadcast block fragments
	pm.blockFragmentCh = make(chan core.NewBlockFragmentEvent, blockFragmentChSize)
	pm.blockFragmentSub = pm.validator.SubscribeNewBlockFragmentEvent(pm.blockFragmentCh)
	go pm.blockFragmentBroadcastLoop()

	// broadcast votes
	pm.voteCh = make(chan core.NewVoteEvent, voteChSize)
	pm.blockFragmentSub = pm.votingSystem.SubscribeNewVoteEvent(pm.voteCh)
}

func (pm *ProtocolManager) Stop() {
	pm.log.Info("Stopping mining protocol")

	pm.proposalSub.Unsubscribe()
	pm.blockFragmentSub.Unsubscribe()
	pm.voteSub.Unsubscribe()

	close(pm.quitSync)

	// Disconnect existing sessions.
	// This also closes the gate for any new registrations on the peer set.
	// sessions which are already established but not added to pm.peers yet
	// will exit when they try to register.
	pm.peers.Close()

	// Wait for all peer handler goroutines and the loops to come down.
	pm.wg.Wait()

	pm.log.Info("Mining protocol stopped")
}

func (pm *ProtocolManager) newPeer(pv int, p *p2p.Peer, rw p2p.MsgReadWriter) *peer {
	return newPeer(pv, p, newMeteredMsgWriter(rw))
}

func (pm *ProtocolManager) handle(p *peer) error {
	// Ignore maxPeers if this is a trusted peer
	if pm.peers.Len() >= pm.maxPeers {
		return p2p.DiscTooManyPeers
	}

	p.Log().Debug("Consensus peer connected", "name", p.Name())

	// Execute the Kowala handshake
	var (
		genesis     = pm.chain.Genesis()
		head        = pm.chain.CurrentHeader()
		hash        = head.Hash()
		blockNumber = head.Number
	)
	if err := p.Handshake(pm.networkID, blockNumber, hash, genesis.Hash()); err != nil {
		p.Log().Debug("Kowala handshake failed", "err", err)
		return err
	}
	if rw, ok := p.rw.(*meteredMsgReadWriter); ok {
		rw.Init(p.version)
	}
	// Register the peer locally
	if err := pm.peers.Register(p); err != nil {
		p.Log().Error("Kowala peer registration failed", "err", err)
		return err
	}
	defer pm.removePeer(p.id)

	// main loop. handle incoming messages.
	for {
		if err := pm.handleMsg(p); err != nil {
			p.Log().Debug("Kowala message handling failed", "err", err)
			return err
		}
	}

	return nil
}

func (pm *ProtocolManager) handleMsg(p *peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > Constants.MaxMsgSize {
		return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, Constants.MaxMsgSize)
	}
	defer msg.Discard()

	// Handle the message depending on its contents
	switch {
	case msg.Code == ProposalMsg:
		if !pm.validator.Validating() {
			break
		}
		// Retrieve and decode the propagated proposal
		var proposal types.Proposal
		if err := msg.Decode(&proposal); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		p.MarkProposal(proposal.Hash())
		if err := pm.validator.AddProposal(&proposal); err != nil {
			// ignore
			break
		}
	case msg.Code == VoteMsg:
		if !pm.validator.Validating() {
			break
		}
		// Retrieve and decode the propagated vote
		var vote types.Vote
		if err := msg.Decode(&vote); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		p.MarkVote(vote.Hash())
		if err := pm.validator.AddVote(&vote); err != nil {
			// ignore
			break
		}
	case msg.Code == BlockFragmentMsg:
		if !pm.validator.Validating() {
			break
		}
		// Retrieve and decode the propagated block fragment
		var request blockFragmentData
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		p.MarkBlockFragment(request.Data.Proof)
		if err := pm.validator.AddBlockFragment(request.BlockNumber, request.Round, request.Data); err != nil {
			log.Error("error while adding a new block fragment", "err", err, "round", request.Round, "block", request.BlockNumber, "fragment", request.Data)
			// ignore
			break
		}

	default:
		return errResp(ErrInvalidMsgCode, "%v", msg.Code)
	}

	return nil
}

func (pm *ProtocolManager) proposalBroadcastLoop() {
	for {
		select {
		case event := <-pm.proposalCh:
			for _, peer := range pm.peers.PeersWithoutProposal(event.Proposal.Hash()) {
				peer.SendProposal(event.Proposal)
			}
		case <-pm.proposalSub.Err():
			return
		}
	}
}

func (pm *ProtocolManager) blockFragmentBroadcastLoop() {
	for {
		select {
		case event := <-pm.blockFragmentCh:
			for _, peer := range pm.peers.PeersWithoutBlockFragment(event.Data.Proof) {
				peer.SendBlockFragment(event.BlockNumber, event.Round, event.Data)
			}
		case <-pm.blockFragmentSub.Err():
			return
		}
	}
}

func (pm *ProtocolManager) voteBroadcastLoop() {
	for {
		select {
		case event := <-pm.voteCh:
			for _, peer := range pm.peers.PeersWithoutVote(event.Vote.Hash()) {
				peer.SendVote(event.Vote)
			}
		case <-pm.voteSub.Err():
			return
		}
	}
}

func errResp(code errCode, format string, v ...interface{}) error {
	return fmt.Errorf("%v - %v", code, fmt.Sprintf(format, v...))
}
