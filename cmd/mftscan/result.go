package main

// Result is the assessment of a file.
type Result int

// File assessment results.
const (
	Skipped      Result = 0
	Inaccessible Result = 1
	Dir          Result = 2
	File         Result = 3
	Reparse      Result = 4
)
