package usn

// USN represents a volume update sequence number.
type USN int64

const (
	// Min is the smallest valid update sequence number.
	Min = USN(0)

	// Max is the largest valid update sequence number.
	Max = USN(^uint64(0) >> 1)
)
