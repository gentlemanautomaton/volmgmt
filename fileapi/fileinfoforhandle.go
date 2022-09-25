package fileapi

import (
	"os"
	"syscall"
	"time"
)

// FileInfoForHandle implements the fs.FileInfo interface. It stores data in
// a form that is convenient when accessing files by their handle or file
// reference number.
type FileInfoForHandle struct {
	FileName string
	syscall.ByHandleFileInformation
}

// Name returns the name of the file.
func (fi FileInfoForHandle) Name() string {
	return fi.FileName
}

// Size returns the size of the file in bytes.
func (fi FileInfoForHandle) Size() int64 {
	return int64(fi.FileSizeHigh)<<32 + int64(fi.FileSizeLow)
}

// Mode returns the mode of the file.
//
// FIXME: This does not properly indicate whether a file is a symlink.
func (fi FileInfoForHandle) Mode() (m os.FileMode) {
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
func (fi FileInfoForHandle) ModTime() time.Time {
	return time.Unix(0, fi.LastWriteTime.Nanoseconds())
}

// IsDir returns true if the file is a directory.
func (fi FileInfoForHandle) IsDir() bool {
	return fi.Mode().IsDir()
}

// Sys returns implementation-specific details. It returns nil.
func (fi FileInfoForHandle) Sys() any {
	return nil
}
