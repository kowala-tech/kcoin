package mining

import (
	"github.com/kowala-tech/kcoin/client/metrics"
	"github.com/kowala-tech/kcoin/client/p2p"
)

var (
	propProposalInPacketsMeter  = metrics.NewRegisteredMeter("mining/prop/proposal/in/packets", nil)
	propProposalInTrafficMeter  = metrics.NewRegisteredMeter("mining/prop/proposal/in/traffic", nil)
	propProposalOutPacketsMeter = metrics.NewRegisteredMeter("mining/prop/proposal/out/packets", nil)
	propProposalOutTrafficMeter = metrics.NewRegisteredMeter("mining/prop/proposal/out/traffic", nil)

	propBlockFragmentInPacketsMeter  = metrics.NewRegisteredMeter("mining/prop/fragment/in/packets", nil)
	propBlockFragmentInTrafficMeter  = metrics.NewRegisteredMeter("mining/prop/fragment/in/traffic", nil)
	propBlockFragmentOutPacketsMeter = metrics.NewRegisteredMeter("mining/prop/fragment/out/packets", nil)
	propBlockFragmentOutTrafficMeter = metrics.NewRegisteredMeter("mining/prop/fragment/out/traffic", nil)

	propVoteInPacketsMeter  = metrics.NewRegisteredMeter("mining/prop/votes/in/packets", nil)
	propVoteInTrafficMeter  = metrics.NewRegisteredMeter("mining/prop/votes/in/traffic", nil)
	propVoteOutPacketsMeter = metrics.NewRegisteredMeter("mining/prop/votes/out/packets", nil)
	propVoteOutTrafficMeter = metrics.NewRegisteredMeter("mining/prop/votes/out/traffic", nil)

	miscInPacketsMeter  = metrics.NewRegisteredMeter("eth/misc/in/packets", nil)
	miscInTrafficMeter  = metrics.NewRegisteredMeter("eth/misc/in/traffic", nil)
	miscOutPacketsMeter = metrics.NewRegisteredMeter("eth/misc/out/packets", nil)
	miscOutTrafficMeter = metrics.NewRegisteredMeter("eth/misc/out/traffic", nil)
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
	case msg.Code == VoteMsg:
		packets, traffic = propVoteInPacketsMeter, propVoteInTrafficMeter
	case msg.Code == BlockFragmentMsg:
		packets, traffic = propBlockFragmentInPacketsMeter, propBlockFragmentInTrafficMeter
	case msg.Code == ProposalMsg:
		packets, traffic = propProposalInPacketsMeter, propProposalInTrafficMeter
	}
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	return msg, err
}

func (rw *meteredMsgReadWriter) WriteMsg(msg p2p.Msg) error {
	// Account for the data traffic
	packets, traffic := miscOutPacketsMeter, miscOutTrafficMeter
	switch {
	case msg.Code == VoteMsg:
		packets, traffic = propVoteOutPacketsMeter, propVoteOutTrafficMeter
	case msg.Code == BlockFragmentMsg:
		packets, traffic = propBlockFragmentOutPacketsMeter, propBlockFragmentOutTrafficMeter
	case msg.Code == ProposalMsg:
		packets, traffic = propProposalOutPacketsMeter, propProposalOutTrafficMeter
	}
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	// Send the packet to the p2p layer
	return rw.MsgReadWriter.WriteMsg(msg)
}
