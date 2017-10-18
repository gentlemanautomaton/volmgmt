package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/gentlemanautomaton/volmgmt/usn"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr, "%s\n\n", errmsg)
	flag.Usage()
	os.Exit(2)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [-t type[,type...]] [-i regexp] [-e regexp] <volume>\n", os.Args[0])
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
	flag.StringVar(&includeStr, "i", "", "regular expression for file match (inclusion)")
	flag.StringVar(&excludeStr, "e", "", "regular expression for file match (exclusion)")
	flag.Parse()

	if flag.NArg() == 0 {
		usage("No volume specified.")
	}
	if flag.NArg() > 1 {
		usage("Only a single volume may be specified.")
	}

	path := flag.Arg(0)

	reason, err := usn.ParseReason(reasonStr)
	if err != nil {
		usage(fmt.Sprintf("%v", err))
	}

	include = compileRegex(includeStr)
	exclude = compileRegex(excludeStr)

	journal, err := usn.NewJournal(path)
	if err != nil {
		fmt.Printf("Unable to create monitor: %v\n", err)
		os.Exit(2)
	}
	defer journal.Close()

	data, err := journal.Query()
	if err != nil {
		fmt.Printf("Unable to access USN Journal: %v\n", err)
		os.Exit(2)
	}

	monitor := journal.Monitor()
	defer monitor.Close()

	feed := monitor.Listen(64) // Register the feed before starting the monitor

	err = monitor.Run(data.NextUSN, time.Millisecond*100, reason)
	if err != nil {
		fmt.Printf("Unable to monitor USN Journal: %v\n", err)
		os.Exit(2)
	}

	done := make(chan struct{})
	go run(feed, include, exclude, done)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ch:
	case <-done:
	}
}

func run(feed <-chan usn.Record, include, exclude *regexp.Regexp, done chan struct{}) {
	defer close(done)

	for record := range feed {
		if include != nil && !include.MatchString(record.FileName) {
			continue
		}

		if exclude != nil && exclude.MatchString(record.FileName) {
			continue
		}

		action := strings.ToUpper(record.Reason.Join("|", usn.ReasonFormatShort))

		fmt.Printf("%s  %s  %s\n", record.TimeStamp.Format("2006-01-02 15:04:05.000000"), action, record.FileName)
	}
}

func compileRegex(re string) *regexp.Regexp {
	if re == "" {
		return nil
	}

	c, err := regexp.Compile(re)
	if err != nil {
		usage(fmt.Sprintf("Unable to compile regular expression \"%s\": %v\n", re, err))
	}
	return c
}
