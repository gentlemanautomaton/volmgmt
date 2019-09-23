package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

func parseSize(val string) int64 {
	if val == "" {
		return 0
	}

	s, err := humanize.ParseBytes(val)
	if err != nil {
		usage(fmt.Sprintf("Unable to parse file size \"%s\": %v", val, err))
	}

	return int64(s)
}
