package usn

// Iter is an iterable source of change journal records.
//
// The provided buffer is used to hold raw journal data temporarily.
// The unmarshaled records that are returned will be appended to data.
type Iter interface {
	Next(buffer []byte, data []Record) (records []Record, err error)
}
