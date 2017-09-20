package volumeapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procGetVolumeInformationByHandle     = modkernel32.NewProc("GetVolumeInformationByHandleW")
	procGetVolumeNameForVolumeMountPoint = modkernel32.NewProc("GetVolumeNameForVolumeMountPointW")
	procGetVolumePathNamesForVolumeName  = modkernel32.NewProc("GetVolumePathNamesForVolumeNameW")
)

const (
	volPathNamesBufSize = 512
)

// GetVolumeInformationByHandle retrieves information about the volume
// represented by the given system handle.
func GetVolumeInformationByHandle(handle syscall.Handle) (volumeName string, serialNumber uint32, maxComponentLength uint32, flags uint32, fileSystem string, err error) {
	var vnBuffer [MaxVolumeNameLength]uint16
	var fsnBuffer [MaxFileSystemNameLength]uint16

	p0 := &vnBuffer[0]
	p1 := &fsnBuffer[0]

	r0, _, e := syscall.Syscall9(
		procGetVolumeInformationByHandle.Addr(),
		8,
		uintptr(handle),
		uintptr(unsafe.Pointer(p0)),
		uintptr(MaxVolumeNameLength),
		uintptr(unsafe.Pointer(&serialNumber)),
		uintptr(unsafe.Pointer(&maxComponentLength)),
		uintptr(unsafe.Pointer(&flags)),
		uintptr(unsafe.Pointer(p1)),
		uintptr(MaxFileSystemNameLength),
		0)
	if r0 == 0 {
		if e != 0 {
			err = syscall.Errno(e)
		} else {
			err = syscall.EINVAL
		}
		return
	}
	volumeName = syscall.UTF16ToString(vnBuffer[:])
	fileSystem = syscall.UTF16ToString(fsnBuffer[:])
	return
}

// GetVolumeNameForVolumeMountPoint is a low level API wrapper for the syscall.
func GetVolumeNameForVolumeMountPoint(volumeMountPoint string) (volumeName string, err error) {
	if len(volumeMountPoint) == 0 {
		return "", syscall.EINVAL
	}

	vmpp, err := syscall.UTF16PtrFromString(volumeMountPoint)
	if err != nil {
		return "", err
	}

	var buffer [64]uint16
	p0 := &buffer[0]
	r0, _, e := syscall.Syscall(
		procGetVolumeNameForVolumeMountPoint.Addr(),
		3,
		uintptr(unsafe.Pointer(vmpp)),
		uintptr(unsafe.Pointer(p0)),
		uintptr(len(buffer)))
	if r0 == 0 {
		err = syscall.Errno(e)
	}

	volumeName = syscall.UTF16ToString(buffer[:])
	return
}

// GetVolumePathNamesForVolumeName returns a complete set of volume paths and
// mount points for the given volume name, which must be a volume GUID path in
// this form:
//
//   \\?\Volume{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}\
func GetVolumePathNamesForVolumeName(volumeName string) (pathNames []string, err error) {
	if len(volumeName) == 0 {
		return nil, syscall.EINVAL
	}

	vnp, err := syscall.UTF16PtrFromString(volumeName)
	if err != nil {
		return nil, err
	}

	var (
		length int
		b      [volPathNamesBufSize]uint16
		buffer = b[:]
	)

	// Make up to 3 attempts to get the list of path names.
	//
	// 1: Using a fixed buffer allocated on the stack
	// 2: Using a dynamic buffer based on the reported length (first attempt)
	// 3: Using a dynamic buffer based on the reported length (second attempt)
	//
	// It is extremely unlikely, but possible, that the length could change
	// between calls if a new path or mount is created.
L:
	for i := 0; i < 3; i++ {
		length, err = getVolumePathNamesForVolumeName(vnp, buffer)
		switch err {
		case nil:
			break L
		case ErrMoreData:
			buffer = make([]uint16, length)
		default:
			return
		}
	}
	if err == nil {
		pathNames = utf16ToSplitString(buffer[:length])
	}
	return
}

// getVolumePathNamesForVolumeName is a low level API wrapper for the syscall.
func getVolumePathNamesForVolumeName(volumeName *uint16, buffer []uint16) (length int, err error) {
	var p0 *uint16
	if len(buffer) > 0 {
		p0 = &buffer[0]
	}
	r0, _, e := syscall.Syscall6(
		procGetVolumePathNamesForVolumeName.Addr(),
		4,
		uintptr(unsafe.Pointer(volumeName)),
		uintptr(unsafe.Pointer(p0)),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&length)),
		0,
		0)
	if r0 == 0 {
		err = syscall.Errno(e)
	}
	return
}
