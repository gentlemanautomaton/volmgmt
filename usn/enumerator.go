package usn

import (
	"io"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/volmgmt/hsync"
	"github.com/gentlemanautomaton/volmgmt/volumeapi"
)

// Enumerator reads records from a master file table.
type Enumerator struct {
	data RawJournalData
	h    *hsync.Handle
	pos  int64 // File reference number or USN
	low  USN
	high USN
}

// NewEnumerator returns a master file table enumerator for the volume described
// by path.
func NewEnumerator(path string, low, high USN) (enumerator *Enumerator, err error) {
	const (
		access = syscall.GENERIC_READ
		mode   = syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE
	)

	h, err := volumeapi.Handle(path, access, mode)
	if err != nil {
		return nil, err
	}

	return NewEnumeratorWithHandle(hsync.New(h), low, high)
}

// NewEnumeratorWithHandle returns a master file table enumerator for the
// volume with the given handle.
//
// When the enumerator is closed its associated handle will also be closed. When
// providing an existing handle that will be used elsewhere be sure to
// clone it first.
func NewEnumeratorWithHandle(handle *hsync.Handle, low, high USN) (*Enumerator, error) {
	data, err := QueryJournal(handle.Handle())
	if err != nil {
		return nil, err
	}

	return &Enumerator{
		h:    handle,
		data: data,
		low:  low,
		high: high,
	}, nil
}

// Read fills the given buffer with data from the master file table. The first
// 8 bytes of the returned data contain a 64 bit file reference number or
// update sequence number.
//
// If no more data is available, io.EOF will be returned.
func (e *Enumerator) Read(p []byte) (n int, err error) {
	opts := RawEnumOptions{
		StartFileReferenceNumber: e.pos,
		Low:             e.low,
		High:            e.high,
		MinMajorVersion: 2,
		MaxMajorVersion: 3,
	}
	length, err := EnumData(e.h.Handle(), opts, p)
	n = int(length)
	if err == nil && length >= 8 {
		// Check the next USN that was returned at the start of the buffer. If it
		// matches the starting USN that we provided then there is no data
		// available.
		next := *(*int64)(unsafe.Pointer(&p[0]))
		if next == e.pos {
			return 0, io.EOF
		}

		// We read something, update the cursor's position
		e.pos = next

		if length == 8 {
			// The cursor moved forward but we didn't get any records back, probably
			// because the reason mask filtered out whatever was written to the
			// journal. This is a normal occurence; treat it as an EOF condition.
			return 0, io.EOF
		}
	}
	return
}

// Next returns a slice of records from the master file table. It returns all unread
// records that are available that can fit within the given buffer.
//
// If there are no more unread records err will be io.EOF.
func (e *Enumerator) Next(buffer []byte) (records []Record, err error) {
	n, err := e.Read(buffer) // Advances the cursor
	if err != nil {
		return
	}

	// Skip the USN at the front of buffer
	buffer = buffer[8:n]
	n -= 8

	for {
		var record Record
		err = record.UnmarshalBinary(buffer)
		if err != nil {
			return
		}

		records = append(records, record)

		n -= int(record.RecordLength)
		if n < recordV2Size {
			// The next record wouldn't fit in the buffer
			return
		}

		buffer = buffer[record.RecordLength:]
	}
}

// Close releases any resources consumed by the enumerator.
func (e *Enumerator) Close() {
	e.h.Close()
}
