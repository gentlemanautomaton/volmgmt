package ioctlcode

// New returns an I/O control code for the given parameters.
func New(deviceType, function uint16, method, access uint8) uint32 {
	return uint32(deviceType)<<16 | uint32(access)<<14 | uint32(function)<<2 | uint32(method)
}
