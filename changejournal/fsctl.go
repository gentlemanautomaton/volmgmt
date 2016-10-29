package volmgmt

import "github.com/gentlemanautomaton/volmgmt/ioctl"

var (
	FSCTL_ENUM_USN_DATA             = ioctl.Code(ioctl.DeviceFileSystem, 44, ioctl.MethodNeither, ioctl.AccessAny)
	FSCTL_READ_USN_JOURNAL          = ioctl.Code(ioctl.DeviceFileSystem, 46, ioctl.MethodNeither, ioctl.AccessAny)
	FSCTL_CREATE_USN_JOURNAL        = ioctl.Code(ioctl.DeviceFileSystem, 57, ioctl.MethodNeither, ioctl.AccessAny)
	FSCTL_READ_FILE_USN_DATA        = ioctl.Code(ioctl.DeviceFileSystem, 58, ioctl.MethodNeither, ioctl.AccessAny)
	FSCTL_QUERY_USN_JOURNAL         = ioctl.Code(ioctl.DeviceFileSystem, 61, ioctl.MethodNeither, ioctl.AccessAny)
	FSCTL_DELETE_USN_JOURNAL        = ioctl.Code(ioctl.DeviceFileSystem, 62, ioctl.MethodNeither, ioctl.AccessAny)
	FSCTL_WRITE_USN_REASON          = ioctl.Code(ioctl.DeviceFileSystem, 180, ioctl.MethodNeither, ioctl.AccessAny)
	FSCTL_USN_TRACK_MODIFIED_RANGES = ioctl.Code(ioctl.DeviceFileSystem, 189, ioctl.MethodNeither, ioctl.AccessAny)
	/*
		FSCTL_DELETE_USN_JOURNAL
		FSCTL_ENUM_USN_DATA
		FSCTL_MARK_HANDLE
		FSCTL_QUERY_USN_JOURNAL
		FSCTL_READ_USN_JOURNAL
	*/
)
