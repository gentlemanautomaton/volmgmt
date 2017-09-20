package volume

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"

	"github.com/gentlemanautomaton/volmgmt/guidconv"
	"github.com/gentlemanautomaton/volmgmt/mountapi"
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

// Name returns the volume GUID name.
//
// BUG: This currently only works for device drivers that supply a stable GUID.
func (v *Volume) Name() (string, error) {
	guid, err := v.GUID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("\\\\?\\Volume%s\\", strings.ToLower(guidconv.Format(&guid))), nil
}

// GUID returns a GUID for the volume that is supplied by the mount manager.
//
// If the underlying device driver supplies a stable GUID, the mount manager
// will use and return that value.
//
// BUG: This currently only works for device drivers that supply a stable GUID.
func (v *Volume) GUID() (guid windows.GUID, err error) {
	guid, err = v.StableGUID()
	if err == nil {
		return
	}
	// TODO: Query the mount manager
	return
}

// StableGUID returns a stable GUID for the volume that is supplied by its
// device driver.
//
// Not all device drivers are capable of supplying a stable GUID. If this
// device doesn't supply one an error will be returned. In such cases the
// mount manager will generate a GUID for the volume, which can be accessed
// via v.GUID().
func (v *Volume) StableGUID() (guid windows.GUID, err error) {
	guid, err = mountapi.QueryStableGUID(v.handle)
	if err != nil {
		err = fmt.Errorf("unable to retrieve stable GUID: %v", err)
	}
	return
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

// DeviceID returns the device ID of the volume.
func (v *Volume) DeviceID() ([]byte, error) {
	return mountapi.QueryUniqueID(v.handle)
}

// DevicePath returns an NT namespace device path for the volume.
func (v *Volume) DevicePath() (string, error) {
	return mountapi.QueryDeviceName(v.handle)
}

// Paths returns all of the volume's mount points.
func (v *Volume) Paths() ([]string, error) {
	// TODO: Consider using IOCTL_MOUNTMGR_QUERY_POINTS instead
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
