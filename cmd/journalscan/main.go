package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/gentlemanautomaton/volmgmt/usn"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr, "%s\n\n", errmsg)
	flag.Usage()
	os.Exit(1)
}

func main() {
	var settings Settings
	{
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "usage: %s [-t type[,type...]] [-include regexp] [-exclude regexp] <volume>[,<volume>...]\n", os.Args[0])
			flag.PrintDefaults()
		}

		var (
			reasonStr  string
			reason     usn.Reason
			includeStr string
			include    *regexp.Regexp
			excludeStr string
			exclude    *regexp.Regexp
		)

		flag.StringVar(&reasonStr, "t", "*", "journal record types to include (comma-separated)")
		flag.StringVar(&includeStr, "include", "", "regular expression for file match (inclusion)")
		flag.StringVar(&excludeStr, "exclude", "", "regular expression for file match (exclusion)")
		flag.Parse()

		if flag.NArg() == 0 {
			usage("No volume specified.")
		}

		reason, err := usn.ParseReason(reasonStr)
		if err != nil {
			usage(fmt.Sprintf("%v", err))
		}

		include = compileRegex(includeStr)
		exclude = compileRegex(excludeStr)

		location, err := time.LoadLocation("Local")
		if err != nil {
			fmt.Printf("Unable to load local timezone information: %v\n", err)
			os.Exit(1)
		}

		settings = Settings{
			Reason:   reason,
			Include:  include,
			Exclude:  exclude,
			Location: location,
		}
	}

	paths := flag.Args()

	for _, path := range paths {
		scan(path, settings)
	}
}
