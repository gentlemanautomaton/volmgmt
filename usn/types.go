package usn

// RawJournalData contains the raw output of USN journal information.
type RawJournalData struct {
	// Version 0 Fields
	JournalID       uint64
	FirstUSN        USN
	NextUSN         USN
	LowestValidUSN  USN
	MaxUSN          USN
	MaximumSize     uint64
	AllocationDelta uint64
	// Version 1 Fields
	MinSupportedMajorVersion uint16
	MaxSupportedMajorVersion uint16
	// Version 2 Fields
	Flags                       uint32
	RangeTrackChunkSize         uint64
	RangeTrackFileSizeThreshold int64
}
