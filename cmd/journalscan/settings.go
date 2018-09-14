package main

import (
	"regexp"
	"time"

	"github.com/gentlemanautomaton/volmgmt/usn"
)

// Settings hold various settings for a scan.
type Settings struct {
	Reason   usn.Reason
	Include  *regexp.Regexp
	Exclude  *regexp.Regexp
	Location *time.Location
}
