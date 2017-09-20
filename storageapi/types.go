package storageapi

// DeviceNumber describes the device type and device number of storage devices,
// and can be used to construct the physical address of mass storage devices
// and logical volumes. For logical volumes it also holds the partition number
// of the volume.
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

// PropertyQuery is used internally in calls to query storage devices.
type PropertyQuery struct {
	PropertyID uint32
	QueryType  uint32
	_          byte
}

// DeviceDescriptor describes a storage device.
type DeviceDescriptor struct {
	DeviceType         byte
	DeviceTypeModifier byte
	RemovableMedia     bool
	CommandQueueing    bool
	VendorID           string
	ProductID          string
	ProductRevision    string
	SerialNumber       string
	BusType            uint32
}

// RawDeviceDescriptorHeader is used internally for system calls that query
// storage device descriptors.
type RawDeviceDescriptorHeader struct {
	Version uint32
	Size    uint32
}

// RawDeviceDescriptor is used internally for system calls that query storage
// device descriptors.
type RawDeviceDescriptor struct {
	Version               uint32
	Size                  uint32
	DeviceType            byte
	DeviceTypeModifier    byte
	RemovableMedia        bool
	CommandQueueing       bool
	VendorIDOffset        uint32
	ProductIDOffset       uint32
	ProductRevisionOffset uint32
	SerialNumberOffset    uint32
	BusType               uint32
	RawPropertiesLength   uint32
}

// AccessAlignmentDescriptor describes the byte alignment characteristics of a
// storage device.
type AccessAlignmentDescriptor struct {
	Version                       uint32
	Size                          uint32
	BytesPerCacheLine             uint32
	BytesOffsetForCacheAlignment  uint32
	BytesPerLogicalSector         uint32
	BytesPerPhysicalSector        uint32
	BytesOffsetForSectorAlignment uint32
}
