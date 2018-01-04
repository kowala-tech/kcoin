package kusd

import (
	"github.com/kowala-tech/kUSD/metrics"
	"github.com/kowala-tech/kUSD/p2p"
)

var (
	propTxnInPacketsMeter     = metrics.NewMeter("kusd/prop/txns/in/packets")
	propTxnInTrafficMeter     = metrics.NewMeter("kusd/prop/txns/in/traffic")
	propTxnOutPacketsMeter    = metrics.NewMeter("kusd/prop/txns/out/packets")
	propTxnOutTrafficMeter    = metrics.NewMeter("kusd/prop/txns/out/traffic")
	propHashInPacketsMeter    = metrics.NewMeter("kusd/prop/hashes/in/packets")
	propHashInTrafficMeter    = metrics.NewMeter("kusd/prop/hashes/in/traffic")
	propHashOutPacketsMeter   = metrics.NewMeter("kusd/prop/hashes/out/packets")
	propHashOutTrafficMeter   = metrics.NewMeter("kusd/prop/hashes/out/traffic")
	propBlockInPacketsMeter   = metrics.NewMeter("kusd/prop/blocks/in/packets")
	propBlockInTrafficMeter   = metrics.NewMeter("kusd/prop/blocks/in/traffic")
	propBlockOutPacketsMeter  = metrics.NewMeter("kusd/prop/blocks/out/packets")
	propBlockOutTrafficMeter  = metrics.NewMeter("kusd/prop/blocks/out/traffic")
	reqHeaderInPacketsMeter   = metrics.NewMeter("kusd/req/headers/in/packets")
	reqHeaderInTrafficMeter   = metrics.NewMeter("kusd/req/headers/in/traffic")
	reqHeaderOutPacketsMeter  = metrics.NewMeter("kusd/req/headers/out/packets")
	reqHeaderOutTrafficMeter  = metrics.NewMeter("kusd/req/headers/out/traffic")
	reqBodyInPacketsMeter     = metrics.NewMeter("kusd/req/bodies/in/packets")
	reqBodyInTrafficMeter     = metrics.NewMeter("kusd/req/bodies/in/traffic")
	reqBodyOutPacketsMeter    = metrics.NewMeter("kusd/req/bodies/out/packets")
	reqBodyOutTrafficMeter    = metrics.NewMeter("kusd/req/bodies/out/traffic")
	reqStateInPacketsMeter    = metrics.NewMeter("kusd/req/states/in/packets")
	reqStateInTrafficMeter    = metrics.NewMeter("kusd/req/states/in/traffic")
	reqStateOutPacketsMeter   = metrics.NewMeter("kusd/req/states/out/packets")
	reqStateOutTrafficMeter   = metrics.NewMeter("kusd/req/states/out/traffic")
	reqReceiptInPacketsMeter  = metrics.NewMeter("kusd/req/receipts/in/packets")
	reqReceiptInTrafficMeter  = metrics.NewMeter("kusd/req/receipts/in/traffic")
	reqReceiptOutPacketsMeter = metrics.NewMeter("kusd/req/receipts/out/packets")
	reqReceiptOutTrafficMeter = metrics.NewMeter("kusd/req/receipts/out/traffic")
	miscInPacketsMeter        = metrics.NewMeter("kusd/misc/in/packets")
	miscInTrafficMeter        = metrics.NewMeter("kusd/misc/in/traffic")
	miscOutPacketsMeter       = metrics.NewMeter("kusd/misc/out/packets")
	miscOutTrafficMeter       = metrics.NewMeter("kusd/misc/out/traffic")

	// consensus
	stateInPacketsMeter        = metrics.NewMeter("kusd/consensus/state/in/packets")
	stateInTrafficMeter        = metrics.NewMeter("kusd/consensus/state/in/traffic")
	stateOutPacketsMeter       = metrics.NewMeter("kusd/consensus/state/out/packets")
	stateOutTrafficMeter       = metrics.NewMeter("kusd/consensus/state/out/traffic")
	proposalInPacketsMeter     = metrics.NewMeter("kusd/consensus/proposals/in/packets")
	proposalInTrafficMeter     = metrics.NewMeter("kusd/consensus/proposals/in/traffic")
	proposalOutPacketsMeter    = metrics.NewMeter("kusd/consensus/proposals/out/packets")
	proposalOutTrafficMeter    = metrics.NewMeter("kusd/consensus/proposals/out/traffic")
	polProposalInPacketsMeter  = metrics.NewMeter("kusd/consensus/polproposals/in/packets")
	polProposalInTrafficMeter  = metrics.NewMeter("kusd/consensus/polproposals/in/traffic")
	polProposalOutPacketsMeter = metrics.NewMeter("kusd/consensus/polproposals/out/packets")
	polProposalOutTrafficMeter = metrics.NewMeter("kusd/consensus/polproposals/out/traffic")
	voteInPacketsMeter         = metrics.NewMeter("kusd/consensus/votes/in/packets")
	voteInTrafficMeter         = metrics.NewMeter("kusd/consensus/votes/in/traffic")
	voteOutPacketsMeter        = metrics.NewMeter("kusd/consensus/votes/out/packets")
	voteOutTrafficMeter        = metrics.NewMeter("kusd/consensus/votes/out/traffic")
	electionInPacketsMeter     = metrics.NewMeter("kusd/consensus/elections/in/packets")
	electionInTrafficMeter     = metrics.NewMeter("kusd/consensus/elections/in/traffic")
	electionOutPacketsMeter    = metrics.NewMeter("kusd/consensus/elections/out/packets")
	electionOutTrafficMeter    = metrics.NewMeter("kusd/consensus/elections/out/traffic")
	fragmentInPacketsMeter     = metrics.NewMeter("kusd/consensus/fragments/in/packets")
	fragmentInTrafficMeter     = metrics.NewMeter("kusd/consensus/fragments/in/traffic")
	fragmentOutPacketsMeter    = metrics.NewMeter("kusd/consensus/fragments/out/packets")
	fragmentOutTrafficMeter    = metrics.NewMeter("kusd/consensus/fragments/out/traffic")
)

