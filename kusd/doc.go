/*

Content Dissemination

Terminology

* Chunk - The basic unit in which the content is divided.
* Hash - The result of applying a cryptographic hash function.

Libswift protocol - Faster Block Propagation

Generic protocol which runs directly on top of TCP, providing quick block broadcasts.
Instead of sending full blocks, validators dispatch block fragments. The
proposal includes the computed Merkle root hash and the validators verify if the
fragments sent are part of the block. The block is finally assembled as soon
as a validator contains all the chunks of data.

Messages & Consensus

The proposal message (ProposalMsg) carries the cryptographic hash that is
necessary for other validators to check the integrity of the chunk.

Content Integrity Verification

Content integrity is provided by employing Merkle hash trees. Assuming that a
validator receives the root hash of the block, it can check the integrity if any
chunk of that content.

References

* Performance analysis of the Libswift P2P streaming protocol
- http://ieeexplore.ieee.org/document/6335790/?reload=true
* https://raw.githubusercontent.com/libswift/libswift/master/doc/draft-ietf-ppsp-peer-protocol-00.txt

*/

package kusd
