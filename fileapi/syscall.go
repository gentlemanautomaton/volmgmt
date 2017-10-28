package fileapi

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/volmgmt/fileref"
	"golang.org/x/sys/windows"
)

var (
	// ErrEmptyBuffer is returned when a nil or zero-sized buffer is provided
	// to a system call.
	ErrEmptyBuffer = errors.New("nil or empty buffer provided")
)

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procOpenFileByID                 = modkernel32.NewProc("OpenFileById")
	procGetFileInformationByHandleEx = modkernel32.NewProc("GetFileInformationByHandleEx")
)

// OpenFileByID opens a file by its file ID. The file will be opened with the
// given access, share mode and flags.
//
// The handle provided can be to any file or on the volume, or to the volume
// itself.
func OpenFileByID(peer syscall.Handle, id fileref.ID, access, shareMode, flags uint32) (handle syscall.Handle, err error) {
	d := id.Descriptor()

	r0, _, e := syscall.Syscall6(
		procOpenFileByID.Addr(),
		6,
		uintptr(peer),
		uintptr(unsafe.Pointer(&d)),
		uintptr(access),
		uintptr(shareMode),
		0,
		uintptr(flags))
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		if e != 0 {
			err = syscall.Errno(e)
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

// GetFileInformationByHandleEx retrieves information about the file
// represented by the given system handle. The type of information returned
// is determined by class.
func GetFileInformationByHandleEx(handle syscall.Handle, class uint32, buffer []byte) (err error) {
	if len(buffer) == 0 {
		return ErrEmptyBuffer
	}

	r0, _, e := syscall.Syscall6(
		procGetFileInformationByHandleEx.Addr(),
		4,
		uintptr(handle),
		uintptr(class),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
		0,
		0)
	if r0 == 0 {
		if e != 0 {
			err = syscall.Errno(e)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
