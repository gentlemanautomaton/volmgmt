package usn

import (
	"errors"
	"io"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/volmgmt/hsync"
	"github.com/gentlemanautomaton/volmgmt/volumeapi"
)

var (
	// ErrNegativeUSN is returned when seeking to a negative update sequence
	// number.
	ErrNegativeUSN = errors.New("the requested operation would result in a negative update sequence number")

	// ErrInvalidWhence is returned when an invalid or unsupported whence value is
	// supplied to a Seek function.
	ErrInvalidWhence = errors.New("invalid whence value")

	// ErrInsufficientBuffer is returned when a record cannot be read because the
	ErrInsufficientBuffer = errors.New("unable to read record data due to insufficient buffer size")
	// buffer is too small to receive its data.
)

// Cursor provides a more idiomatic means of reading USN change journal data
// through an io.ReadSeeker interface.
type Cursor struct {
	data       RawJournalData
	h          *hsync.Handle
	usn        USN
	reasonMask Reason
	// TODO: Consider adding some sort of buffer (or let the user provide one)
}

// NewCursor returns a USN change journal cursor for the volume described by
// path.
func NewCursor(path string, reasonMask Reason) (cursor *Cursor, err error) {
	const (
		access = syscall.GENERIC_READ
		mode   = syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE
	)

	h, err := volumeapi.Handle(path, access, mode)
	if err != nil {
		return nil, err
	}

	return NewCursorWithHandle(hsync.New(h), reasonMask)
}

// NewCursorWithHandle returns a USN Journal cursor for the volume with the
// given handle.
//
// When the cursor is closed its associated handle will also be closed. When
// providing an existing handle that will be used elsewhere be sure to
// clone it first.
func NewCursorWithHandle(handle *hsync.Handle, reasonMask Reason) (*Cursor, error) {
	data, err := QueryJournal(handle.Handle())
	if err != nil {
		return nil, err
	}

	return &Cursor{
		h:          handle,
		data:       data,
		reasonMask: reasonMask,
	}, nil
}

// Read fills the given buffer with data from the underlying USN journal if any
// is present. The first 8 bytes of the returned data contain a 64 bit USN
// value.
//
// If no more data is currently available, io.EOF will be returned.
func (c *Cursor) Read(p []byte) (n int, err error) {
	opts := RawReadOptions{
		StartUSN:        c.usn,
		ReasonMask:      c.reasonMask,
		JournalID:       c.data.JournalID,
		MinMajorVersion: 2,
		MaxMajorVersion: 3,
	}
	length, err := ReadJournal(c.h.Handle(), opts, p)
	n = int(length)
	if err == nil && length >= 8 {
		// Check the next USN that was returned at the start of the buffer. If it
		// matches the starting USN that we provided then there is no data
		// available.
		next := *(*USN)(unsafe.Pointer(&p[0]))
		if next == c.usn {
			return 0, io.EOF
		}

		// We read something, update the cursor's position
		c.usn = next

		if length == 8 {
			// The cursor moved forward but we didn't get any records back, probably
			// because the reason mask filtered out whatever was written to the
			// journal. This is a normal occurence; treat it as an EOF condition.
			return 0, io.EOF
		}
	}
	return
}

// Seek moves the cursor to the USN specified by offset.
func (c *Cursor) Seek(offset int64, whence int) (usn int64, err error) {
	switch whence {
	case io.SeekStart:
		if offset < 0 {
			err = ErrNegativeUSN
		} else {
			c.usn = USN(offset)
		}
	case io.SeekCurrent:
		if offset < 0 {
			// TODO: Make sure that negative USN values don't occur in normal operation
			if c.usn+USN(offset) < 0 {
				err = ErrNegativeUSN
			} else {
				c.usn += USN(offset)
			}
		} else {
			c.usn += USN(offset)
		}
	default:
		err = ErrInvalidWhence
	}
	usn = int64(c.usn)
	return
}

// Next returns a slice of records from the journal. It returns all unread
// records that are available that can fit within the given buffer.
//
// If there are no more unread records err will be io.EOF.
func (c *Cursor) Next(buffer []byte) (records []Record, err error) {
	n, err := c.Read(buffer) // Advances the cursor
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

// USN returns the current update sequence number that the cursor is pointed to.
func (c *Cursor) USN() USN {
	return c.usn
}

// Close releases any resources consumed by the journal.
func (c *Cursor) Close() {
	c.h.Close()
}
