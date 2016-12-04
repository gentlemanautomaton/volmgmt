package volume

import (
	"errors"
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
	handle     syscall.Handle
	devnum     storageapi.DeviceNumber
	descriptor storageapi.DeviceDescriptor
}

// New returns a volume representing the volume of the given path, which must be in one of the
// following formats:
//
//  \\.\X:
//  \\?\Volume{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}\
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
	//access := uint32(0)
	mode := uint32(syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE)
	h, err := syscall.CreateFile(pathp, access, mode, nil, syscall.OPEN_EXISTING, 0, 0)
	if err != nil {
		return nil, err
	}

	v := &Volume{
		handle: h,
	}

	// Query and store the volume's device number
	v.devnum, err = storageapi.GetDeviceNumber(h)
	if err != nil {
		v.Close()
		return nil, err
	}

	// Query and store the volume's device descriptor
	v.descriptor, err = storageapi.QueryDeviceDescriptor(h)
	if err != nil {
		v.Close()
		return nil, err
	}

	return v, nil
}

// Label returns the label of the volume.
func (v *Volume) Label() (string, error) {
	label, _, _, _, _, err := volumeapi.GetVolumeInformationByHandle(v.handle)
	return label, err
}

// DeviceNumber returns the physical device number of the volume.
func (v *Volume) DeviceNumber() uint32 {
	return v.devnum.DeviceNumber
}

// PartitionNumber returns the partition number of the volume.
func (v *Volume) PartitionNumber() int32 {
	return v.devnum.PartitionNumber
}

// DeviceType returns the type of device represented by the volume.
func (v *Volume) DeviceType() uint16 {
	return v.devnum.DeviceType
}

// BusType returns the hardware bus type of the volume.
func (v *Volume) BusType() uint32 {
	return v.descriptor.BusType
}

// RemovableMedia returns true if the volume is located on removable media.
func (v *Volume) RemovableMedia() bool {
	return v.descriptor.RemovableMedia
}

// VendorID returns the hardware vendor ID of the volume.
func (v *Volume) VendorID() string {
	return v.descriptor.VendorID
}

// ProductID returns the hardware product ID of the volume.
func (v *Volume) ProductID() string {
	return v.descriptor.ProductID
}

// ProductRevision returns the hardware product revision of the volume.
func (v *Volume) ProductRevision() string {
	return v.descriptor.ProductRevision
}

// SerialNumber returns the hardware serial number of the volume.
func (v *Volume) SerialNumber() string {
	return v.descriptor.SerialNumber
}

// Paths returns all of the volume's mount points.
func (v *Volume) Paths() ([]string, error) {
	// FIXME: We need the volume name in a form that is usable by
	//        GetVolumePathNamesForVolumeName. Some potential options:
	//        * Call SetupDiGetClassDevs (Requires enumeration? Bleh.)
	//        * GetVolumeNameForVolumeMountPoint?
	return nil, errors.New("volume.Paths() is not yet implemented")
	/*
		name, err := v.Name()
		if err != nil {
			return nil, err
		}

		return volumeapi.GetVolumePathNamesForVolumeName(name)
	*/
}

// Handle returns the system handle of the volume.
func (v *Volume) Handle() syscall.Handle {
	return v.handle
}

// Close releases any resources consumed by the volume.
func (v *Volume) Close() error {
	return syscall.CloseHandle(v.handle)
}
