package fileattr

// Format describes a format for file attributes flags.
type Format map[Value]string

// FormatC maps values to C-style constant strings.
var FormatC = Format{
	Readonly:           "FILE_ATTRIBUTE_READONLY",
	Hidden:             "FILE_ATTRIBUTE_HIDDEN",
	System:             "FILE_ATTRIBUTE_SYSTEM",
	Directory:          "FILE_ATTRIBUTE_DIRECTORY",
	Archive:            "FILE_ATTRIBUTE_ARCHIVE",
	Device:             "FILE_ATTRIBUTE_DEVICE",
	Normal:             "FILE_ATTRIBUTE_NORMAL",
	Temporary:          "FILE_ATTRIBUTE_TEMPORARY",
	SparseFile:         "FILE_ATTRIBUTE_SPARSE_FILE",
	ReparsePoint:       "FILE_ATTRIBUTE_REPARSE_POINT",
	Compressed:         "FILE_ATTRIBUTE_COMPRESSED",
	Offline:            "FILE_ATTRIBUTE_OFFLINE",
	NotContentIndexed:  "FILE_ATTRIBUTE_NOT_CONTENT_INDEXED",
	Encrypted:          "FILE_ATTRIBUTE_ENCRYPTED",
	IntegrityStream:    "FILE_ATTRIBUTE_INTEGRITY_STREAM",
	Virtual:            "FILE_ATTRIBUTE_VIRTUAL",
	NoScrubData:        "FILE_ATTRIBUTE_NO_SCRUB_DATA",
	RecallOnOpen:       "FILE_ATTRIBUTE_RECALL_ON_OPEN",
	RecallOnDataAccess: "FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS",
}

// FormatGo maps values to Go-style constant strings.
var FormatGo = Format{
	Readonly:           "Readonly",
	Hidden:             "Hidden",
	System:             "System",
	Directory:          "Directory",
	Archive:            "Archive",
	Device:             "Device",
	Normal:             "Normal",
	Temporary:          "Temporary",
	SparseFile:         "SparseFile",
	ReparsePoint:       "ReparsePoint",
	Compressed:         "Compressed",
	Offline:            "Offline",
	NotContentIndexed:  "NotContentIndexed",
	Encrypted:          "Encrypted",
	IntegrityStream:    "IntegrityStream",
	Virtual:            "Virtual",
	NoScrubData:        "NoScrubData",
	RecallOnOpen:       "RecallOnOpen",
	RecallOnDataAccess: "RecallOnDataAccess",
}

// FormatCode maps values to single-letter codes.
//
// https://superuser.com/questions/44812/windows-explorers-file-attribute-column-values
var FormatCode = Format{
	Readonly:           "R",
	Hidden:             "H",
	System:             "S",
	Directory:          "D",
	Archive:            "A",
	Device:             "^", // Unofficial
	Normal:             "N",
	Temporary:          "T",
	SparseFile:         "P",
	ReparsePoint:       "L",
	Compressed:         "C",
	Offline:            "O",
	NotContentIndexed:  "I",
	Encrypted:          "E",
	IntegrityStream:    "V", // ReFS
	Virtual:            "-",
	NoScrubData:        "X", // ReFS
	RecallOnOpen:       "!", // Unofficial
	RecallOnDataAccess: "?", // Unofficial
}
