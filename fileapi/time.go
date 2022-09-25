package fileapi

import (
	"syscall"
	"time"
)

func timeToFiletime(t time.Time) syscall.Filetime {
	if t.IsZero() {
		return syscall.Filetime{}
	}
	return syscall.NsecToFiletime(t.UnixNano())
}

func filetimeToTime(ft syscall.Filetime) time.Time {
	if ft.LowDateTime == 0 && ft.HighDateTime == 0 {
		return time.Time{}
	}
	return time.Unix(0, ft.Nanoseconds())
}
