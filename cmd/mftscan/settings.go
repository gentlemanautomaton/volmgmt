package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

// Settings hold various settings for a scan.
type Settings struct {
	Include     *regexp.Regexp
	Exclude     *regexp.Regexp
	After       time.Time
	Before      time.Time
	Location    *time.Location
	BiggerThan  int64 // In bytes
	SmallerThan int64 // In bytes
	List        bool
	Progress    bool
	Verbose     bool
	Limit       int
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

	if s.BiggerThan > 0 {
		output = append(output, fmt.Sprintf("Bigger Than: %s", humanize.Bytes(uint64(s.BiggerThan))))
	}

	if s.SmallerThan > 0 {
		output = append(output, fmt.Sprintf("Smaller Than: %s", humanize.Bytes(uint64(s.SmallerThan))))
	}

	if s.List {
		output = append(output, "List Files: On")
	}

	if s.Progress {
		output = append(output, "Progress: On")
	}

	if s.Verbose {
		output = append(output, "Verbose: On")
	}

	if s.Limit != 1 {
		output = append(output, fmt.Sprintf("Concurrent Reads: %d", s.Limit))
	}

	if len(output) == 0 {
		return ""
	}

	return strings.Join(output, "\n") + "\n"
}
