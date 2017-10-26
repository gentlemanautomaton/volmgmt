package usn

import (
	"errors"
	"fmt"
	"time"
	"unsafe"

	"github.com/gentlemanautomaton/volmgmt/fileattr"
	"github.com/gentlemanautomaton/volmgmt/fileref"

	"golang.org/x/sys/windows"
)

const (
	recordHeaderSize = 8
	recordV2Size     = 56
	recordV3Size     = 60
	recordV4Size     = 80
)

var (
	// ErrTruncatedRecord is returned when the data buffer containing a USN
	// record appears to have been truncated.
	ErrTruncatedRecord = errors.New("USN record data was truncated")

	// ErrInvalidRecordLength is returned when the data buffer containing a USN
	// record contains has an invalid record length value.
	ErrInvalidRecordLength = errors.New("USN record length is invalid (possible data corruption)")
)

// Record represents a change journal record.
type Record struct {
	RecordLength              uint32
	MajorVersion              uint16
	MinorVersion              uint16
	FileReferenceNumber       fileref.ID
	ParentFileReferenceNumber fileref.ID
	USN                       USN
	TimeStamp                 time.Time
	Reason                    Reason
	SourceInfo                uint32
	FileAttributes            fileattr.Value
	FileName                  string
}

// UnmarshalBinary attempts to parse a single record from the given data.
func (r *Record) UnmarshalBinary(data []byte) error {
	bufSize := len(data)
	if bufSize < recordHeaderSize {
		return ErrTruncatedRecord
	}

	hdr := (*RawRecordHeader)(unsafe.Pointer(&data[0]))
	r.RecordLength = hdr.RecordLength
	r.MajorVersion = hdr.MajorVersion
	r.MinorVersion = hdr.MinorVersion

	if r.RecordLength == 0 {
		return ErrInvalidRecordLength
	}

	switch hdr.MajorVersion {
	case 2:
		return r.unmarshal2(data)
	case 3:
		return r.unmarshal3(data)
	default:
		return fmt.Errorf("unsupported USN record version: %d.%d", hdr.MajorVersion, hdr.MinorVersion)
	}
}

// unmarshal2 assumes the header has already been umarshalled.
func (r *Record) unmarshal2(data []byte) error {
	if err := r.validateSize(data, recordV2Size); err != nil {
		return err
	}
	raw := (*RawRecordV2)(unsafe.Pointer(&data[0]))
	r.FileReferenceNumber = fileref.New64(raw.FileReferenceNumber)
	r.ParentFileReferenceNumber = fileref.New64(raw.ParentFileReferenceNumber)
	r.USN = raw.USN
	r.TimeStamp = time.Unix(0, raw.TimeStamp.Nanoseconds())
	r.Reason = raw.Reason
	r.SourceInfo = raw.SourceInfo
	r.FileAttributes = raw.FileAttributes
	return r.unmarshalFileName(data, raw.FileNameOffset, raw.FileNameLength)
}

// unmarshal3 assumes the header has already been umarshalled.
func (r *Record) unmarshal3(data []byte) error {
	if err := r.validateSize(data, recordV3Size); err != nil {
		return err
	}
	raw := (*RawRecordV3)(unsafe.Pointer(&data[0]))
	r.FileReferenceNumber = fileref.LittleEndian(raw.FileReferenceNumber)
	r.ParentFileReferenceNumber = fileref.LittleEndian(raw.ParentFileReferenceNumber)
	r.USN = raw.USN
	r.TimeStamp = time.Unix(0, raw.TimeStamp.Nanoseconds())
	r.Reason = raw.Reason
	r.SourceInfo = raw.SourceInfo
	r.FileAttributes = raw.FileAttributes
	return r.unmarshalFileName(data, raw.FileNameOffset, raw.FileNameLength)
}

// validateSize returns an error if the record length doesn't match the expected
// size. It assumes the header has already been umarshalled. If the record
// length is valid it returns nil.
func (r *Record) validateSize(data []byte, expected uint32) error {
	if uint32(len(data)) < expected {
		return ErrTruncatedRecord
	}
	if r.RecordLength < expected {
		return ErrInvalidRecordLength
	}
	return nil
}

// unmarshalFileName assumes the header has already been umarshalled.
func (r *Record) unmarshalFileName(data []byte, offset, length uint16) error {
	var (
		bufSize    = len(data)
		recordSize = int(r.RecordLength)
		start      = int(offset)
		end        = start + int(length)
	)
	if start > bufSize || end > bufSize {
		return ErrTruncatedRecord
	}
	if start > recordSize || end > recordSize {
		return ErrInvalidRecordLength
	}
	r.FileName = utf16BytesToString(data[start:end])
	return nil
}

// RawRecordHeader represents the raw form of the common USN journal record
// header.
type RawRecordHeader struct {
	RecordLength uint32
	MajorVersion uint16
	MinorVersion uint16
}

// RawRecordV1 represents the raw form of a version 1 change journal record.
type RawRecordV1 struct {
	RawRecordHeader
	FileReferenceNumber       int64
	ParentFileReferenceNumber int64
	USN                       USN
	TimeStamp                 windows.Filetime
}

// RawRecordV2 represents the raw form of a version 2 change journal record.
type RawRecordV2 struct {
	RawRecordHeader
	FileReferenceNumber       int64
	ParentFileReferenceNumber int64
	USN                       USN
	TimeStamp                 windows.Filetime
	Reason                    Reason
	SourceInfo                uint32
	SecurityID                uint32
	FileAttributes            fileattr.Value
	FileNameLength            uint16
	FileNameOffset            uint16
	_                         uint16
}

// RawRecordV3 represents the raw form of a version 3 change journal record.
type RawRecordV3 struct {
	RawRecordHeader
	FileReferenceNumber       [16]byte
	ParentFileReferenceNumber [16]byte
	USN                       USN
	TimeStamp                 windows.Filetime
	Reason                    Reason
	SourceInfo                uint32
	SecurityID                uint32
	FileAttributes            fileattr.Value
	FileNameLength            uint16
	FileNameOffset            uint16
	_                         uint16
}

// RawRecordV4 represents the raw form of a version 4 change journal record.
type RawRecordV4 struct {
	RawRecordHeader
	FileReferenceNumber       [16]byte
	ParentFileReferenceNumber [16]byte
	USN                       USN
	Reason                    Reason
	SourceInfo                uint32
	RemainingExtents          uint32
	NumberOfExtents           uint16
	ExtentSize                uint16
	_                         RecordExtent
}

// RecordExtent represents a change journal record extent.
type RecordExtent struct {
	Offset int64
	Length int64
}

//var RecordSet []Record

//func (rs *RecordSet) Read()
