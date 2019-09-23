package main

import (
	"fmt"
	"regexp"
	"strings"
)

func compileRegex(re string) *regexp.Regexp {
	if re == "" {
		return nil
	}

	// Force case-insensitive matching
	if !strings.HasPrefix(re, "(?i)") {
		re = "(?i)" + re
	}

	c, err := regexp.Compile(re)
	if err != nil {
		usage(fmt.Sprintf("Unable to compile regular expression \"%s\": %v\n", re, err))
	}
	return c
}
