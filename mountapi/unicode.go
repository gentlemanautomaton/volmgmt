package mountapi

import (
	"syscall"
	"unsafe"
)

func utf16BytesToString(s []byte) string {
	p := (*[0xffff]uint16)(unsafe.Pointer(&s[0]))
	return syscall.UTF16ToString(p[:len(s)/2])
}
