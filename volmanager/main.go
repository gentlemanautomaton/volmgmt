package main

import (
	"flag"
	"fmt"

	"github.com/gentlemanautomaton/volmgmt/volume"
)

func main() {
	flag.Parse()

	for _, path := range flag.Args() {
		fmt.Printf("Querying volume information for \"%s\"...\n--------\n", path)

		vol, err := volume.New(path)
		if err != nil {
			fmt.Print(fmt.Errorf("Unable to create volume handle: %v", err))
			continue
		}
		defer vol.Close()

		label, err := vol.Label()
		if err != nil {
			fmt.Print(fmt.Errorf("Unable to ascertain volume label: %v", err))
			continue
		}

		fmt.Printf("Volume Label: \"%s\"\n", label)
		fmt.Printf("Device Information: Number %d, Partition %d, Type %d\n", vol.DeviceNumber(), vol.PartitionNumber(), vol.DeviceType())
		fmt.Printf("Device Description: Removable: %t, Vendor: %s, Product: %s, Revision: %s, S/N: %s\n", vol.RemovableMedia(), vol.VendorID(), vol.ProductID(), vol.ProductRevision(), vol.SerialNumber())

		paths, err := vol.Paths()
		if err != nil {
			fmt.Printf("%v\n", fmt.Errorf("Unable to ascertain volume paths: %v", err))
			continue
		}
		if len(paths) > 0 {
			fmt.Printf("\n")
			for i, path := range paths {
				fmt.Printf("  Mount %d: %s\n", i, path)
			}
		}
	}
	fmt.Printf("--------\n")
}
