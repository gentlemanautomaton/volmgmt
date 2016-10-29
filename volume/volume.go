package volume

import "syscall"

// Volume represents a storage volume.
type Volume interface {
	Name() string
	Paths() []string
	Handle() syscall.Handle
	Close() error
}

type volume struct {
	handle syscall.Handle
}

// New returns a volume representing the volume of the given path, which must be in one of the
// following formats:
//
//  \\.\X:
//  \\?\Volume{GUID}\
//  \\.\PhysicalDrive0
//
// The returned volume will wrap a system handle and will consume system
// resources until the volume is closed. It is the caller's responsibility to
// close the volume when finished with it.
func New(path string) (Volume, error) {
	if len(path) == 0 {
		return nil, syscall.ERROR_FILE_NOT_FOUND
	}
	pathp, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}
	sharemode := uint32(syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE)
	h, err := syscall.CreateFile(pathp, 0, sharemode, nil, syscall.OPEN_EXISTING, 0, 0)
	if err != nil {
		return nil, err
	}
	// TODO: Query and store the volume information
	return &volume{
		handle: h,
	}, nil
}

// Name returns the name of the volume.
func (v *volume) Name() string {
	// TODO: Call GetVolumeInformationByHandleW()
	return ""
}

// Paths returns all of the volume's mount points.
func (v *volume) Paths() []string {
	// TODO: Call GetVolumePathNamesForVolumeName()
	return nil
}

// Handle returns the system handle of the volume.
func (v *volume) Handle() syscall.Handle {
	return v.handle
}

// Close releases any resources consumed by the volume.
func (v *volume) Close() error {
	return syscall.CloseHandle(v.handle)
}
