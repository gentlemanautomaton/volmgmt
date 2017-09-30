package usn

// RawReadOptions are used to specify optional parameters when reading from the
// USN journal.
type RawReadOptions struct {
	// Version 0 Fields
	StartUSN          USN
	ReasonMask        uint32
	ReturnOnlyOnClose uint32
	Timeout           int64
	BytesToWaitFor    uint64
	JournalID         uint64
	// Version 1 Fields
	MinMajorVersion uint16
	MaxMajorVersion uint16
}
