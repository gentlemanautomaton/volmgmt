package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/gentlemanautomaton/signaler"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr, "%s\n\n", errmsg)
	flag.Usage()
	os.Exit(1)
}

func main() {
	// Shutdown when we receive a termination signal
	shutdown := signaler.New().Capture(os.Interrupt, syscall.SIGTERM)

	// Ensure that we cleanup even if we panic
	defer shutdown.Trigger()

	// Prepare the update channels
	ctx := shutdown.Context()

	var settings Settings
	{
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "usage: %s [-include regexp] [-exclude regexp] [-bigger size] [-smaller size] <volume>[,<volume>...]\n", os.Args[0])
			flag.PrintDefaults()
		}

		var (
			includeStr     string
			include        *regexp.Regexp
			excludeStr     string
			exclude        *regexp.Regexp
			afterStr       string
			after          time.Time
			beforeStr      string
			before         time.Time
			biggerThanStr  string
			biggerThan     int64
			smallerThanStr string
			smallerThan    int64
			limit          int
			list           bool
			verbose        bool
			progress       bool
		)

		flag.StringVar(&includeStr, "include", "", "regular expression for file match (inclusion)")
		flag.StringVar(&excludeStr, "exclude", "", "regular expression for file match (exclusion)")
		flag.StringVar(&afterStr, "after", "", "only include entries at or after this time")
		flag.StringVar(&beforeStr, "before", "", "only include entries at or before this time")
		flag.StringVar(&biggerThanStr, "bigger", "", "only include entries bigger than this file size")
		flag.StringVar(&smallerThanStr, "smaller", "", "only include entries smaller than this file size")
		flag.BoolVar(&list, "list", false, "print matched file paths")
		flag.BoolVar(&verbose, "v", false, "print errors")
		flag.BoolVar(&progress, "p", false, "print progress messages")
		flag.IntVar(&limit, "limit", runtime.NumCPU(), "number of concurrent file operations to perform")
		flag.Parse()

		if flag.NArg() == 0 {
			usage("No volume specified.")
		}

		location, err := time.LoadLocation("Local")
		if err != nil {
			fmt.Printf("Unable to load local timezone information: %v\n", err)
			os.Exit(1)
		}

		include = compileRegex(includeStr)
		exclude = compileRegex(excludeStr)
		after = parseTime(afterStr, location)
		before = parseTime(beforeStr, location)
		biggerThan = parseSize(biggerThanStr)
		smallerThan = parseSize(smallerThanStr)

		settings = Settings{
			Include:     include,
			Exclude:     exclude,
			After:       after,
			Before:      before,
			Location:    location,
			BiggerThan:  biggerThan,
			SmallerThan: smallerThan,
			List:        list,
			Progress:    progress,
			Verbose:     verbose,
			Limit:       limit,
		}
	}

	paths := flag.Args()

	var summaries []Summary

	start := time.Now()
	for _, path := range paths {
		if ctx.Err() != nil {
			break
		}
		summary := scan(ctx, path, settings)
		summaries = append(summaries, summary)
	}
	end := time.Now()
	duration := end.Sub(start)

	if len(summaries) == 1 {
		fmt.Printf("%s\n", summaries[0])
		return
	}

	fmt.Printf("Total Time: %s.\n", duration)

	for i := range summaries {
		fmt.Printf("[%d] \"%s\" %s\n", i, paths[i], summaries[i])
	}

	fmt.Printf("[*] \"%s\" %s\n", strings.Join(paths, "|"), Combine(summaries...))
}
