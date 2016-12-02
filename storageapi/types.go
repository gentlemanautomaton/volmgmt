package storageapi

// DeviceNumber describes the device type and device number of any device. It is
// particularly useful for identifying the physical address of mass storage
// devices and partitioned volumes. For volumes it also holds the partition
// number of the volume.
type DeviceNumber struct {
	// DeviceType is a device type code as defined in the ioctlcode package.
	DeviceType uint16

	// DeviceNumber is the physical device number of the device.
	DeviceNumber uint32

	// PartitionNumber is the partition number of the device if it represents
	// a partitioned volume. If the device is not a partitioned volume it will
	// be set to PartitionNotAvailable.
	//
	// PartitionNumber is formally defined as a ulong in the official
	// documentation on MSDN, but the documentation also clearly states that a
	// value of -1 is returned for devices that are not parititioned, so the
	// package authors have elected to change its type from uint32 to int32 here.
	// See the MSDN documentation for details:
	//
	// https://msdn.microsoft.com/library/bb968801
	//
	// Users of this package should test its value against the
	// PartitionNotAvailable constant, which is guaranteed to match the "missing"
	// value, however the GetDeviceNumber function in this package will return it.
	PartitionNumber int32
}
