package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// Summary holds the summary of a scan.
type Summary struct {
	Skipped     int64
	Directories int64
	Files       int64
	TotalBytes  int64
	Sizes       []int64
}

// String returns a string representation of the summary.
func (s Summary) String() string {
	return fmt.Sprintf("Skipped: %d, Directories: %d, Files: %d (Total: %s, Mean: %s, Median: %s)",
		s.Skipped,
		s.Directories,
		s.Files,
		humanize.Bytes(uint64(s.TotalBytes)),
		humanize.Bytes(uint64(s.Mean())),
		humanize.Bytes(uint64(s.Median())))
}

// Mean returns the mean file size.
func (s Summary) Mean() int64 {
	if s.Files == 0 {
		return 0
	}
	return s.TotalBytes / s.Files
}

// Median returns the median file size.
func (s Summary) Median() int64 {
	length := len(s.Sizes)
	if length == 0 {
		return 0
	}
	middle := length / 2
	if length%2 == 0 {
		return (s.Sizes[middle-1] + s.Sizes[middle]) / 2
	}
	return s.Sizes[middle]
}

// Combine combines a series of summaries
func Combine(summaries ...Summary) Summary {
	var combined Summary
	sizeCount := 0
	for _, s := range summaries {
		sizeCount += len(s.Sizes)
	}
	combined.Sizes = make([]int64, 0, sizeCount)
	for _, s := range summaries {
		combined.Skipped += s.Skipped
		combined.Directories += s.Directories
		combined.Files += s.Files
		combined.TotalBytes += s.TotalBytes
		combined.Sizes = append(combined.Sizes, s.Sizes...)
	}
	return combined
}
