package usnfilter

import (
	"strings"

	"github.com/gentlemanautomaton/volmgmt/usn"
)

// PathContains returns a filter that returns true when a record's path contains
// the given substring. If the record does not specify a path its filename will
// be used instead. The match is case-insensitive.
func PathContains(substr string) usn.Filter {
	ci := strings.ToLower(substr)
	return func(record usn.Record) bool {
		if record.Path == "" {
			return strings.Contains(strings.ToLower(record.FileName), ci)
		}
		return strings.Contains(strings.ToLower(record.Path), ci)
	}
}

// PathPrefix returns a filter that returns true when a record's path starts
// with the given prefix. If the record does not specify a path its filename
// will be used instead. The match is case-insensitive.
func PathPrefix(prefix string) usn.Filter {
	ci := strings.ToLower(prefix)
	return func(record usn.Record) bool {
		if record.Path == "" {
			return strings.HasPrefix(strings.ToLower(record.FileName), ci)
		}
		return strings.HasPrefix(strings.ToLower(record.Path), ci)
	}
}
