package storageapi

import "syscall"

// GetDeviceNumber returns the physical device ID and storage media type for
// the given volume handle.
func GetDeviceNumber(handle syscall.Handle) {
	//syscall.DeviceIoControl(handle, )
}
