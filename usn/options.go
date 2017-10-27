package usn

// RawReadOptions are used to specify  parameters when reading from the
// USN journal.
type RawReadOptions struct {
	// Version 0 Fields
	StartUSN          USN
	ReasonMask        Reason
	ReturnOnlyOnClose uint32
	Timeout           int64
	BytesToWaitFor    uint64
	JournalID         uint64
	// Version 1 Fields
	MinMajorVersion uint16
	MaxMajorVersion uint16
}

// RawEnumOptions are used to specify parameters when enumerating the master
// file table.
type RawEnumOptions struct {
	// Version 0 Fields
	StartFileReferenceNumber int64 // Can be a file reference number or USN
	Low                      USN
	High                     USN
	// Version 1 Fields
	MinMajorVersion uint16
	MaxMajorVersion uint16
}
