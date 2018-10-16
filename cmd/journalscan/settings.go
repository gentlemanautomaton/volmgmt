package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gentlemanautomaton/volmgmt/usn"
)

// Settings hold various settings for a scan.
type Settings struct {
	Reason   usn.Reason
	Include  *regexp.Regexp
	Exclude  *regexp.Regexp
	After    time.Time
	Before   time.Time
	Location *time.Location
}

// Summary returns a multiline summary of the settings.
func (s Settings) Summary() string {
	var output []string
	if s.Include != nil {
		output = append(output, fmt.Sprintf("Include: %s", s.Include))
	}

	if s.Exclude != nil {
		output = append(output, fmt.Sprintf("Exclude: %s", s.Exclude))
	}

	if !s.After.IsZero() {
		output = append(output, fmt.Sprintf("After: %s", s.After))
	}

	if !s.Before.IsZero() {
		output = append(output, fmt.Sprintf("Before: %s", s.Before))
	}

	if len(output) == 0 {
		return ""
	}

	return strings.Join(output, "\n") + "\n"
}
