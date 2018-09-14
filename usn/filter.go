package usn

// Filter is capable of filtering USN change journal records. It returns true
// when a record should be included in a set and false when it should not.
type Filter func(Record) bool

// Match returns f(r) if f is non-nil. If f is nil it returns true.
func (f Filter) Match(r Record) bool {
	if f == nil {
		return true
	}
	return f(r)
}
