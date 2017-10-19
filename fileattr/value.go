package fileattr

import "strings"

// Value describes a set of file attributes stored as part of a file's metadata.
type Value uint32

// File Attributes:
const (
	Readonly           Value = 0x000001 // FILE_ATTRIBUTE_READONLY
	Hidden             Value = 0x000002 // FILE_ATTRIBUTE_HIDDEN
	System             Value = 0x000004 // FILE_ATTRIBUTE_SYSTEM
	Directory          Value = 0x000010 // FILE_ATTRIBUTE_DIRECTORY
	Archive            Value = 0x000020 // FILE_ATTRIBUTE_ARCHIVE
	Device             Value = 0x000040 // FILE_ATTRIBUTE_DEVICE
	Normal             Value = 0x000080 // FILE_ATTRIBUTE_NORMAL
	Temporary          Value = 0x000100 // FILE_ATTRIBUTE_TEMPORARY
	SparseFile         Value = 0x000200 // FILE_ATTRIBUTE_SPARSE_FILE
	ReparsePoint       Value = 0x000400 // FILE_ATTRIBUTE_REPARSE_POINT
	Compressed         Value = 0x000800 // FILE_ATTRIBUTE_COMPRESSED
	Offline            Value = 0x001000 // FILE_ATTRIBUTE_OFFLINE
	NotContentIndexed  Value = 0x002000 // FILE_ATTRIBUTE_NOT_CONTENT_INDEXED
	Encrypted          Value = 0x004000 // FILE_ATTRIBUTE_ENCRYPTED
	IntegrityStream    Value = 0x008000 // FILE_ATTRIBUTE_INTEGRITY_STREAM
	Virtual            Value = 0x010000 // FILE_ATTRIBUTE_VIRTUAL
	NoScrubData        Value = 0x020000 // FILE_ATTRIBUTE_NO_SCRUB_DATA
	RecallOnOpen       Value = 0x040000 // FILE_ATTRIBUTE_RECALL_ON_OPEN
	RecallOnDataAccess Value = 0x400000 // FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS
)

// Match reports whether v contains all of the file attributes specified by
// c.
func (v Value) Match(c Value) bool {
	return v&c == c
}

// String returns a string representation of the file attributes using a default
// separator and format.
func (v Value) String() string {
	return v.Join("|", FormatGo)
}

// Join returns a string representation of the file attributes using the given
// separator and format.
func (v Value) Join(sep string, format Format) string {
	if s, ok := format[v]; ok {
		return s
	}

	var matched []string
	for i := 0; i < 32; i++ {
		flag := Value(1 << uint32(i))
		if v.Match(flag) {
			if s, ok := format[flag]; ok {
				matched = append(matched, s)
			}
		}
	}

	return strings.Join(matched, sep)
}
