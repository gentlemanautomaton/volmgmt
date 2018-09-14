package usnfilter

import (
	"regexp"

	"github.com/gentlemanautomaton/volmgmt/usn"
)

// PathRegexp returns a filter that returns true when records have a filename
// or path matching the given regular expression.
func PathRegexp(re *regexp.Regexp) usn.Filter {
	return func(record usn.Record) bool {
		if record.Path == "" {
			return re.MatchString(record.FileName)
		}
		return re.MatchString(record.Path)
	}
}
