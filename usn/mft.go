package usn

import (
	"errors"
	"syscall"

	"github.com/gentlemanautomaton/volmgmt/fileref"
	"github.com/gentlemanautomaton/volmgmt/hsync"
	"github.com/gentlemanautomaton/volmgmt/volumeapi"
)

const mftBufSize = 8192

// MFT provides access to the master file table.
type MFT struct {
	h *hsync.Handle
}

// NewMFT returns a master file table accessor for the volume with the given
// path.
func NewMFT(path string) (*MFT, error) {
	const (
		access = syscall.GENERIC_READ
		mode   = syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE
	)

	h, err := volumeapi.Handle(path, access, mode)
	if err != nil {
		return nil, err
	}

	return NewMFTWithHandle(hsync.New(h)), nil
}

// NewMFTWithHandle returns a master file table accessor for the volume with the
// given handle.
//
// When the journal is closed its associated handle will also be closed. When
// providing an existing handle that will be used elsewhere be sure to
// clone it first.
func NewMFTWithHandle(handle *hsync.Handle) *MFT {
	return &MFT{
		h: handle,
	}
}

// File attempts to locate the file with the given ID in the master file table.
func (mft *MFT) File(id fileref.ID) (record Record, err error) {
	if !id.IsInt64() {
		err = errors.New("unable to search for files with 128-bit reference numbers")
		return
	}

	opts := RawEnumOptions{
		StartFileReferenceNumber: id.Int64(),
		MinMajorVersion:          2,
		MaxMajorVersion:          3,
		Low:                      Min,
		High:                     Max,
	}

	var (
		buffer [mftBufSize]byte
		b      = buffer[:]
	)

	length, err := EnumData(mft.h.Handle(), opts, b)
	if err != nil {
		return
	}

	if length <= 8 {
		// No data returned
		err = errors.New("insufficient buffer")
		return
	}

	// Skip the USN at the front of buffer
	b = b[8:length]

	err = record.UnmarshalBinary(b)
	return
}

// Enumerate returns a new enumerator for the master file table. It will
// return records with update sequence numbers between low and high, inclusive.
func (mft *MFT) Enumerate(low, high USN) (*Enumerator, error) {
	return NewEnumeratorWithHandle(mft.h.Clone(), low, high)
}

// Close releases any resources consumed by the MFT.
func (mft *MFT) Close() {
	mft.h.Close()
}
