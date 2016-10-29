package ioctl

// Code returns an IO Control Code for the given parameters.
func Code(deviceType, function uint16, method, access uint8) uint32 {
	return uint32(deviceType)<<16 | uint32(access)<<14 | uint32(function)<<2 | uint32(method)
}
