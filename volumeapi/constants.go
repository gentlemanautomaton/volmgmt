// +build windows

package volumeapi

import "syscall"

const (
	// ErrMoreData indicates that a buffer was too small to hold the
	// data that is to be recieved.
	ErrMoreData = syscall.Errno(234)
)

const (
	// MaxVolumeNameLength is the maximum number of characters in a volume name.
	MaxVolumeNameLength = 261

	// MaxFileSystemNameLength is the maximum number of characters in a file
	// system name.
	MaxFileSystemNameLength = 261
)
