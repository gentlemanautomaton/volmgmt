package storageapi

import (
	"strings"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/volmgmt/ioctl"
)

const (
	descriptorBufSize    = 1024
	descriptorBufMaxSize = 1024 * 1024 // 1 MiB maximum size that we're willing to allocate
)

// GetDeviceNumber retrieves the device numbering of the storage device
// represented by the provided handle. The returned information includes the
// device type and physical device number, as well as the partition number if
// the device is a partitioned volume.
func GetDeviceNumber(handle syscall.Handle) (dev DeviceNumber, err error) {
	var length uint32
	err = syscall.DeviceIoControl(handle, ioctl.StorageGetDeviceNumber, nil, 0, (*byte)(unsafe.Pointer(&dev)), uint32(unsafe.Sizeof(dev)), &length, nil)
	return
}

// QueryDeviceDescriptor retrieves the device descriptor of the storage device
// represented by the provided handle.
func QueryDeviceDescriptor(handle syscall.Handle) (descriptor DeviceDescriptor, err error) {
	var (
		length uint32
		b      [descriptorBufSize]byte
		buffer = b[:]
	)

	// Make up to 3 attempts to get the raw descriptor data.
	//
	// 1: Using a fixed buffer allocated on the stack
	// 2: Using a dynamic buffer based on the reported size (first attempt)
	// 3: Using a dynamic buffer based on the reported size (second attempt)
	//
	// It is extremely unlikely, but feasible, that the length could change
	// between calls if a device descriptor changes.

	for i := 0; i < 3; i++ {
		length, err = QueryProperty(handle, StorageDeviceProperty, PropertyStandardQuery, buffer)
		if err != nil && err != syscall.ERROR_INSUFFICIENT_BUFFER {
			return
		}

		hdr := (*RawDeviceDescriptorHeader)(unsafe.Pointer(&buffer[0]))
		if hdr.Size <= uint32(len(buffer)) {
			buffer = buffer[:length]
			break
		}

		err = syscall.ERROR_INSUFFICIENT_BUFFER

		if hdr.Size > descriptorBufMaxSize {
			return
		}

		buffer = make([]byte, hdr.Size)
	}
	if err != nil {
		return
	}

	raw := (*RawDeviceDescriptor)(unsafe.Pointer(&buffer[0]))

	descriptor.DeviceType = raw.DeviceType
	descriptor.DeviceTypeModifier = raw.DeviceTypeModifier
	descriptor.RemovableMedia = raw.RemovableMedia
	descriptor.CommandQueueing = raw.CommandQueueing
	descriptor.BusType = raw.BusType

	if raw.VendorIDOffset > 0 {
		descriptor.VendorID = strings.TrimSpace(asciiToString(buffer[raw.VendorIDOffset:]))
	}
	if raw.ProductIDOffset > 0 {
		descriptor.ProductID = strings.TrimSpace(asciiToString(buffer[raw.ProductIDOffset:]))
	}
	if raw.ProductRevisionOffset > 0 {
		descriptor.ProductRevision = strings.TrimSpace(asciiToString(buffer[raw.ProductRevisionOffset:]))
	}
	if raw.SerialNumberOffset > 0 {
		descriptor.SerialNumber = strings.TrimSpace(asciiToString(buffer[raw.SerialNumberOffset:]))
	}

	return
}

// QueryProperty is a low level API call that retrieves many different kinds
// of storage device properties. It is generally prefable to use one of the
// higher level functions like QueryDeviceProperty to retrieve a particular
// kind of data.
func QueryProperty(handle syscall.Handle, propertyID, queryType uint32, buffer []byte) (length uint32, err error) {
	pq := PropertyQuery{
		PropertyID: StorageDeviceProperty,
		QueryType:  queryType,
	}
	p0 := (*byte)(unsafe.Pointer(&pq))
	s0 := uint32(unsafe.Sizeof(pq))

	var p1 *byte
	s1 := uint32(len(buffer))
	if s1 > 0 {
		p1 = &buffer[0]
	}

	err = syscall.DeviceIoControl(handle, ioctl.StorageQueryProperty, p0, s0, p1, s1, &length, nil)

	return
}
