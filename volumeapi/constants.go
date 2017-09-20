package volumeapi

import "syscall"

const (
	// ErrMoreData indicates that a buffer was too small to hold the
	// data that is to be recieved.
	ErrMoreData = syscall.Errno(234)
)

const (
	// MaxVolumeLabelLength is the maximum number of characters in a volume label.
	MaxVolumeLabelLength = 261

	// MaxVolumeNameLength is the maximum number of characters in a volume name.
	//
	// Volume names are also known as volume GUID paths, and are based on the
	// globally unique identifier of the volume. Volume names are not supplied by
	// or editable by users, unlike volume labels.
	//
	// Volume names will always be of this form:
	//
	//   \\?\Volume{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}\
	MaxVolumeNameLength = 50

	// MaxFileSystemNameLength is the maximum number of characters in a file
	// system name.
	MaxFileSystemNameLength = 261
)
