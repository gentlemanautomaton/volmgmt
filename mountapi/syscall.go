package mountapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/gentlemanautomaton/volmgmt/ioctl"
)

const (
	valueBufSize = 1024
)

// QueryStableGUID returns the stable globally unique ID returned by device
// and storage class drivers.
func QueryStableGUID(handle syscall.Handle) (guid windows.GUID, err error) {
	var length uint32
	err = syscall.DeviceIoControl(handle, ioctl.MountDevQueryStableGUID, nil, 0, (*byte)(unsafe.Pointer(&guid)), uint32(unsafe.Sizeof(guid)), &length, nil)
	return
}

// QueryUniqueID returns the unique ID returned by device and storage class
// drivers.
func QueryUniqueID(handle syscall.Handle) (id []byte, err error) {
	return queryBytes(handle, ioctl.MountDevQueryUniqueID)
}

// QueryDeviceName returns the name returned by device and storage class
// drivers.
func QueryDeviceName(handle syscall.Handle) (name string, err error) {
	b, err := queryBytes(handle, ioctl.MountDevQueryDeviceName)
	if err != nil {
		return
	}
	name = utf16BytesToString(b)
	return
}

/*
// QueryMountPoints retrieves returns the set of mount points for the storage
// device represented by the provided handle.
func QueryMountPoints() {
	// See: https://stackoverflow.com/questions/42426931/how-to-get-from-device-manager-device-e-g-from-its-physical-device-object-nam
	var handle syscall.Handle // TODO: Get a handle to the mount manager
	var length uint32
	err = syscall.DeviceIoControl(handle, ioctl.MountMgrQueryPoints, nil, 0, (*byte)(unsafe.Pointer(&devnum)), uint32(unsafe.Sizeof(devnum)), &length, nil)
	return
}
*/

func queryBytes(handle syscall.Handle, ioControlCode uint32) (value []byte, err error) {
	const valueOffset = 2 // Byte offset of value within the buffer; i.e. sizeof(header)

	type valueHeader struct {
		Length uint16
	}

	var (
		length uint32
		b      [valueBufSize]byte
		buffer = b[:]
	)

	// Make up to 3 attempts to get the name data.
	//
	// 1: Using a fixed buffer allocated on the stack
	// 2: Using a dynamic buffer based on the reported size (first attempt)
	// 3: Using a dynamic buffer based on the reported size (second attempt)
	//
	// It is extremely unlikely, but feasible, that the length could change
	// between calls if a name changes.

	for i := 0; i < 3; i++ {
		length, err = queryName(handle, ioControlCode, buffer)
		if err != nil && err != syscall.ERROR_INSUFFICIENT_BUFFER {
			return
		}

		hdr := (*valueHeader)(unsafe.Pointer(&buffer[0]))
		if valueOffset+hdr.Length <= uint16(len(buffer)) {
			buffer = buffer[valueOffset:length]
			break
		}

		err = syscall.ERROR_INSUFFICIENT_BUFFER

		buffer = make([]byte, hdr.Length)
	}
	if err != nil {
		return
	}

	value = make([]byte, len(buffer))
	copy(value, buffer)

	return
}

func queryName(handle syscall.Handle, ioControlCode uint32, buffer []byte) (length uint32, err error) {
	var p1 *byte
	s1 := uint32(len(buffer))
	length = s1
	if s1 > 0 {
		p1 = &buffer[0]
	}
	err = syscall.DeviceIoControl(handle, ioControlCode, nil, 0, p1, s1, &length, nil)
	return
}
