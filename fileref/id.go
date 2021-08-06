// Package fileref provides a unified representation of 64-bit and 128-bit
// file identifiers used in NTFS and ReFS file systems.
//
// New identifiers are created by calling New64, New128, BigEndian or
// LittleEndian.
package fileref

import (
	"math/big"
	"strconv"
)

// ID is a file identifer capable of holding 64-bit and 128-bit values.
//
// ID is stored as a big-endian byte sequence. NTFS and ReFS use a little-endian
// byte sequence, so the LittleEndian function should be used to create
// identifiers for byte sequences taken from those file systems.
type ID [16]byte

// New64 returns a 64-bit file identifier.
func New64(value int64) (id ID) {
	return New128(value, 0)
}

// New128 returns a 128-bit file identifier for a set of lower and upper
// 64 bit values.
func New128(lower int64, upper int64) (id ID) {
	id[0] = byte(upper >> 56)
	id[1] = byte(upper >> 48)
	id[2] = byte(upper >> 40)
	id[3] = byte(upper >> 32)
	id[4] = byte(upper >> 24)
	id[5] = byte(upper >> 16)
	id[6] = byte(upper >> 8)
	id[7] = byte(upper)
	id[8] = byte(lower >> 56)
	id[9] = byte(lower >> 48)
	id[10] = byte(lower >> 40)
	id[11] = byte(lower >> 32)
	id[12] = byte(lower >> 24)
	id[13] = byte(lower >> 16)
	id[14] = byte(lower >> 8)
	id[15] = byte(lower)
	return
}

// BigEndian creates a file identifier from a sequence of bytes in big-endian
// byte order.
func BigEndian(value [16]byte) ID {
	return ID(value)
}

// LittleEndian creates a file identifier from a sequence of bytes in
// little-endian byte order.
func LittleEndian(value [16]byte) ID {
	for i, j := 0, 15; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
	return ID(value)
}

// Int64 returns the file identifier as a 64-bit signed integer if it can be
// represented as one. If it cannot, -1 is returned.
func (id ID) Int64() int64 {
	upper, lower := id.Split()
	if upper != 0 {
		return -1
	}
	return lower
}

// IsInt64 returns true if the file identifier can be represented as a 64-bit
// signed integer.
func (id ID) IsInt64() bool {
	upper := int64(id[0])<<56 | int64(id[1])<<48 | int64(id[2])<<40 | int64(id[3])<<32 | int64(id[4])<<24 | int64(id[5])<<16 | int64(id[6])<<8 | int64(id[7])
	return upper == 0
}

// IsZero returns true if the file identifier is zero.
func (id ID) IsZero() bool {
	upper, lower := id.Split()
	return upper == 0 && lower == 0
}

// BigEndian returns the ID as a sequence of bytes in big-endian byte order.
func (id ID) BigEndian() (value [16]byte) {
	return [16]byte(id)
}

// LittleEndian returns the ID as a sequence of bytes in little-endian byte
// order.
func (id ID) LittleEndian() (value [16]byte) {
	for i := 0; i < 16; i++ {
		value[i] = id[15-i]
	}
	return
}

// Descriptor returns a descriptor for the file id.
func (id ID) Descriptor() Descriptor {
	if id.IsInt64() {
		return Descriptor{
			Size: 24,
			Type: FileType,
			Data: id.LittleEndian(),
		}
	}
	return Descriptor{
		Size: 24,
		Type: ExtendedFileIDType,
		Data: id.LittleEndian(),
	}
}

// String returns a string representation of the file identifier.
func (id ID) String() string {
	if id.IsInt64() {
		return strconv.FormatInt(id.Int64(), 10)
	}
	bi := big.NewInt(0)
	bi.SetBytes(id[:]) // This assumes that id is in big-endian byte order and that it is unsigned
	return bi.String()
}

// Split breaks the ID into upper and lower 64-bit values.
func (id ID) Split() (upper, lower int64) {
	upper = int64(id[0])<<56 | int64(id[1])<<48 | int64(id[2])<<40 | int64(id[3])<<32 | int64(id[4])<<24 | int64(id[5])<<16 | int64(id[6])<<8 | int64(id[7])
	lower = int64(id[8])<<56 | int64(id[9])<<48 | int64(id[10])<<40 | int64(id[11])<<32 | int64(id[12])<<24 | int64(id[13])<<16 | int64(id[14])<<8 | int64(id[15])
	return
}
