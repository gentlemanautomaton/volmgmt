package fsctl

import (
	"github.com/gentlemanautomaton/volmgmt/ioctlcode"
	"github.com/gentlemanautomaton/volmgmt/ioctltype"
)

// I/O control codes for file systems.
var (
	RequestOplockLevel1    = ioctlcode.New(ioctltype.DeviceFileSystem, 0, ioctlcode.MethodBuffered, ioctlcode.AccessAny)        // FSCTL_REQUEST_OPLOCK_LEVEL_1
	RequestOplockLevel2    = ioctlcode.New(ioctltype.DeviceFileSystem, 1, ioctlcode.MethodBuffered, ioctlcode.AccessAny)        // FSCTL_REQUEST_OPLOCK_LEVEL_2
	RequestBatchOplock     = ioctlcode.New(ioctltype.DeviceFileSystem, 2, ioctlcode.MethodBuffered, ioctlcode.AccessAny)        // FSCTL_REQUEST_BATCH_OPLOCK
	OplockBreakAcknowledge = ioctlcode.New(ioctltype.DeviceFileSystem, 3, ioctlcode.MethodBuffered, ioctlcode.AccessAny)        // FSCTL_OPLOCK_BREAK_ACKNOWLEDGE
	OpbatchAckClosePending = ioctlcode.New(ioctltype.DeviceFileSystem, 4, ioctlcode.MethodBuffered, ioctlcode.AccessAny)        // FSCTL_OPBATCH_ACK_CLOSE_PENDING
	OplockBreakNotify      = ioctlcode.New(ioctltype.DeviceFileSystem, 5, ioctlcode.MethodBuffered, ioctlcode.AccessAny)        // FSCTL_OPLOCK_BREAK_NOTIFY
	LockVolume             = ioctlcode.New(ioctltype.DeviceFileSystem, 6, ioctlcode.MethodBuffered, ioctlcode.AccessAny)        // FSCTL_LOCK_VOLUME
	UnlockVolume           = ioctlcode.New(ioctltype.DeviceFileSystem, 7, ioctlcode.MethodBuffered, ioctlcode.AccessAny)        // FSCTL_UNLOCK_VOLUME
	DismountVolume         = ioctlcode.New(ioctltype.DeviceFileSystem, 8, ioctlcode.MethodBuffered, ioctlcode.AccessAny)        // FSCTL_DISMOUNT_VOLUME
	IsVolumeMounted        = ioctlcode.New(ioctltype.DeviceFileSystem, 10, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_IS_VOLUME_MOUNTED
	IsPathnameValid        = ioctlcode.New(ioctltype.DeviceFileSystem, 11, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_IS_PATHNAME_VALID
	MarkVolumeDirty        = ioctlcode.New(ioctltype.DeviceFileSystem, 12, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_MARK_VOLUME_DIRTY
	QueryRetrievalPointers = ioctlcode.New(ioctltype.DeviceFileSystem, 14, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_QUERY_RETRIEVAL_POINTERS
	GetCompression         = ioctlcode.New(ioctltype.DeviceFileSystem, 15, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_GET_COMPRESSION
	SetCompression         = ioctlcode.New(ioctltype.DeviceFileSystem, 16, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // FSCTL_SET_COMPRESSION
	SetBootloaderAccessed  = ioctlcode.New(ioctltype.DeviceFileSystem, 19, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_SET_BOOTLOADER_ACCESSED
	OplockBreakAckNo2      = ioctlcode.New(ioctltype.DeviceFileSystem, 20, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_OPLOCK_BREAK_ACK_NO_2
	InvalidateVolumes      = ioctlcode.New(ioctltype.DeviceFileSystem, 21, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_INVALIDATE_VOLUMES
	QueryFATBPB            = ioctlcode.New(ioctltype.DeviceFileSystem, 22, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_QUERY_FAT_BPB
	RequestFilterOplock    = ioctlcode.New(ioctltype.DeviceFileSystem, 23, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_REQUEST_FILTER_OPLOCK
	GetStatistics          = ioctlcode.New(ioctltype.DeviceFileSystem, 24, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_FILESYSTEM_GET_STATISTICS
)

// I/O control codes for file systems added in the Windows NT 4.0 release.
var (
	GetNTFSVolumeData    = ioctlcode.New(ioctltype.DeviceFileSystem, 25, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_GET_NTFS_VOLUME_DATA
	GetNTFSFileRecord    = ioctlcode.New(ioctltype.DeviceFileSystem, 26, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_GET_NTFS_FILE_RECORD
	GetVolumeBitmap      = ioctlcode.New(ioctltype.DeviceFileSystem, 27, ioctlcode.MethodNeither, ioctlcode.AccessAny)      // FSCTL_GET_VOLUME_BITMAP
	GetRetrievalPointers = ioctlcode.New(ioctltype.DeviceFileSystem, 28, ioctlcode.MethodNeither, ioctlcode.AccessAny)      // FSCTL_GET_RETRIEVAL_POINTERS
	MoveFile             = ioctlcode.New(ioctltype.DeviceFileSystem, 29, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial) // FSCTL_MOVE_FILE
	IsVolumeDirty        = ioctlcode.New(ioctltype.DeviceFileSystem, 30, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_IS_VOLUME_DIRTY
	AllowExtendedDASDIO  = ioctlcode.New(ioctltype.DeviceFileSystem, 32, ioctlcode.MethodNeither, ioctlcode.AccessAny)      // FSCTL_ALLOW_EXTENDED_DASD_IO
)

// I/O control codes for file systems added in the Windows 2000 release.
var (
	FindFilesBySID       = ioctlcode.New(ioctltype.DeviceFileSystem, 35, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_FIND_FILES_BY_SID
	SetObjectID          = ioctlcode.New(ioctltype.DeviceFileSystem, 38, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial)   // FSCTL_SET_OBJECT_ID
	GetObjectID          = ioctlcode.New(ioctltype.DeviceFileSystem, 39, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_GET_OBJECT_ID
	DeleteObjectID       = ioctlcode.New(ioctltype.DeviceFileSystem, 40, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial)   // FSCTL_DELETE_OBJECT_ID
	SetReparsePoint      = ioctlcode.New(ioctltype.DeviceFileSystem, 41, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial)   // FSCTL_SET_REPARSE_POINT
	GetReparsePoint      = ioctlcode.New(ioctltype.DeviceFileSystem, 42, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_GET_REPARSE_POINT
	DeleteReparsePoint   = ioctlcode.New(ioctltype.DeviceFileSystem, 43, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial)   // FSCTL_DELETE_REPARSE_POINT
	EnumUSNData          = ioctlcode.New(ioctltype.DeviceFileSystem, 44, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_ENUM_USN_DATA
	SecurityIDCheck      = ioctlcode.New(ioctltype.DeviceFileSystem, 45, ioctlcode.MethodNeither, ioctlcode.AccessRead)       // FSCTL_SECURITY_ID_CHECK
	ReadUSNJournal       = ioctlcode.New(ioctltype.DeviceFileSystem, 46, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_READ_USN_JOURNAL
	SetObjectIDExtended  = ioctlcode.New(ioctltype.DeviceFileSystem, 47, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial)   // FSCTL_SET_OBJECT_ID_EXTENDED
	CreateOrGetObjectID  = ioctlcode.New(ioctltype.DeviceFileSystem, 48, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_CREATE_OR_GET_OBJECT_ID
	SetSparse            = ioctlcode.New(ioctltype.DeviceFileSystem, 49, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial)   // FSCTL_SET_SPARSE
	SetZeroData          = ioctlcode.New(ioctltype.DeviceFileSystem, 50, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)     // FSCTL_SET_ZERO_DATA
	QueryAllocatedRanges = ioctlcode.New(ioctltype.DeviceFileSystem, 51, ioctlcode.MethodNeither, ioctlcode.AccessRead)       // FSCTL_QUERY_ALLOCATED_RANGES
	EnableUpgrade        = ioctlcode.New(ioctltype.DeviceFileSystem, 52, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)     // FSCTL_ENABLE_UPGRADE
	SetEncryption        = ioctlcode.New(ioctltype.DeviceFileSystem, 53, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_SET_ENCRYPTION
	EncryptionIO         = ioctlcode.New(ioctltype.DeviceFileSystem, 54, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_ENCRYPTION_FSCTL_IO
	WriteRawEncrypted    = ioctlcode.New(ioctltype.DeviceFileSystem, 55, ioctlcode.MethodNeither, ioctlcode.AccessSpecial)    // FSCTL_WRITE_RAW_ENCRYPTED
	ReadRawEncrypted     = ioctlcode.New(ioctltype.DeviceFileSystem, 56, ioctlcode.MethodNeither, ioctlcode.AccessSpecial)    // FSCTL_READ_RAW_ENCRYPTED
	CreateUSNJournal     = ioctlcode.New(ioctltype.DeviceFileSystem, 57, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_CREATE_USN_JOURNAL
	ReadFileUSNData      = ioctlcode.New(ioctltype.DeviceFileSystem, 58, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_READ_FILE_USN_DATA
	USNCloseRecord       = ioctlcode.New(ioctltype.DeviceFileSystem, 59, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_WRITE_USN_CLOSE_RECORD
	ExtendVolume         = ioctlcode.New(ioctltype.DeviceFileSystem, 60, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_EXTEND_VOLUME
	QueryUSNJournal      = ioctlcode.New(ioctltype.DeviceFileSystem, 61, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_QUERY_USN_JOURNAL
	DeleteUSNJournal     = ioctlcode.New(ioctltype.DeviceFileSystem, 62, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_DELETE_USN_JOURNAL
	MarkHandle           = ioctlcode.New(ioctltype.DeviceFileSystem, 63, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_MARK_HANDLE
	SISCopyfile          = ioctlcode.New(ioctltype.DeviceFileSystem, 64, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_SIS_COPYFILE
	SISLinkFiles         = ioctlcode.New(ioctltype.DeviceFileSystem, 65, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // FSCTL_SIS_LINK_FILES
	RecallFile           = ioctlcode.New(ioctltype.DeviceFileSystem, 69, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_RECALL_FILE
	ReadFromPlex         = ioctlcode.New(ioctltype.DeviceFileSystem, 71, ioctlcode.MethodOutDirect, ioctlcode.AccessRead)     // FSCTL_READ_FROM_PLEX
	FilePrefetch         = ioctlcode.New(ioctltype.DeviceFileSystem, 72, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial)   // FSCTL_FILE_PREFETCH
)

// I/O control codes for file systems added in the Windows Vista release.
var (
	MakeMediaCompatible            = ioctlcode.New(ioctltype.DeviceFileSystem, 76, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_MAKE_MEDIA_COMPATIBLE
	SetDefectManagement            = ioctlcode.New(ioctltype.DeviceFileSystem, 77, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_SET_DEFECT_MANAGEMENT
	QuerySparingInfo               = ioctlcode.New(ioctltype.DeviceFileSystem, 78, ioctlcode.MethodBuffered, ioctlcode.AccessAny)      // FSCTL_QUERY_SPARING_INFO
	QueryOnDiskVolumeInfo          = ioctlcode.New(ioctltype.DeviceFileSystem, 79, ioctlcode.MethodBuffered, ioctlcode.AccessAny)      // FSCTL_QUERY_ON_DISK_VOLUME_INFO
	SetVolumeCompressionState      = ioctlcode.New(ioctltype.DeviceFileSystem, 80, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial)  // FSCTL_SET_VOLUME_COMPRESSION_STATE
	TXFSModifyRM                   = ioctlcode.New(ioctltype.DeviceFileSystem, 81, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_TXFS_MODIFY_RM
	TXFSQueryRMInformation         = ioctlcode.New(ioctltype.DeviceFileSystem, 82, ioctlcode.MethodBuffered, ioctlcode.AccessRead)     // FSCTL_TXFS_QUERY_RM_INFORMATION
	TXFSRollforwardRedo            = ioctlcode.New(ioctltype.DeviceFileSystem, 84, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_TXFS_ROLLFORWARD_REDO
	TXFSRollforwardUndo            = ioctlcode.New(ioctltype.DeviceFileSystem, 85, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_TXFS_ROLLFORWARD_UNDO
	TXFSStartRM                    = ioctlcode.New(ioctltype.DeviceFileSystem, 86, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_TXFS_START_RM
	TXFSShutdownRM                 = ioctlcode.New(ioctltype.DeviceFileSystem, 87, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_TXFS_SHUTDOWN_RM
	TXFSReadBackupInformation      = ioctlcode.New(ioctltype.DeviceFileSystem, 88, ioctlcode.MethodBuffered, ioctlcode.AccessRead)     // FSCTL_TXFS_READ_BACKUP_INFORMATION
	TXFSWriteBackupInformation     = ioctlcode.New(ioctltype.DeviceFileSystem, 89, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_TXFS_WRITE_BACKUP_INFORMATION
	TXFSCreateSecondaryRM          = ioctlcode.New(ioctltype.DeviceFileSystem, 90, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_TXFS_CREATE_SECONDARY_RM
	TXFSGetMetadataInfo            = ioctlcode.New(ioctltype.DeviceFileSystem, 91, ioctlcode.MethodBuffered, ioctlcode.AccessRead)     // FSCTL_TXFS_GET_METADATA_INFO
	TXFSGetTransactedVersion       = ioctlcode.New(ioctltype.DeviceFileSystem, 92, ioctlcode.MethodBuffered, ioctlcode.AccessRead)     // FSCTL_TXFS_GET_TRANSACTED_VERSION
	TXFSSavepointInformation       = ioctlcode.New(ioctltype.DeviceFileSystem, 94, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_TXFS_SAVEPOINT_INFORMATION
	TXFSCreateMiniversion          = ioctlcode.New(ioctltype.DeviceFileSystem, 95, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)    // FSCTL_TXFS_CREATE_MINIVERSION
	TXFSTransactionActive          = ioctlcode.New(ioctltype.DeviceFileSystem, 99, ioctlcode.MethodBuffered, ioctlcode.AccessRead)     // FSCTL_TXFS_TRANSACTION_ACTIVE
	SetVolumeZeroOnDeallocation    = ioctlcode.New(ioctltype.DeviceFileSystem, 101, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial) // FSCTL_SET_ZERO_ON_DEALLOCATION
	SetRepair                      = ioctlcode.New(ioctltype.DeviceFileSystem, 102, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_SET_REPAIR
	GetRepair                      = ioctlcode.New(ioctltype.DeviceFileSystem, 103, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_GET_REPAIR
	WaitForRepair                  = ioctlcode.New(ioctltype.DeviceFileSystem, 104, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_WAIT_FOR_REPAIR
	InitiateRepair                 = ioctlcode.New(ioctltype.DeviceFileSystem, 106, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_INITIATE_REPAIR
	CSCInternal                    = ioctlcode.New(ioctltype.DeviceFileSystem, 107, ioctlcode.MethodNeither, ioctlcode.AccessAny)      // FSCTL_CSC_INTERNAL
	ShrinkVolume                   = ioctlcode.New(ioctltype.DeviceFileSystem, 108, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial) // FSCTL_SHRINK_VOLUME
	SetShortNameBehavior           = ioctlcode.New(ioctltype.DeviceFileSystem, 109, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_SET_SHORT_NAME_BEHAVIOR
	DFSRSetGhostHandleState        = ioctlcode.New(ioctltype.DeviceFileSystem, 110, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_DFSR_SET_GHOST_HANDLE_STATE
	TXFSListTransactionLockedFiles = ioctlcode.New(ioctltype.DeviceFileSystem, 120, ioctlcode.MethodBuffered, ioctlcode.AccessRead)    // FSCTL_TXFS_LIST_TRANSACTION_LOCKED_FILES
	TXFSListTransactions           = ioctlcode.New(ioctltype.DeviceFileSystem, 121, ioctlcode.MethodBuffered, ioctlcode.AccessRead)    // FSCTL_TXFS_LIST_TRANSACTIONS
	QueryPagefileEncryption        = ioctlcode.New(ioctltype.DeviceFileSystem, 122, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_QUERY_PAGEFILE_ENCRYPTION
	ResetVolumeAllocationHints     = ioctlcode.New(ioctltype.DeviceFileSystem, 123, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_RESET_VOLUME_ALLOCATION_HINTS
	TXFSReadBackupInformation2     = ioctlcode.New(ioctltype.DeviceFileSystem, 126, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_TXFS_READ_BACKUP_INFORMATION2
)

// I/O control codes for file systems added in the Windows 7 release.
var (
	QueryDependentVolume                = ioctlcode.New(ioctltype.DeviceFileSystem, 124, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_QUERY_DEPENDENT_VOLUME
	SDGlobalChange                      = ioctlcode.New(ioctltype.DeviceFileSystem, 125, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_SD_GLOBAL_CHANGE
	LookupStreamFromCluster             = ioctlcode.New(ioctltype.DeviceFileSystem, 127, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_LOOKUP_STREAM_FROM_CLUSTER
	TXFSWriteBackupInformation2         = ioctlcode.New(ioctltype.DeviceFileSystem, 128, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_TXFS_WRITE_BACKUP_INFORMATION2
	FileTypeNotification                = ioctlcode.New(ioctltype.DeviceFileSystem, 129, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_FILE_TYPE_NOTIFICATION
	BootAreaInfo                        = ioctlcode.New(ioctltype.DeviceFileSystem, 140, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_GET_BOOT_AREA_INFO
	GetRetrievalPointerBase             = ioctlcode.New(ioctltype.DeviceFileSystem, 141, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_GET_RETRIEVAL_POINTER_BASE
	SetPersistentVolumeState            = ioctlcode.New(ioctltype.DeviceFileSystem, 142, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_SET_PERSISTENT_VOLUME_STATE
	QueryPersistentVolumeState          = ioctlcode.New(ioctltype.DeviceFileSystem, 143, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_QUERY_PERSISTENT_VOLUME_STATE
	RequestOplock                       = ioctlcode.New(ioctltype.DeviceFileSystem, 144, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_REQUEST_OPLOCK
	CSVTunnelRequest                    = ioctlcode.New(ioctltype.DeviceFileSystem, 145, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_CSV_TUNNEL_REQUEST
	IsCSVFile                           = ioctlcode.New(ioctltype.DeviceFileSystem, 146, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_IS_CSV_FILE
	QueryFileSystemRecognition          = ioctlcode.New(ioctltype.DeviceFileSystem, 147, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_QUERY_FILE_SYSTEM_RECOGNITION
	GetVolumePathName                   = ioctlcode.New(ioctltype.DeviceFileSystem, 148, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_CSV_GET_VOLUME_PATH_NAME
	CSVGetVolumeNameForVolumeMountPoint = ioctlcode.New(ioctltype.DeviceFileSystem, 149, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_CSV_GET_VOLUME_NAME_FOR_VOLUME_MOUNT_POINT
	CSVGetVolumePathNamesForVolumeName  = ioctlcode.New(ioctltype.DeviceFileSystem, 150, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_CSV_GET_VOLUME_PATH_NAMES_FOR_VOLUME_NAME
	IsFileOnCSVVolume                   = ioctlcode.New(ioctltype.DeviceFileSystem, 151, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_IS_FILE_ON_CSV_VOLUME
	CSVInternal                         = ioctlcode.New(ioctltype.DeviceFileSystem, 155, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_CSV_INTERNAL
	SetExternalBacking                  = ioctlcode.New(ioctltype.DeviceFileSystem, 195, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial) // FSCTL_SET_EXTERNAL_BACKING
	GetExternalBacking                  = ioctlcode.New(ioctltype.DeviceFileSystem, 196, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_GET_EXTERNAL_BACKING
	DeleteExternalBacking               = ioctlcode.New(ioctltype.DeviceFileSystem, 197, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial) // FSCTL_DELETE_EXTERNAL_BACKING
	EnumExternalBacking                 = ioctlcode.New(ioctltype.DeviceFileSystem, 198, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_ENUM_EXTERNAL_BACKING
	EnumOverlay                         = ioctlcode.New(ioctltype.DeviceFileSystem, 199, ioctlcode.MethodNeither, ioctlcode.AccessAny)      // FSCTL_ENUM_OVERLAY
	AddOverlay                          = ioctlcode.New(ioctltype.DeviceFileSystem, 204, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)   // FSCTL_ADD_OVERLAY
	RemoveOverlay                       = ioctlcode.New(ioctltype.DeviceFileSystem, 205, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)   // FSCTL_REMOVE_OVERLAY
	UpdateOverlay                       = ioctlcode.New(ioctltype.DeviceFileSystem, 206, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)   // FSCTL_UPDATE_OVERLAY
	GetWOFVersion                       = ioctlcode.New(ioctltype.DeviceFileSystem, 218, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_GET_WOF_VERSION
	SuspendOverlay                      = ioctlcode.New(ioctltype.DeviceFileSystem, 225, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_SUSPEND_OVERLAY
)

// I/O control codes for file systems added in the Windows 8 release.
var (
	FileLevelTrim                              = ioctlcode.New(ioctltype.DeviceFileSystem, 130, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)     // FSCTL_FILE_LEVEL_TRIM
	CorruptionHandling                         = ioctlcode.New(ioctltype.DeviceFileSystem, 152, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_CORRUPTION_HANDLING
	OffloadRead                                = ioctlcode.New(ioctltype.DeviceFileSystem, 153, ioctlcode.MethodBuffered, ioctlcode.AccessRead)      // FSCTL_OFFLOAD_READ
	OffloadWrite                               = ioctlcode.New(ioctltype.DeviceFileSystem, 154, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)     // FSCTL_OFFLOAD_WRITE
	SetPurgeFailureMode                        = ioctlcode.New(ioctltype.DeviceFileSystem, 156, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_SET_PURGE_FAILURE_MODE
	QueryFileLayout                            = ioctlcode.New(ioctltype.DeviceFileSystem, 157, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_QUERY_FILE_LAYOUT
	IsVolumeOwnedByCSVFS                       = ioctlcode.New(ioctltype.DeviceFileSystem, 158, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_IS_VOLUME_OWNED_BYCSVFS
	GetIntegrityInformation                    = ioctlcode.New(ioctltype.DeviceFileSystem, 159, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_GET_INTEGRITY_INFORMATION
	SetIntegrityInformation                    = ioctlcode.New(ioctltype.DeviceFileSystem, 160, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // FSCTL_SET_INTEGRITY_INFORMATION
	QueryFileRegions                           = ioctlcode.New(ioctltype.DeviceFileSystem, 161, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_QUERY_FILE_REGIONS
	DedupFile                                  = ioctlcode.New(ioctltype.DeviceFileSystem, 165, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_DEDUP_FILE
	DedupQueryFileHashes                       = ioctlcode.New(ioctltype.DeviceFileSystem, 166, ioctlcode.MethodNeither, ioctlcode.AccessRead)       // FSCTL_DEDUP_QUERY_FILE_HASHES
	DedupQueryRangeState                       = ioctlcode.New(ioctltype.DeviceFileSystem, 167, ioctlcode.MethodNeither, ioctlcode.AccessRead)       // FSCTL_DEDUP_QUERY_RANGE_STATE
	DedupQueryReparseInfo                      = ioctlcode.New(ioctltype.DeviceFileSystem, 168, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_DEDUP_QUERY_REPARSE_INFO
	RKFInternal                                = ioctlcode.New(ioctltype.DeviceFileSystem, 171, ioctlcode.MethodNeither, ioctlcode.AccessAny)        // FSCTL_RKF_INTERNAL
	ScrubData                                  = ioctlcode.New(ioctltype.DeviceFileSystem, 172, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_SCRUB_DATA
	RepairCopies                               = ioctlcode.New(ioctltype.DeviceFileSystem, 173, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // FSCTL_REPAIR_COPIES
	DisableLocalBuffering                      = ioctlcode.New(ioctltype.DeviceFileSystem, 174, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_DISABLE_LOCAL_BUFFERING
	CSVMgmtLock                                = ioctlcode.New(ioctltype.DeviceFileSystem, 175, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_CSV_MGMT_LOCK
	CSVQueryDownLevelFileSystemCharacteristics = ioctlcode.New(ioctltype.DeviceFileSystem, 176, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_CSV_QUERY_DOWN_LEVEL_FILE_SYSTEM_CHARACTERISTICS
	AdvanceFileID                              = ioctlcode.New(ioctltype.DeviceFileSystem, 177, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_ADVANCE_FILE_ID
	CSVSyncTunnelRequest                       = ioctlcode.New(ioctltype.DeviceFileSystem, 178, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_CSV_SYNC_TUNNEL_REQUEST
	CSVQueryVetoFileDirectIO                   = ioctlcode.New(ioctltype.DeviceFileSystem, 179, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_CSV_QUERY_VETO_FILE_DIRECT_IO
	WriteUSNReason                             = ioctlcode.New(ioctltype.DeviceFileSystem, 180, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_WRITE_USN_REASON
	CSVControl                                 = ioctlcode.New(ioctltype.DeviceFileSystem, 181, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_CSV_CONTROL
	GetREFSVolumeData                          = ioctlcode.New(ioctltype.DeviceFileSystem, 182, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_GET_REFS_VOLUME_DATA
	CSVHBreakingSyncTunnelRequest              = ioctlcode.New(ioctltype.DeviceFileSystem, 185, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // FSCTL_CSV_H_BREAKING_SYNC_TUNNEL_REQUEST
)

// I/O control codes for file systems added in the Windows 8.1 release.
var (
	QueryStorageClasses           = ioctlcode.New(ioctltype.DeviceFileSystem, 187, ioctlcode.MethodBuffered, ioctlcode.AccessAny)    // FSCTL_QUERY_STORAGE_CLASSES
	QueryRegionInfo               = ioctlcode.New(ioctltype.DeviceFileSystem, 188, ioctlcode.MethodBuffered, ioctlcode.AccessAny)    // FSCTL_QUERY_REGION_INFO
	USNTrackModifiedRanges        = ioctlcode.New(ioctltype.DeviceFileSystem, 189, ioctlcode.MethodBuffered, ioctlcode.AccessAny)    // FSCTL_USN_TRACK_MODIFIED_RANGES
	QuerySharedVirtualDiskSupport = ioctlcode.New(ioctltype.DeviceFileSystem, 192, ioctlcode.MethodBuffered, ioctlcode.AccessAny)    // FSCTL_QUERY_SHARED_VIRTUAL_DISK_SUPPORT
	SVHDXSyncTunnelRequest        = ioctlcode.New(ioctltype.DeviceFileSystem, 193, ioctlcode.MethodBuffered, ioctlcode.AccessAny)    // FSCTL_SVHDX_SYNC_TUNNEL_REQUEST
	SVHDXSetInitiatorInformation  = ioctlcode.New(ioctltype.DeviceFileSystem, 194, ioctlcode.MethodBuffered, ioctlcode.AccessAny)    // FSCTL_SVHDX_SET_INITIATOR_INFORMATION
	DuplicateExtentsToFile        = ioctlcode.New(ioctltype.DeviceFileSystem, 209, ioctlcode.MethodBuffered, ioctlcode.AccessWrite)  // FSCTL_DUPLICATE_EXTENTS_TO_FILE
	SparseOverallocate            = ioctlcode.New(ioctltype.DeviceFileSystem, 211, ioctlcode.MethodNeither, ioctlcode.AccessSpecial) // FSCTL_SPARSE_OVERALLOCATE
	StorageQOSControl             = ioctlcode.New(ioctltype.DeviceFileSystem, 212, ioctlcode.MethodNeither, ioctlcode.AccessAny)     // FSCTL_STORAGE_QOS_CONTROL
	SVHDXAsyncTunnelRequest       = ioctlcode.New(ioctltype.DeviceFileSystem, 217, ioctlcode.MethodBuffered, ioctlcode.AccessAny)    // FSCTL_SVHDX_ASYNC_TUNNEL_REQUEST
)

// I/O control codes for file systems added in the Windows 10 release.
var (
	InitiateFileMetadataOptimization = ioctlcode.New(ioctltype.DeviceFileSystem, 215, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial) // FSCTL_INITIATE_FILE_METADATA_OPTIMIZATION
	QueryFileMetadataOptimization    = ioctlcode.New(ioctltype.DeviceFileSystem, 216, ioctlcode.MethodBuffered, ioctlcode.AccessSpecial) // FSCTL_QUERY_FILE_METADATA_OPTIMIZATION
	HCSSYncTunnelRequest             = ioctlcode.New(ioctltype.DeviceFileSystem, 219, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_HCS_SYNC_TUNNEL_REQUEST
	HCSAsyncTunnelRequest            = ioctlcode.New(ioctltype.DeviceFileSystem, 220, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_HCS_ASYNC_TUNNEL_REQUEST
	QueryExtentReadCacheInfo         = ioctlcode.New(ioctltype.DeviceFileSystem, 221, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_QUERY_EXTENT_READ_CACHE_INFO
	QueryREFSVolumeCounterInfo       = ioctlcode.New(ioctltype.DeviceFileSystem, 222, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_QUERY_REFS_VOLUME_COUNTER_INFO
	CleanVolumeMetadata              = ioctlcode.New(ioctltype.DeviceFileSystem, 223, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_CLEAN_VOLUME_METADATA
	SetIntegrityInformationEx        = ioctlcode.New(ioctltype.DeviceFileSystem, 224, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_SET_INTEGRITY_INFORMATION_EX
	VirtualStorageQueryProperty      = ioctlcode.New(ioctltype.DeviceFileSystem, 226, ioctlcode.MethodBuffered, ioctlcode.AccessAny)     // FSCTL_VIRTUAL_STORAGE_QUERY_PROPERTY
)
