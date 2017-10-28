package fileref

// Descriptor is a file reference descriptor that can be used in file system
// API calls.
type Descriptor struct {
	Size uint32
	Type uint32
	Data [16]byte
}
