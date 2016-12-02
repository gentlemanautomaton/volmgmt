package storageapi

import (
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/volmgmt/ioctl"
)

// GetDeviceNumber retrieves the device numbering of the device represented by
// the provided handle. The returned information includes the device type and
// physical device number, as well as the partition number if the device
// handle represents a partitioned volume.
func GetDeviceNumber(handle syscall.Handle) (dev DeviceNumber, err error) {
	var length uint32
	err = syscall.DeviceIoControl(handle, ioctl.StorageGetDeviceNumber, nil, 0, (*byte)(unsafe.Pointer(&dev)), uint32(unsafe.Sizeof(dev)), &length, nil)
	return
}
