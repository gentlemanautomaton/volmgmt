package main

import (
	"os"

	"github.com/gentlemanautomaton/volmgmt/usn"
)

// FileInfoFilter is a filter based on a file's attributes.
type FileInfoFilter func(os.FileInfo) bool

func buildRecordFilter(settings Settings) usn.Filter {
	return func(record usn.Record) bool {
		if settings.Include != nil {
			if record.Path == "" {
				if !settings.Include.MatchString(record.FileName) {
					return false
				}
			} else {
				if !settings.Include.MatchString(record.Path) {
					return false
				}
			}
		}

		if settings.Exclude != nil {
			if record.Path == "" {
				if settings.Exclude.MatchString(record.FileName) {
					return false
				}
			} else {
				if settings.Exclude.MatchString(record.Path) {
					return false
				}
			}
		}

		return true
	}
}

func buildFileInfoFilter(settings Settings) FileInfoFilter {
	return func(fi os.FileInfo) bool {
		if settings.BiggerThan > 0 {
			if fi.Size() <= settings.BiggerThan {
				return false
			}
		}

		if settings.SmallerThan > 0 {
			if fi.Size() >= settings.SmallerThan {
				return false
			}
		}

		if !settings.After.IsZero() {
			if fi.ModTime().Before(settings.After) {
				return false
			}
		}

		if !settings.Before.IsZero() {
			if fi.ModTime().After(settings.Before) {
				return false
			}
		}

		return true
	}
}
