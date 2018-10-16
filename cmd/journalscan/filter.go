package main

import "github.com/gentlemanautomaton/volmgmt/usn"

func buildFilter(settings Settings) usn.Filter {
	return func(record usn.Record) bool {
		if !settings.After.IsZero() {
			if record.TimeStamp.Before(settings.After) {
				return false
			}
		}

		if !settings.Before.IsZero() {
			if record.TimeStamp.After(settings.Before) {
				return false
			}
		}

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
