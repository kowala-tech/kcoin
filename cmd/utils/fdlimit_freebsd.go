// +build freebsd

package utils

import "syscall"

// This file is largely identical to fdlimit_unix.go,
// but Rlimit fields have type int64 on FreeBSD so it needs
// an extra conversion.

// raiseFdLimit tries to maximize the file descriptor allowance of this process
// to the maximum hard-limit allowed by the OS.
func raiseFdLimit(max uint64) error {
	// Get the current limit
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return err
	}
	// Try to update the limit to the max allowance
	limit.Cur = limit.Max
	if limit.Cur > int64(max) {
		limit.Cur = int64(max)
	}
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return err
	}
	return nil
}

// getFdLimit retrieves the number of file descriptors allowed to be opened by this
// process.
func getFdLimit() (int, error) {
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return 0, err
	}
	return int(limit.Cur), nil
}
