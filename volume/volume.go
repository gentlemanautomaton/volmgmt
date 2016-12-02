package volume

import (
	"syscall"

	"github.com/gentlemanautomaton/volmgmt/storageapi"
	"github.com/gentlemanautomaton/volmgmt/volumeapi"
)

/*
type Volume interface {
	Name() (string, error)
	Paths() ([]string, error)
	Handle() syscall.Handle
	Close() error
}
*/

// Volume represents a storage volume. It must be created with a call to New.
type Volume struct {
	handle syscall.Handle
	n      storageapi.DeviceNumber
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
func New(path string) (*Volume, error) {
	if len(path) == 0 {
		return nil, syscall.ERROR_FILE_NOT_FOUND
	}

	// Create volume handle
	pathp, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}
	//access := uint32(syscall.GENERIC_READ | syscall.GENERIC_WRITE)
	access := uint32(syscall.GENERIC_READ)
	mode := uint32(syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE)
	h, err := syscall.CreateFile(pathp, access, mode, nil, syscall.OPEN_EXISTING, 0, 0)
	if err != nil {
		return nil, err
	}

	v := &Volume{
		handle: h,
	}

	// Query and store the volume's device number
	v.n, err = storageapi.GetDeviceNumber(h)
	if err != nil {
		v.Close()
		return nil, err
	}

	return v, nil
}

// Name returns the label of the volume.
func (v *Volume) Name() (string, error) {
	name, _, _, _, _, err := volumeapi.GetVolumeInformationByHandle(v.handle)
	return name, err
}

// DeviceNumber returns the physical device number of the volume.
func (v *Volume) DeviceNumber() uint32 {
	return v.n.DeviceNumber
}

// PartitionNumber returns the partition number of the volume.
func (v *Volume) PartitionNumber() int32 {
	return v.n.PartitionNumber
}

// DeviceType returns the type of device represented by the volume.
func (v *Volume) DeviceType() uint16 {
	return v.n.DeviceType
}

// Paths returns all of the volume's mount points.
func (v *Volume) Paths() ([]string, error) {
	name, err := v.Name()
	if err != nil {
		return nil, err
	}

	return volumeapi.GetVolumePathNamesForVolumeName(name)
}

// Handle returns the system handle of the volume.
func (v *Volume) Handle() syscall.Handle {
	return v.handle
}

// Close releases any resources consumed by the volume.
func (v *Volume) Close() error {
	return syscall.CloseHandle(v.handle)
}
