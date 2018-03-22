/*
Package p2p implements the Kowala p2p network protocols

Peer Discovery

The network has nodes that are assumed to be always available - bootstrap nodes.
The bootstrap nodes maintain a list of all nodes that connected to them in
period of time. When a node connects to the Kowala network, they first connect
to the bootstrap nodes which share the list of peers.

The list of hardcoded bootstrap nodes per network can be found under
params/bootnodes.go

Peer Selection

Kowala uses a kademlia-like system to discover further peers.
The current codebase supports two peer discovery mechanisms: v4 and v5.
The v5 version enables discovery based on topics - protocols that the
node has enabled (matching protocols).
*/

package p2p
