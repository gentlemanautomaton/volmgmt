package fileapi

import (
	"errors"
	"syscall"
	"time"
	"unsafe"

	"github.com/gentlemanautomaton/volmgmt/fileattr"
)

// BasicInfo stores basic information about a file. It can be marsheled into
// the FILE_BASIC_INFO windows API structure.
type BasicInfo struct {
	CreationTime   time.Time
	LastAccessTime time.Time
	LastWriteTime  time.Time
	ChangeTime     time.Time
	FileAttributes fileattr.Value
}

// Class returns the file information class.
func (info BasicInfo) Class() FileInfoClass {
	return FileBasicInfo
}

// Bytes marshals info as binary data suitable for use with API calls.
func (info BasicInfo) Bytes() (data []byte) {
	raw := rawBasicInfo{
		CreationTime:   timeToFiletime(info.CreationTime),
		LastAccessTime: timeToFiletime(info.LastAccessTime),
		LastWriteTime:  timeToFiletime(info.LastWriteTime),
		ChangeTime:     timeToFiletime(info.ChangeTime),
		FileAttributes: info.FileAttributes,
	}
	b := unsafe.Slice((*byte)(unsafe.Pointer(&raw)), unsafe.Sizeof(raw))
	return b
}

// MarshalBinary marshals info as binary data suitable for use with API calls.
func (info BasicInfo) MarshalBinary() (data []byte, err error) {
	return info.Bytes(), nil
}

// Size returns the the number of bytes required to store BasicInfo in its
// marshaled form.
func (info *BasicInfo) Size() int {
	return int(unsafe.Sizeof(rawBasicInfo{}))
}

// UnmarshalBinary unmarshals the given data into info.
func (info *BasicInfo) UnmarshalBinary(data []byte) error {
	if len(data) < info.Size() {
		return errors.New("insufficient data for BasicInfo unmarshaling")
	}

	raw := (*rawBasicInfo)(unsafe.Pointer(&data[0]))

	info.CreationTime = filetimeToTime(raw.CreationTime)
	info.LastAccessTime = filetimeToTime(raw.LastAccessTime)
	info.LastWriteTime = filetimeToTime(raw.LastWriteTime)
	info.ChangeTime = filetimeToTime(raw.ChangeTime)
	info.FileAttributes = raw.FileAttributes

	return nil
}

// https://learn.microsoft.com/en-us/windows/win32/api/winbase/ns-winbase-file_basic_info
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fscc/16023025-8a78-492f-8b96-c873b042ac50
type rawBasicInfo struct {
	CreationTime   syscall.Filetime
	LastAccessTime syscall.Filetime
	LastWriteTime  syscall.Filetime
	ChangeTime     syscall.Filetime
	FileAttributes fileattr.Value
	reserved       uint32
}
