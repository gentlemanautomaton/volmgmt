package usn

import (
	"errors"
	"fmt"
	"time"
	"unsafe"

	"github.com/gentlemanautomaton/volmgmt/fileattr"

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
	FileReferenceNumber       uint64
	ParentFileReferenceNumber uint64
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
		if bufSize < recordV2Size {
			return ErrTruncatedRecord
		}
		if r.RecordLength < recordV2Size {
			return ErrInvalidRecordLength
		}
		raw := (*RawRecordV2)(unsafe.Pointer(&data[0]))
		r.USN = raw.USN
		r.TimeStamp = time.Unix(0, raw.TimeStamp.Nanoseconds())
		r.Reason = raw.Reason
		r.SourceInfo = raw.SourceInfo
		r.FileAttributes = raw.FileAttributes
		start := int(raw.FileNameOffset)
		end := start + int(raw.FileNameLength)
		if end > bufSize {
			return ErrTruncatedRecord
		}
		if end > int(r.RecordLength) {
			return ErrInvalidRecordLength
		}
		r.FileName = utf16BytesToString(data[start:end])
	default:
		return fmt.Errorf("unsupported USN record version: %d.%d", hdr.MajorVersion, hdr.MinorVersion)
	}
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
	FileReferenceNumber       uint64
	ParentFileReferenceNumber uint64
	USN                       USN
	TimeStamp                 windows.Filetime
}

// RawRecordV2 represents the raw form of a version 2 change journal record.
type RawRecordV2 struct {
	RawRecordHeader
	FileReferenceNumber       uint64
	ParentFileReferenceNumber uint64
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
	FileReferenceNumber       FileID128
	ParentFileReferenceNumber FileID128
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
	FileReferenceNumber       FileID128
	ParentFileReferenceNumber FileID128
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
