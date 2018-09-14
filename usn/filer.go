package usn

import (
	"github.com/gentlemanautomaton/volmgmt/fileref"
)

// Filer returns master file table records by looking up file identifiers.
type Filer func(id fileref.ID) (Record, error)

// Parents returns a slice of parent records for r, starting with the
// immediate parent.
func (f Filer) Parents(r Record) (records []Record, err error) {
	for r.FileReferenceNumber != r.ParentFileReferenceNumber && !r.ParentFileReferenceNumber.IsZero() {
		last := r.ParentFileReferenceNumber
		r, err = f(r.ParentFileReferenceNumber)
		if err != nil || r.ParentFileReferenceNumber == last {
			if err == ErrNotFound {
				err = nil
			}
			return
		}
		records = append(records, r)
	}
	return
}
