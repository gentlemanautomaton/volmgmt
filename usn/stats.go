package usn

import "time"

// Stats hold statistics about a set of records.
type Stats struct {
	Bytes   uint64
	Records uint64
	First   time.Time
	Last    time.Time
}

// Add updates s to reflect the inclusion of r.
func (s *Stats) Add(r *Record) {
	s.Bytes += uint64(r.RecordLength)
	s.Records++
	if s.Records == 1 {
		s.First = r.TimeStamp
		s.Last = r.TimeStamp
	} else {
		if s.First.After(r.TimeStamp) {
			s.First = r.TimeStamp
		}
		if s.Last.Before(r.TimeStamp) {
			s.Last = r.TimeStamp
		}
	}
}
