package usnsource

import "strings"

// Info encodes information about the source of USN change journal records.
type Info uint32

// USN Journal Source Info Codes
const (
	Local                       Info = 0x00000000
	DataManagement              Info = 0x00000001 // USN_SOURCE_DATA_MANAGEMENT
	AuxilaryData                Info = 0x00000002 // USN_SOURCE_AUXILIARY_DATA
	ReplicationManagement       Info = 0x00000004 // USN_SOURCE_REPLICATION_MANAGEMENT
	ClientReplicationManagement Info = 0x00000008 // USN_SOURCE_CLIENT_REPLICATION_MANAGEMENT
)

// Match reports whether info contains all of the codes specified by c.
func (info Info) Match(c Info) bool {
	return info&c == c
}

// String returns a string representation of info.
func (info Info) String() string {
	return info.Join("|", FormatShort)
}

// Join returns a string representation of the source information using the
// given separator and format.
func (info Info) Join(sep string, format Format) string {
	if s, ok := format[info]; ok {
		return s
	}

	var matched []string
	for i := 0; i < 32; i++ {
		flag := Info(1 << uint32(i))
		if info.Match(flag) {
			if s, ok := format[flag]; ok {
				matched = append(matched, s)
			}
		}
	}

	return strings.Join(matched, sep)
}
