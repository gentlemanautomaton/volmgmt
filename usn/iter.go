package usn

// Iter is an iterable source of change journal records.
type Iter interface {
	Next(buffer []byte) (records []Record, err error)
}
