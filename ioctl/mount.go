package ioctl

import (
	"github.com/gentlemanautomaton/volmgmt/ioctlcode"
	"github.com/gentlemanautomaton/volmgmt/ioctltype"
)

// I/O control codes for interaction with the mount manager.
var (
	MountMgrCreatePoint               = ioctlcode.New(ioctltype.MountMgr, 0, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTMGR_CREATE_POINT
	MountMgrDeletePoints              = ioctlcode.New(ioctltype.MountMgr, 1, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTMGR_DELETE_POINTS
	MountMgrQueryPoints               = ioctlcode.New(ioctltype.MountMgr, 2, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_MOUNTMGR_QUERY_POINTS
	MountMgrDeletePointsDBOnly        = ioctlcode.New(ioctltype.MountMgr, 3, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTMGR_DELETE_POINTS_DBONLY
	MountMgrNextDriveLetter           = ioctlcode.New(ioctltype.MountMgr, 4, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTMGR_NEXT_DRIVE_LETTER
	MountMgrAutoDLAssignments         = ioctlcode.New(ioctltype.MountMgr, 5, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTMGR_AUTO_DL_ASSIGNMENTS
	MountMgrVolumeMountPointCreated   = ioctlcode.New(ioctltype.MountMgr, 6, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTMGR_VOLUME_MOUNT_POINT_CREATED
	MountMgrVolumeMountPointDeleted   = ioctlcode.New(ioctltype.MountMgr, 7, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTMGR_VOLUME_MOUNT_POINT_DELETED
	MountMgrChangeNotify              = ioctlcode.New(ioctltype.MountMgr, 8, ioctlcode.MethodBuffered, ioctlcode.AccessRead)      // IOCTL_MOUNTMGR_CHANGE_NOTIFY
	MountMgrKeepLinksWhenOffline      = ioctlcode.New(ioctltype.MountMgr, 9, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTMGR_KEEP_LINKS_WHEN_OFFLINE
	MountMgrCheckUnprocessedVolumes   = ioctlcode.New(ioctltype.MountMgr, 10, ioctlcode.MethodBuffered, ioctlcode.AccessRead)     // IOCTL_MOUNTMGR_CHECK_UNPROCESSED_VOLUMES
	MountMgrVolumeArrivalNotification = ioctlcode.New(ioctltype.MountMgr, 11, ioctlcode.MethodBuffered, ioctlcode.AccessRead)     // IOCTL_MOUNTMGR_VOLUME_ARRIVAL_NOTIFICATION
)

// I/O control codes for interaction with mounted devices.
var (
	MountDevQueryUniqueID          = ioctlcode.New(ioctltype.MountDev, 0, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_MOUNTDEV_QUERY_UNIQUE_ID
	MountDevUniqueIDChangeNotify   = ioctlcode.New(ioctltype.MountDev, 1, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTDEV_UNIQUE_ID_CHANGE_NOTIFY
	MountDevQueryDeviceName        = ioctlcode.New(ioctltype.MountDev, 2, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_MOUNTDEV_QUERY_DEVICE_NAME
	MountDevQuerySuggestedLinkName = ioctlcode.New(ioctltype.MountDev, 3, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_MOUNTDEV_QUERY_SUGGESTED_LINK_NAME
	MountDevLinkCreated            = ioctlcode.New(ioctltype.MountDev, 4, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTDEV_LINK_CREATED
	MountDevLinkDeleted            = ioctlcode.New(ioctltype.MountDev, 5, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_MOUNTDEV_LINK_DELETED
	MountDevQueryStableGUID        = ioctlcode.New(ioctltype.MountDev, 6, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_MOUNTDEV_QUERY_STABLE_GUID
)
