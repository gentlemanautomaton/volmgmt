package fileapi

import (
	"os"
	"syscall"
	"time"
)

// FileInfo implements the fs.FileInfo interface. It stores data in a form
// that is convenient when accessing files by their handle or file reference
// number.
type FileInfo struct {
	FileName string
	syscall.ByHandleFileInformation
}

// Name returns the name of the file.
func (fi FileInfo) Name() string {
	return fi.FileName
}

// Size returns the size of the file in bytes.
func (fi FileInfo) Size() int64 {
	return int64(fi.FileSizeHigh)<<32 + int64(fi.FileSizeLow)
}

// Mode returns the mode of the file.
//
// FIXME: This does not properly indicate whether a file is a symlink.
func (fi FileInfo) Mode() (m os.FileMode) {
	if fi.FileAttributes&syscall.FILE_ATTRIBUTE_READONLY != 0 {
		m |= 0444
	} else {
		m |= 0666
	}
	if fi.FileAttributes&syscall.FILE_ATTRIBUTE_DIRECTORY != 0 {
		m |= os.ModeDir | 0111
	}
	return m
}

// ModTime returns the last modification time of the file.
func (fi FileInfo) ModTime() time.Time {
	return time.Unix(0, fi.LastWriteTime.Nanoseconds())
}

// IsDir returns true if the file is a directory.
func (fi FileInfo) IsDir() bool {
	return fi.Mode().IsDir()
}

// Sys returns implementation-specific details. It returns nil.
func (fi FileInfo) Sys() any {
	return nil
}
