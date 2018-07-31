package protocol

import (
	"strconv"
	"strings"
)

// Constants to match up protocol versions and messages
const (
	Kcoin1 = 1

	// Official short name of the protocol used during capability negotiation.
	ProtocolName = "kcoin"
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
	strconv.Itoa(Kcoin1),
	strings.ToUpper(ProtocolName) + strconv.Itoa(Kcoin1),         // ProtocolNameUpper+ProtocolVersionStr
	[]byte(strings.ToUpper(ProtocolName) + strconv.Itoa(Kcoin1)), // ProtocolNameUpper+ProtocolVersionStr
	[]uint{Kcoin1},
	[]uint64{21},
	10 * 1024 * 1024,
}
