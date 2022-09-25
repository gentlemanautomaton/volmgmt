package fileapi

import (
	"syscall"
	"unsafe"
)

// RenameInfo stores basic information about a file. It can be marsheled into
// the FILE_RENAME_INFO windows API structure.
type RenameInfo struct {
	ReplaceIfExists bool
	FileName        string
}

// Class returns the file information class.
func (info RenameInfo) Class() FileInfoClass {
	return FileRenameInfo
}

// MarshalBinary marshals info as binary data suitable for use with API calls.
func (info RenameInfo) MarshalBinary() (data []byte, err error) {
	utf16FileName, err := syscall.UTF16FromString(info.FileName)
	if err != nil {
		return nil, err
	}

	raw := rawRenameInfo{
		ReplaceIfExists: info.ReplaceIfExists,
		FileNameLength:  uint32((len(utf16FileName) - 1) * 2), // In bytes, excludes trailing null
	}

	// unsafe.Sizeof(raw) pads to 24, which is not what we want, so we
	// hardcode 20 bytes here
	const headerSize = 20

	data = make([]byte, headerSize+len(utf16FileName)*2)
	hdr := unsafe.Slice((*byte)(unsafe.Pointer(&raw)), headerSize)
	name := unsafe.Slice((*byte)(unsafe.Pointer(&utf16FileName[0])), len(utf16FileName)*2)
	copy(data, hdr)
	copy(data[headerSize:], name)

	return data, nil
}

// https://learn.microsoft.com/en-us/windows/win32/api/winbase/ns-winbase-file_rename_info
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fscc/1d2673a8-8fb9-4868-920a-775ccaa30cf8
type rawRenameInfo struct {
	ReplaceIfExists bool
	_               syscall.Handle
	FileNameLength  uint32 // In bytes, excludes trailing null
}
