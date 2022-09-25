package fileapi

// FileInfoClass identifies the class of file information being accessed in
// file system API calls.
type FileInfoClass int

// File information classes.
const (
	FileBasicInfo = iota
	FileStandardInfo
	FileNameInfo
	FileRenameInfo
	FileDispositionInfo
	FileAllocationInfo
	FileEndOfFileInfo
	FileStreamInfo
	FileCompressionInfo
	FileAttributeTagInfo
	FileIDBothDirectoryInfo
	FileIDBothDirectoryRestartInfo
	FileIoPriorityHintInfo
	FileRemoteProtocolInfo
	FileFullDirectoryInfo
	FileFullDirectoryRestartInfo
	FileStorageInfo
	FileAlignmentInfo
	FileIDInfo
	FileIDExtdDirectoryInfo
	FileIDExtdDirectoryRestartInfo
	FileDispositionInfoEx
	FileRenameInfoEx
	FileCaseSensitiveInfo
	FileNormalizedNameInfo
)
