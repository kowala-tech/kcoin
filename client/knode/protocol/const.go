package protocol

import (
	"strings"
	"strconv"
)

// Constants to match up protocol versions and messages
const (
	Kcoin1 = 1

	// Official short name of the protocol used during capability negotiation.
	ProtocolName = "kcoin"
)

var (
	ProtocolNameUpper = strings.ToUpper("kcoin")
	ProtocolVersionStr = strconv.Itoa(Kcoin1)
	ProtocolPrefix = ProtocolNameUpper+ProtocolVersionStr
	ProtocolPrefixBytes = []byte(ProtocolNameUpper+ProtocolVersionStr)
)

// Supported versions of the kcoin protocol (first is primary).
var ProtocolVersions = []uint{Kcoin1}

// Number of implemented message corresponding to different protocol versions.
var ProtocolLengths = []uint64{21}

const ProtocolMaxMsgSize = 10 * 1024 * 1024 // Maximum cap on the size of a protocol message
