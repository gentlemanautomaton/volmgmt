package usn

import (
	"context"
	"errors"
	"io"

	"github.com/gentlemanautomaton/volmgmt/fileref"
)

// ErrNotFound is returned by cache.Filer when a record cannot be found.
var ErrNotFound = errors.New("not found")

/*
// FileTable describes a mutable repository of file table records.
type FileTable interface {
	Get(usn USN) (Record, bool)
	Set(usn USN) Record
}
*/

// Calculate a reasonable buffer size for the cache to use when reading
// data.
//
// Records are a variable size because file names are a variable length.
// Here we use 32 UTF-16 characters as a rough estimate for the per-record
// file name length.
const fileNameSizeEstimate = 32 * 2
const cacheBufferSize = 8 + (recordV3Size+fileNameSizeEstimate)*2048 // USN + 2048 records

// Cache is a usn change journal cache.
type Cache struct {
	m      map[fileref.ID]Record
	buffer [cacheBufferSize]byte
}

// NewCache prepares a new cache object.
func NewCache() *Cache {
	return &Cache{
		m: make(map[fileref.ID]Record),
	}
}

// ReadFrom reads records from iter and inserts them into the record cache.
// It returns when the iterator returns an error or io.EOF, or if the given
// context is cancelled.
func (c *Cache) ReadFrom(ctx context.Context, iter Iter) error {
	var records []Record
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		var err error
		records, err = iter.Next(c.buffer[:], records[:0])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		for i := range records {
			frn := records[i].FileReferenceNumber
			c.m[frn] = records[i]
		}
	}
}

// Get returns the record for the given file reference number.
func (c *Cache) Get(frn fileref.ID) (record Record, ok bool) {
	record, ok = c.m[frn]
	return
}

// Set updates the record for the given file reference number.
func (c *Cache) Set(r Record) {
	c.m[r.FileReferenceNumber] = r
}

// Size returns the number of records in the cache
func (c *Cache) Size() int {
	return len(c.m)
}

// Filer is a Filer that uses the cache to retrieve values.
func (c *Cache) Filer(frn fileref.ID) (record Record, err error) {
	record, ok := c.m[frn]
	if !ok {
		err = ErrNotFound
	}
	return
}

// Records returns a slice of all records in the cache. The order of the
// returned records is unspecified.
func (c *Cache) Records() []Record {
	filer := Filer(c.Filer)
	records := make([]Record, 0, len(c.m))
	for _, record := range c.m {
		record.Path = record.FileName
		if !record.ParentFileReferenceNumber.IsZero() {
			parents, pErr := filer.Parents(record)
			if pErr == nil {
				for p := range parents {
					record.Path = parents[p].FileName + `\` + record.Path
				}
			}
		}
		records = append(records, record)
	}
	return records
}

/*
// FileTable describes a mutable repository of file table records.
type FileTable interface {
	Get(usn USN) (Record, bool)
	Set(usn USN) Record
}

// MFT manages a master file table. It uses a cache to
type MFT struct {
	cache FileTable
	h     hsync.Handle
}

// NewMFT returns a new master file table cache. It uses the provided
// repository to store data
func NewMFT(handle *hsync.Handle) {
}

// Init reads the entire master file table into the cache.
func (mft *MFT) Init() {

}

// Read returns the record matching the file reference number.
func (mft *MFT) Read(fileReferenceNumber uint64) Record {
	return Record{}
}

// ReadFrom reads records from Iter and inserts them into mft. It returns
// when the iterator returns an error or io.EOF.
//
// TODO: Add context?
func (mft *MFT) ReadFrom(iter Iter) error {
	return nil
}

// Update updates the record for the given file reference number.
func (mft *MFT) Update(r Record) {
}

*/
