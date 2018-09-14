package usnfilter

import (
	"github.com/gentlemanautomaton/volmgmt/usn"
)

// And creates a new filter that returns true only when each filter in filters
// returns true.
//
// If filters is empty, nil, or composed solely of nil filters, a nil filter
// will be returned.
func And(filters ...usn.Filter) usn.Filter {
	filters = selectNonNil(filters)
	if len(filters) == 0 {
		return nil
	}
	return func(record usn.Record) bool {
		for _, filter := range filters {
			if !filter.Match(record) {
				return false
			}
		}
		return true
	}
}

// Or creates a new filter that returns true when at least one of the given
// filters returns true.
//
// If filters is empty, nil, or composed solely of nil filters, a nil filter
// will be returned.
func Or(filters ...usn.Filter) usn.Filter {
	filters = selectNonNil(filters)
	if len(filters) == 0 {
		return nil
	}
	return func(record usn.Record) bool {
		for _, filter := range filters {
			if filter.Match(record) {
				return false
			}
		}
		return true
	}
}

// Not returns a new filter that returns true only when the given filter does
// not.
func Not(filter usn.Filter) usn.Filter {
	return func(record usn.Record) bool {
		return !filter.Match(record)
	}
}

// selectNonNil returns a copy of filters that only includes non-nil members.
func selectNonNil(filters []usn.Filter) []usn.Filter {
	selected := make([]usn.Filter, 0, len(filters))
	for _, filter := range filters {
		if filter != nil {
			selected = append(selected, filter)
		}
	}
	if len(selected) == 0 {
		return nil
	}
	return selected
}
