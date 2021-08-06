package usn

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/volmgmt/fsctl"
)

// See:       https://www.microsoft.com/msj/1099/journal2/journal2.aspx
// Archived: https://web.archive.org/web/20171018212725/https://www.microsoft.com/msj/1099/journal2/journal2.aspx

// CreateJournal will create a change journal on the file system volume
// represented by the provided handle and the parameters maximum journal size
// and allocation delta. If the parameters are zero, the journal will be created
// with defaults
func CreateJournal(handle syscall.Handle, maxSize, allocDelta uint64) (err error) {
	var length uint32
	var options = struct {
		MaximumSize     uint64
		AllocationDelta uint64
	}{maxSize, allocDelta}

	err = syscall.DeviceIoControl(handle, fsctl.CreateUSNJournal,
		(*byte)(unsafe.Pointer(&options)), uint32(unsafe.Sizeof(options)),
		nil, 0, &length, nil)
	return
}

// DeleteJournal will delete a change journal on the file system volume
// represented by the provided handle.
func DeleteJournal(handle syscall.Handle) (err error) {
	return errors.New("Not yet implemented")
}

// EnumData will enumerate change journal data on the file system volume
// represented by the provided handle.
func EnumData(handle syscall.Handle, opts RawEnumOptions, buffer []byte) (length uint32, err error) {
	var p1 *byte
	s1 := uint32(len(buffer))
	if s1 > 0 {
		p1 = &buffer[0]
	}
	err = syscall.DeviceIoControl(handle, fsctl.EnumUSNData, (*byte)(unsafe.Pointer(&opts)), uint32(unsafe.Sizeof(opts)), p1, s1, &length, nil)
	return
}

// MarkHandle will add file change information about a specified file and to
// the files metadata and to the USN change journal of the file system volume
// on which the file resides.
func MarkHandle(handle syscall.Handle) (err error) {
	return errors.New("Not yet implemented")
}

// QueryJournal retrieves information about the change journal affiliated with
// the given volume handle.
func QueryJournal(handle syscall.Handle) (data RawJournalData, err error) {
	var length uint32
	err = syscall.DeviceIoControl(handle, fsctl.QueryUSNJournal, nil, 0, (*byte)(unsafe.Pointer(&data)), uint32(unsafe.Sizeof(data)), &length, nil)
	return
}

// ReadJournal retrieves change journal data from the USN journal.
func ReadJournal(handle syscall.Handle, opts RawReadOptions, buffer []byte) (length uint32, err error) {
	var p1 *byte
	s1 := uint32(len(buffer))
	if s1 > 0 {
		p1 = &buffer[0]
	}
	err = syscall.DeviceIoControl(handle, fsctl.ReadUSNJournal, (*byte)(unsafe.Pointer(&opts)), uint32(unsafe.Sizeof(opts)), p1, s1, &length, nil)
	return
}
