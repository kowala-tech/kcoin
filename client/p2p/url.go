package p2p

import (
	pstore "github.com/libp2p/go-libp2p-peerstore"
	maddr "github.com/multiformats/go-multiaddr"
)

// ParseURL returns the peer info for a given URL
func ParseURL(url string) (*pstore.PeerInfo, error) {
	multiAddr := maddr.StringCast(url)
	return pstore.InfoFromP2pAddr(multiAddr)
}
