package utils

import "errors"

// raiseFdLimit tries to maximize the file descriptor allowance of this process
// to the maximum hard-limit allowed by the OS.
func raiseFdLimit(max uint64) error {
	// This method is NOP by design:
	//  * Linux/Darwin counterparts need to manually increase per process limits
	//  * On Windows Go uses the CreateFile API, which is limited to 16K files, non
	//    changeable from within a running process
	// This way we can always "request" raising the limits, which will either have
	// or not have effect based on the platform we're running on.
	if max > 16384 {
		return errors.New("file descriptor limit (16384) reached")
	}
	return nil
}

// getFdLimit retrieves the number of file descriptors allowed to be opened by this
// process.
func getFdLimit() (int, error) {
	// Please see raiseFdLimit for the reason why we use hard coded 16K as the limit
	return 16384, nil
}
