package utils

import "testing"

// TestFileDescriptorLimits simply tests whether the file descriptor allowance
// per this process can be retrieved.
func TestFileDescriptorLimits(t *testing.T) {
	target := 4096

	if limit, err := getFdLimit(); err != nil || limit <= 0 {
		t.Fatalf("failed to retrieve file descriptor limit (%d): %v", limit, err)
	}
	if err := raiseFdLimit(uint64(target)); err != nil {
		t.Fatalf("failed to raise file allowance")
	}
	if limit, err := getFdLimit(); err != nil || limit < target {
		t.Fatalf("failed to retrieve raised descriptor limit (have %v, want %v): %v", limit, target, err)
	}
}
