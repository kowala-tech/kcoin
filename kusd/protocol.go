package kusd

// Constants to match up protocol versions and messages
const (
	kusd1 = 1
)

// Official short name of the protocol used during capability negotiation.
var ProtocolName = "kusd"

// Supported versions of the kusd protocol (first is primary).
var ProtocolVersions = []uint{kusd1}

// @TODO(rgeraldes)
// Number of implemented message corresponding to different protocol versions.
//var ProtocolLengths = []uint64{17, 8}

const ProtocolMaxMsgSize = 10 * 1024 * 1024 // Maximum cap on the size of a protocol message

// kusd protocol message codes
const (
	// Protocol messages belonging to kusd/1
	StatusMsg          = 0x00
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
)
