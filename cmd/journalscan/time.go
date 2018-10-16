package main

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

func parseTime(val string, loc *time.Location) time.Time {
	if val == "" {
		return time.Time{}
	}

	t, err := dateparse.ParseIn(val, loc)
	if err != nil {
		usage(fmt.Sprintf("Unable to parse time \"%s\": %v", val, err))
	}

	return t
}
