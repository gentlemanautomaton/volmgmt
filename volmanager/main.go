package main

import (
	"flag"
	"fmt"

	"github.com/gentlemanautomaton/volmgmt/volume"
)

func main() {
	flag.Parse()

	for _, path := range flag.Args() {
		fmt.Printf("\"%s\" - ", path)
		vol, err := volume.New(path)
		if err != nil {
			fmt.Print(fmt.Errorf("Unable to create volume handle: %v", err))
			continue
		}
		defer vol.Close()
		name, err := vol.Name()
		if err != nil {
			fmt.Print(fmt.Errorf("Unable to ascertain volume name: %v", err))
			continue
		}
		paths, err := vol.Paths()
		if err != nil {
			fmt.Print(fmt.Errorf("Unable to ascertain volume paths: %v", err))
			continue
		}
		fmt.Printf("\"%s\"", name)
		if len(paths) > 0 {
			fmt.Printf("\n")
			for i, path := range paths {
				fmt.Printf("  Mount %d: %s\n", i, path)
			}
		}
	}
}
