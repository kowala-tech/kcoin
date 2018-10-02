package mining

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

const (
	StatusMsg        = 0x00
	ProposalMsg      = 0x10
	ProposalPOLMsg   = 0x11
	VoteMsg          = 0x12
	ElectionMsg      = 0x13
	BlockFragmentMsg = 0x14
)

const (
	ErrMsgTooLarge = iota
	ErrDecode
	ErrInvalidMsgCode
	ErrProtocolVersionMismatch
	ErrNetworkIdMismatch
	ErrGenesisBlockMismatch
	ErrNoStatusMsg
	ErrExtraStatusMsg
	ErrSuspendedPeer
)

// Constants to match up protocol versions and messages
const (
	Validator1 = 1

	// Official short name of the protocol used during capability negotiation.
	ProtocolName = "validator"
)

var Constants = struct {
	NameUpper   string
	VersionStr  string
	Prefix      string
	PrefixBytes []byte

	// Supported versions of the kcoin protocol (first is primary).
	Versions []uint

	// Number of implemented message corresponding to different protocol versions.
	Lengths []uint64

	// Maximum cap on the size of a protocol message
	MaxMsgSize uint32
}{
	strings.ToUpper(ProtocolName),
	strconv.Itoa(Validator1),
	strings.ToUpper(ProtocolName) + strconv.Itoa(Validator1),         // ProtocolNameUpper+ProtocolVersionStr
	[]byte(strings.ToUpper(ProtocolName) + strconv.Itoa(Validator1)), // ProtocolNameUpper+ProtocolVersionStr
	[]uint{Validator1},
	[]uint64{5},
	10 * 1024 * 1024,
}

type errCode int

func (e errCode) String() string {
	return errorToString[int(e)]
}

// XXX change once legacy code is out
var errorToString = map[int]string{
	ErrMsgTooLarge:             "Message too long",
	ErrDecode:                  "Invalid message",
	ErrInvalidMsgCode:          "Invalid message code",
	ErrProtocolVersionMismatch: "Protocol version mismatch",
	ErrNetworkIdMismatch:       "NetworkId mismatch",
	ErrGenesisBlockMismatch:    "Genesis block mismatch",
	ErrNoStatusMsg:             "No status message",
	ErrExtraStatusMsg:          "Extra status message",
	ErrSuspendedPeer:           "Suspended peer",
}

// statusData is the network packet for the status message.
type statusData struct {
	ProtocolVersion uint32
	NetworkId       uint64
	BlockNumber     *big.Int
	CurrentBlock    common.Hash
	GenesisBlock    common.Hash
}

// blockFragmentData is the network packet that is sent to let the other validators have a part of the proposed block
type blockFragmentData struct {
	BlockNumber *big.Int
	Round       uint64
	Data        *types.BlockFragment
}
