package usn

import (
	"errors"
	"fmt"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	recordHeaderSize = 8
	recordV2Size     = 56
	recordV3Size     = 60
	recordV4Size     = 80
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
	Reason                    uint32
	SourceInfo                uint32
	FileName                  string
}

// UnmarshalBinary attempts to parse a single record from the given data.
func (r *Record) UnmarshalBinary(data []byte) error {
	bufSize := len(data)
	if bufSize < recordHeaderSize {
		return errors.New("insufficient data for USN record header")
	}

	hdr := (*RawRecordHeader)(unsafe.Pointer(&data[0]))
	r.RecordLength = hdr.RecordLength
	r.MajorVersion = hdr.MajorVersion
	r.MinorVersion = hdr.MinorVersion

	switch hdr.MajorVersion {
	case 2:
		if bufSize < recordV2Size {
			return errors.New("insufficient data for v2 USN record")
		}
		raw := (*RawRecordV2)(unsafe.Pointer(&data[0]))
		r.USN = raw.USN
		r.TimeStamp = time.Unix(0, raw.TimeStamp.Nanoseconds())
		r.Reason = raw.Reason
		r.SourceInfo = raw.SourceInfo
		start := int(raw.FileNameOffset)
		end := start + int(raw.FileNameLength)
		if end > bufSize {
			return errors.New("insufficient data for v2 USN record file name")
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
	Reason                    uint32
	SourceInfo                uint32
	SecurityID                uint32
	FileAttributes            uint32
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
	Reason                    uint32
	SourceInfo                uint32
	SecurityID                uint32
	FileAttributes            uint32
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
	Reason                    uint32
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
