package storageapi

const (
	// PartitionNotAvailable is the partition number returned for devices that
	// do not have partitions. This value may be returned from calls to
	// GetDeviceNumber.
	PartitionNotAvailable = -1
)

// STORAGE_PROPERTY_ID enumeration.
const (
	StorageDeviceProperty                 = iota // 0 StorageDeviceProperty
	StorageAdapterProperty                       // 1 StorageAdapterProperty
	StorageDeviceIDProperty                      // 2 StorageDeviceIdProperty
	StorageDeviceUniqueIDProperty                // 3 StorageDeviceUniqueIdProperty
	StorageDeviceWriteCacheProperty              // 4 StorageDeviceWriteCacheProperty
	StorageMiniportProperty                      // 5 StorageMiniportProperty
	StorageAccessAlignmentProperty               // 6 StorageAccessAlignmentProperty
	StorageDeviceSeekPenaltyProperty             // 7 StorageDeviceSeekPenaltyProperty
	StorageDeviceTrimProperty                    // 8 StorageDeviceTrimProperty
	StorageDeviceWriteAggregationProperty        // 9 StorageDeviceWriteAggregationProperty
	StorageDeviceDeviceTelemetryProperty         // 10 StorageDeviceDeviceTelemetryProperty
	StorageDeviceLBProvisioningProperty          // 11 StorageDeviceLBProvisioningProperty
	StorageDevicePowerProperty                   // 12 StorageDevicePowerProperty
	StorageDeviceCopyOffloadProperty             // 13 StorageDeviceCopyOffloadProperty
	StorageDeviceResiliencyProperty              // 14 StorageDeviceResiliencyProperty
	StorageDeviceMediumProductType               // 15 StorageDeviceMediumProductType
)

// STORAGE_PROPERTY_ID enumeration.
const (
	StorageDeviceIOCapabilityProperty      = iota + 48 // 48 StorageDeviceIoCapabilityProperty
	StorageAdapterProtocolSpecificProperty             // 49 StorageAdapterProtocolSpecificProperty
	StorageDeviceProtocolSpecificProperty              // 50 StorageDeviceProtocolSpecificProperty
	StorageAdapterTemperatureProperty                  // 51 StorageAdapterTemperatureProperty
	StorageDeviceTemperatureProperty                   // 52 StorageDeviceTemperatureProperty
	StorageAdapterPhysicalTopologyProperty             // 53 StorageAdapterPhysicalTopologyProperty
	StorageDevicePhysicalTopologyProperty              // 54 StorageDevicePhysicalTopologyProperty
	StorageDeviceAttributesProperty                    // 55 StorageDeviceAttributesProperty
)

// STORAGE_QUERY_TYPE enumeration.
const (
	PropertyStandardQuery   = iota // 0
	PropertyExistsQuery            // 1
	PropertyMaskQuery              // 2
	PropertyQueryMaxDefined        // 3
)