// meteredMsgReadWriter is a wrapper around a p2p.MsgReadWriter, capable of
// accumulating the above defined metrics based on the data stream contents.
type meteredMsgReadWriter struct {
	p2p.MsgReadWriter     // Wrapped message stream to meter
	version           int // Protocol version to select correct meters
}

// newMeteredMsgWriter wraps a p2p MsgReadWriter with metering support. If the
// metrics system is disabled, this function returns the original object.
func newMeteredMsgWriter(rw p2p.MsgReadWriter) p2p.MsgReadWriter {
	if !metrics.Enabled {
		return rw
	}
	return &meteredMsgReadWriter{MsgReadWriter: rw}
}

// Init sets the protocol version used by the stream to know which meters to
// increment in case of overlapping message ids between protocol versions.
func (rw *meteredMsgReadWriter) Init(version int) {
	rw.version = version
}

func (rw *meteredMsgReadWriter) ReadMsg() (p2p.Msg, error) {
	// Read the message and short circuit in case of an error
	msg, err := rw.MsgReadWriter.ReadMsg()
	if err != nil {
		return msg, err
	}
	// Account for the data traffic
	packets, traffic := miscInPacketsMeter, miscInTrafficMeter
	switch {
	// @NOTE(rgeraldes) - order is important (most common ones at the top)
	case msg.Code == TxMsg:
		packets, traffic = propTxnInPacketsMeter, propTxnInTrafficMeter

	case msg.Code == NewStateMsg:
		packets, traffic = stateInPacketsMeter, stateInTrafficMeter
	case msg.Code == BlockFragmentMsg:
		packets, traffic = fragmentInPacketsMeter, fragmentInTrafficMeter
	case msg.Code == VoteMsg:
		packets, traffic = voteInPacketsMeter, voteInTrafficMeter
	case msg.Code == ElectionMsg:
		packets, traffic = electionInPacketsMeter, electionInTrafficMeter
	case msg.Code == ProposalMsg:
		packets, trafic = proposalInPacketsMeter, proposalInTrafficMeter
	case msg.Code == ProposalPOLMsg:
		packets, traffic = polProposalInPacketsMeter, polProposalInTrafficMeter

	case msg.Code == BlockHeadersMsg:
		packets, traffic = reqHeaderInPacketsMeter, reqHeaderInTrafficMeter
	case msg.Code == BlockBodiesMsg:
		packets, traffic = reqBodyInPacketsMeter, reqBodyInTrafficMeter

	case msg.Code == NodeDataMsg:
		packets, traffic = reqStateInPacketsMeter, reqStateInTrafficMeter
	case msg.Code == ReceiptsMsg:
		packets, traffic = reqReceiptInPacketsMeter, reqReceiptInTrafficMeter

	case msg.Code == NewBlockHashesMsg:
		packets, traffic = propHashInPacketsMeter, propHashInTrafficMeter
	case msg.Code == NewBlockMsg:
		packets, traffic = propBlockInPacketsMeter, propBlockInTrafficMeter
	}

	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	return msg, err
}

func (rw *meteredMsgReadWriter) WriteMsg(msg p2p.Msg) error {
	// Account for the data traffic
	packets, traffic := miscOutPacketsMeter, miscOutTrafficMeter
	switch {
	case msg.Code == TxMsg:
		packets, traffic = propTxnOutPacketsMeter, propTxnOutTrafficMeter

	case msg.Code == NewStateMsg:
		packets, traffic = stateOutPacketsMeter, stateOutTrafficMeter
	case msg.Code == BlockFragmentMsg:
		packets, traffic = fragmentnOutPacketsMeter, fragmentOutTrafficMeter
	case msg.Code == VoteMsg:
		packets, traffic = voteOutPacketsMeter, voteOutTrafficMeter
	case msg.Code == ElectionMsg:
		packets, traffic = electionOutPacketsMeter, electionOutTrafficMeter
	case msg.Code == ProposalMsg:
		packets, trafic = proposalOutPacketsMeter, proposalOutTrafficMeter
	case msg.Code == ProposalPOLMsg:
		packets, traffic = polProposalOutPacketsMeter, polProposalOutTrafficMeter

	case msg.Code == BlockHeadersMsg:
		packets, traffic = reqHeaderOutPacketsMeter, reqHeaderOutTrafficMeter
	case msg.Code == BlockBodiesMsg:
		packets, traffic = reqBodyOutPacketsMeter, reqBodyOutTrafficMeter

	case msg.Code == NodeDataMsg:
		packets, traffic = reqStateOutPacketsMeter, reqStateOutTrafficMeter
	case msg.Code == ReceiptsMsg:
		packets, traffic = reqReceiptOutPacketsMeter, reqReceiptOutTrafficMeter

	case msg.Code == NewBlockHashesMsg:
		packets, traffic = propHashOutPacketsMeter, propHashOutTrafficMeter
	case msg.Code == NewBlockMsg:
		packets, traffic = propBlockOutPacketsMeter, propBlockOutTrafficMeter
	}

	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	// Send the packet to the p2p layer
	return rw.MsgReadWriter.WriteMsg(msg)
}
