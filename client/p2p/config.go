package p2p

import (
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/libp2p/go-libp2p-crypto"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// Config holds the p2p host options.
type Config struct {
	// PrivateKey private key to identify itself
	PrivateKey *crypto.PrivKey

	// Logger represents a custom logger
	Logger log.Logger

	BootstrapNodes []pstore.PeerInfo

	IsBootstrapNode bool

	ListenAddr string
}
