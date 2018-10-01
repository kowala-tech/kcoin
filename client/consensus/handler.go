package consensus

import (
	"github.com/kowala-tech/kcoin/client/event"
	"github.com/kowala-tech/kcoin/client/consensus/validator"
)

const (
	// voteChSize is the size of channel listening to NewVoteEvent.
	voteChSize = 4096

	// proposalChSize is the size of channel listening to NewVoteEvent.
	proposalChSize = 4096
)

type ProtocolManager struct {
	SubProtocols []p2p.Protocol

	networkID uint64
	
	votingSystem votingSystem
	validator  validator.Validator
	
	maxPeers int
	peers      *peerSet
	newPeerCh chan *peer
	
	voteCh chan core.NewVoteEvent
	voteSub event.Subscription
	proposalCh chan core.NewProposalEvent
	proposalSub event.Subscription
}

func NewProtocolManager(networkID uint64, validator validator.Validator, votingSystem votingSystem) (*ProtocolManager, error) {
	manager := &ProtocolManager{
		networkID: networkID,
		validator:    validator,
		votingSystem: votingSystem,
		peers: newPeerSet(),
		newPeerCh: make(chan *peer),
	}
}

func (pm *ProtocolManager) Start(maxPeers int) {
	pm.maxPeers = maxPeers

	// broadcast proposals
	pm.proposalSub = pm.eventMux.Subscribe(core.NewProposalEvent{}, core.NewBlockFragmentEvent{})
	go pm.proposalBroadcastLoop()

	// broadcast votes
	pm.voteSub = pm.voteSystem.SubscribeNewVoteEvent(core.NewVoteEvent{})
	go pm.voteBroadcastLoop()
}

func (pm *ProtocolManager) Stop() {
	log.Info("Stopping mining protocol")

	pm.proposalSub.Unsubscribe()   // quits proposalBroadcastLoop
	pm.voteSub.Unsubscribe()       // quits voteBroadcastLoop

	// Quit the sync loop.
	// After this send has completed, no new peers will be accepted.
	pm.noMorePeers <- struct{}{}

	// Disconnect existing sessions.
	// This also closes the gate for any new registrations on the peer set.
	// sessions which are already established but not added to pm.peers yet
	// will exit when they try to register.
	pm.peers.Close()

	// Wait for all peer handler goroutines and the loops to come down.
	pm.wg.Wait()

	log.Info("Mining protocol stopped")
}


func (pm *ProtocolManager) handle(p *peer) error {
	// Ignore maxPeers if this is a trusted peer
	if pm.peers.Len() >= pm.maxPeers {
		return p2p.DiscTooManyPeers
	}

	p.Log().Debug("Consensus peer connected", "name", p.Name())
}


func (pm *ProtocolManager) handleMsg(p *peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > protocol.Constants.MaxMsgSize {
		return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, protocol.Constants.MaxMsgSize)
	}
	defer msg.Discard()

	// Handle the message depending on its contents
	switch {
	case msg.Code == ProposalMsg:
		// Retrieve and decode the propagated proposal
		var proposal types.Proposal
		if err := msg.Decode(&proposal); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}

		p.MarkProposal(proposal.Hash())

		if !pm.validator.Validating() {
			// propagate proposal
			go pm.eventMux.Post(core.NewProposalEvent{Proposal: proposal})
			break
		}

		if err := pm.validator.AddProposal(&proposal); err != nil {
			// ignore
			break
		}

	case msg.Code == VoteMsg:
		// Retrieve and decode the propagated vote
		var vote types.Vote
		if err := msg.Decode(&vote); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}

		p.MarkVote(vote.Hash())

		if !pm.validator.Validating() {
			// propagate vote
			go pm.eventMux.Post(core.NewVoteEvent{Vote: vote})
			break
		}

		if err := pm.validator.AddVote(&vote); err != nil {
			// ignore
			break
		}

	case msg.Code == BlockFragmentMsg:
		// Retrieve and decode the propagated block fragment
		var request blockFragmentData
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}

		p.MarkFragment(request.Data.Proof)

		if !pm.validator.Validating() {
			// propagate block fragment
			pm.eventMux.Post(core.NewBlockFragmentEvent{BlockNumber: request.BlockNumber, Round: request.Round, Data: request.Data})
			break
		}

		if err := pm.validator.AddBlockFragment(request.BlockNumber, request.Round, request.Data); err != nil {
			log.Error("error while adding a new block fragment", "err", err, "round", request.Round, "block", request.BlockNumber, "fragment", request.Data)
			// ignore
			break
		}
	}
}

// Proposal broadcast loop
func (pm *ProtocolManager) proposalBroadcastLoop() {
	for obj := range pm.proposalSub.Chan() {
		switch ev := obj.Data.(type) {
		case core.NewProposalEvent:
			for _, peer := range pm.peers.Peers() {
				peer.SendNewProposal(ev.Proposal)
			}
		case core.NewBlockFragmentEvent:
			for _, peer := range pm.peers.PeersWithoutFragment(ev.Data.Proof) {
				peer.SendBlockFragment(ev.BlockNumber, ev.Round, ev.Data)
			}
		}
	}
}

// Vote broadcast loop
func (pm *ProtocolManager) voteBroadcastLoop() {
	for obj := range pm.voteSub.Chan() {
		switch ev := obj.Data.(type) {
		case core.NewVoteEvent:
			peers := pm.peers.PeersWithoutVote(ev.Vote.Hash())
			for _, peer := range peers {
				peer.SendVote(ev.Vote)
			}
		}
	}
}