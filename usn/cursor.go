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
	// buffer is too small to receive its data.
	ErrInsufficientBuffer = errors.New("unable to read record data due to insufficient buffer size")
)

// Cursor provides a more idiomatic means of reading USN change journal data
// through an io.ReadSeeker interface.
//
// TODO: Attempt to merge Enumerator and Cursor into one type.
type Cursor struct {
	data       RawJournalData
	h          *hsync.Handle
	usn        USN
	processor  Processor
	reasonMask Reason
	filter     Filter
	filer      Filer
	total      Stats
	filtered   Stats
	// TODO: Consider adding some sort of buffer (or let the user provide one)
}

// NewCursor returns a USN change journal cursor for the volume described by
// path. Only records matching the provided reason mask will be returned.
//
// If filer is non-nil, it will be used to return records with a populated
// path field.
//
// TODO: Make processors, filters and filers fulfill a CursorOption interface, then
// accept a variadic set of options.
func NewCursor(path string, reasonMask Reason, processor Processor, filter Filter, filer Filer) (cursor *Cursor, err error) {
	const (
		access = syscall.GENERIC_READ
		mode   = syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE
	)

	h, err := volumeapi.Handle(path, access, mode)
	if err != nil {
		return nil, err
	}

	return NewCursorWithHandle(hsync.New(h), processor, reasonMask, filter, filer)
}

// NewCursorWithHandle returns a USN Journal cursor for the volume with the
// given handle.
//
// When the cursor is closed its associated handle will also be closed. When
// providing an existing handle that will be used elsewhere be sure to
// clone it first.
func NewCursorWithHandle(handle *hsync.Handle, processor Processor, reasonMask Reason, filter Filter, filer Filer) (*Cursor, error) {
	data, err := QueryJournal(handle.Handle())
	if err != nil {
		return nil, err
	}

	return &Cursor{
		h:          handle,
		data:       data,
		processor:  processor,
		reasonMask: reasonMask,
		filter:     filter,
		filer:      filer,
	}, nil
}

// Read fills the given buffer with data from the underlying USN journal if any
// is present. The first 8 bytes of the returned data contain a 64 bit USN
// value.
//
// Read does not apply the cursor's filter. To retrieve filtered records call
// Next instead.
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

/*
func (c *Cursor) Summary(buffer []byte) {
	for {
		var records []Record
		records, err = c.Next(buffer[:])
		if err != nil {
			break
		}

		for r := range records {
			record := &records[r]
			if record.USN < start {
				continue
			}
			if record.USN > end {
				break current
			}
			if pass == 0 {
				count++
			} else {
				count--
				if count == 0 {
					pos = record.USN
					fmt.Printf("Found: %d/%d\n", count, loop)
					return
				}
			}
		}
	}
}
*/

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

		matched := c.process(&record, c.filter, c.filer)
		if matched {
			records = append(records, record)
		}

		n -= int(record.RecordLength)
		if n < recordV2Size {
			// The next record wouldn't fit in the buffer
			return
		}

		buffer = buffer[record.RecordLength:]
	}
}

// End causes the cursor to read forward until there are no more records left
// to be read. The cursor's statistics will be updated to reflect the ro
func (c *Cursor) End(buffer []byte) {
	return
}

// USN returns the current update sequence number that the cursor is pointed to.
func (c *Cursor) USN() USN {
	return c.usn
}

// Stats returns the current total and filtered statistics for the cursor.
//
// Only records processed by calls to Next or End will be included in the
// returned statistics.
func (c *Cursor) Stats() (total, filtered Stats) {
	total, filtered = c.total, c.filtered
	return
}

// Close releases any resources consumed by the journal.
func (c *Cursor) Close() {
	c.h.Close()
}

// process performs record post-processing after it has been marshaled.
func (c *Cursor) process(record *Record, filter Filter, filer Filer) (matched bool) {
	c.processor.Process(*record)

	if filer != nil && !record.ParentFileReferenceNumber.IsZero() {
		record.Path = record.FileName
		parents, pErr := filer.Parents(*record)
		if pErr == nil {
			for p := range parents {
				record.Path = parents[p].FileName + `\` + record.Path
			}
		}
	}

	c.total.Add(record)

	if filter == nil || filter(*record) {
		matched = true
		c.filtered.Add(record)
	}
	return
}
