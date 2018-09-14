package usn

import (
	"syscall"

	"github.com/gentlemanautomaton/volmgmt/hsync"
	"github.com/gentlemanautomaton/volmgmt/volumeapi"
)

// Journal provides access to USN journal information and records.
type Journal struct {
	h *hsync.Handle
}

// NewJournal returns a USN Journal accessor for the volume with the given path.
func NewJournal(path string) (*Journal, error) {
	const (
		access = syscall.GENERIC_READ
		mode   = syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE
	)

	h, err := volumeapi.Handle(path, access, mode)
	if err != nil {
		return nil, err
	}

	return NewJournalWithHandle(hsync.New(h)), nil
}

// NewJournalWithHandle returns a USN Journal accessor for the volume with the
// given handle.
//
// When the journal is closed its associated handle will also be closed. When
// providing an existing handle that will be used elsewhere be sure to
// clone it first.
//
// NewJournal does not force the creation of a change journal when one does not
// already exist on the volume. To bring a new journal into existence call
// Journal.Create().
func NewJournalWithHandle(handle *hsync.Handle) *Journal {
	return &Journal{
		h: handle,
	}
}

// Create will either
/*
func (j *Journal) Create() {
	volumeapi.CreateUSNJournal(j.handle)
}
*/

// Query returns information about the current condition of the change journal.
func (j *Journal) Query() (data RawJournalData, err error) {
	return QueryJournal(j.h.Handle())
}

// Cursor returns a new cursor for the journal.
//
// If filer is non-nil, it will be used to return records with a populated
// path field.
func (j *Journal) Cursor(processor Processor, reasonMask Reason, filter Filter, filer Filer) (*Cursor, error) {
	return NewCursorWithHandle(j.h.Clone(), processor, reasonMask, filter, filer)
}

// MFT returns an MFT for the journal.
func (j *Journal) MFT() *MFT {
	return NewMFTWithHandle(j.h.Clone())
}

// Cache builds up a cache of MFT records matching the given filter with USN
// values between low and high, inclusive.
func (j *Journal) Cache(filter Filter, low, high USN) (*Cache, error) {
	mft := j.MFT()
	defer mft.Close()

	iter, err := mft.Enumerate(filter, low, high)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	cache := NewCache()
	err = cache.ReadFrom(iter)
	if err != nil {
		return nil, err
	}
	return cache, nil
}

// Monitor returns a new monitor for the journal.
func (j *Journal) Monitor() *Monitor {
	return NewMonitorWithHandle(j.h.Clone())
}

// Close releases any resources consumed by the journal.
func (j *Journal) Close() {
	j.h.Close()
}
