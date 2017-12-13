package state

import (
	"bytes"
	"testing"

	"github.com/kowala-tech/kUSD/common"
)

// Tests that the node iterator indeed walks over the entire database contents.
func TestNodeIteratorCoverage(t *testing.T) {
	// Create some arbitrary test state to iterate
	db, mem, root, _ := makeTestState()

	state, err := New(root, db)
	if err != nil {
		t.Fatalf("failed to create state trie at %x: %v", root, err)
	}
	// Gather all the node hashes found by the iterator
	hashes := make(map[common.Hash]struct{})
	for it := NewNodeIterator(state); it.Next(); {
		if it.Hash != (common.Hash{}) {
			hashes[it.Hash] = struct{}{}
		}
	}

	// Cross check the hashes and the database itself
	for hash := range hashes {
		if _, err := mem.Get(hash.Bytes()); err != nil {
			t.Errorf("failed to retrieve reported node %x: %v", hash, err)
		}
	}
	for _, key := range mem.Keys() {
		if bytes.HasPrefix(key, []byte("secure-key-")) {
			continue
		}
		if _, ok := hashes[common.BytesToHash(key)]; !ok {
			t.Errorf("state entry not reported %x", key)
		}
	}
}
