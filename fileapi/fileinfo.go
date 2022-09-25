package fileapi

// FileInfo is a type that holds file information data.
type FileInfo interface {
	Class() FileInfoClass
}

// FileInfoMarshaler is a file information data type that can be marshaled
// into a form acceptable to the file system APIs.
type FileInfoMarshaler interface {
	FileInfo
	MarshalBinary() (data []byte, err error)
}

// FileInfoMarshaler is a file information data type that can be unmarshaled
// from file system APIs.
type FileInfoUnmarshaler interface {
	FileInfo
	Size() int
	UnmarshalBinary(data []byte) error
}
