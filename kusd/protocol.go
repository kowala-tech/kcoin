package kusd

import (
	"fmt"
	"io"
	"math/big"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/rlp"
)

// Constants to match up protocol versions and messages
const (
	kusd1 = 1
)

// ProtocolName is the official short name of the protocol used during capability negotiation.
var ProtocolName = "kusd"

// ProtocolVersions represents the supported versions of the kusd protocol (first is primary).
var ProtocolVersions = []uint{kusd1}

// Number of implemented message corresponding to different protocol versions.
var ProtocolLengths = []uint64{18}

const ProtocolMaxMsgSize = 10 * 1024 * 1024 // Maximum cap on the size of a protocol message

// kusd1 protocol message codes
const (
	// handshake
	StatusMsg = 0x00

	// sync
	NewBlockHashesMsg  = 0x01
	TxMsg              = 0x02
	GetBlockHeadersMsg = 0x03
	BlockHeadersMsg    = 0x04
	GetBlockBodiesMsg  = 0x05
	BlockBodiesMsg     = 0x06
	NewBlockMsg        = 0x07
	GetNodeDataMsg     = 0x0d
	NodeDataMsg        = 0x0e
	GetReceiptsMsg     = 0x0f
	ReceiptsMsg        = 0x10

	// consensus
	NewStateMsg = 0x11
	ProposalMsg    = 0x12 
	ProposalPOLMsg	= 0x13
	VoteMsg	= 0x14 
	ElectionMsg = 0x15 
	BlockFragmentMsg = 0x17
)

type errCode int

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
	NetworkID       uint64
	CurrentBlock    common.Hash
	GenesisBlock    common.Hash
}

// newBlockHashesData is the network packet for the block announcements.
type newBlockHashesData []struct {
	Hash   common.Hash // Hash of one particular block being announced
	Number uint64      // Number of one particular block being announced
}

// getBlockHeadersData represents a block header query.
type getBlockHeadersData struct {
	Origin  hashOrNumber // Block from which to retrieve headers
	Amount  uint64       // Maximum number of headers to retrieve
	Skip    uint64       // Blocks to skip between consecutive headers
	Reverse bool         // Query direction (false = rising towards latest, true = falling towards genesis)
}

// hashOrNumber is a combined field for specifying an origin block.
type hashOrNumber struct {
	Hash   common.Hash // Block hash from which to retrieve headers (excludes Number)
	Number uint64      // Block hash from which to retrieve headers (excludes Hash)
}

// EncodeRLP is a specialized encoder for hashOrNumber to encode only one of the
// two contained union fields.
func (hn *hashOrNumber) EncodeRLP(w io.Writer) error {
	if hn.Hash == (common.Hash{}) {
		return rlp.Encode(w, hn.Number)
	}
	if hn.Number != 0 {
		return fmt.Errorf("both origin hash (%x) and number (%d) provided", hn.Hash, hn.Number)
	}
	return rlp.Encode(w, hn.Hash)
}

// DecodeRLP is a specialized decoder for hashOrNumber to decode the contents
// into either a block hash or a block number.
func (hn *hashOrNumber) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	origin, err := s.Raw()
	if err == nil {
		switch {
		case size == 32:
			err = rlp.DecodeBytes(origin, &hn.Hash)
		case size <= 8:
			err = rlp.DecodeBytes(origin, &hn.Number)
		default:
			err = fmt.Errorf("invalid input size %d for origin", size)
		}
	}
	return err
}

// newBlockData is the network packet for the block propagation message.
type newBlockData struct {
	Block *types.Block
}

// blockBody represents the data content of a single block.
type blockBody struct {
	Transactions []*types.Transaction // Transactions contained within a block
}

// blockBodiesData is the network packet for block content distribution.
type blockBodiesData []*blockBody

// newStateData is the network packet that is sent to announce a new validator status
type newStateData struct {
	BlockNumber *big.Int
	Round       int
	State        validator.
}

// proposalPOLData is the network packet for the re-proposed proposal message
type proposalPOLData struct {
	BlockNumber *big.Int
	POLRound    int
	POL         *types.BitArray
}

// ElectionData is the network packet that is sent to indicate that a given candidate (block) has seen +2/3 votes
type ElectionData struct {
	Type types.VoteType
	BlockNumber   *big.Int
	BlockHash common.Hash
	Round    int
}

type blockFragmentData struct {
	BlockNumber *big.Int
	Round int
	Data *types.Chunk
}