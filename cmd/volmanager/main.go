package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/gentlemanautomaton/volmgmt/usn"
	"github.com/gentlemanautomaton/volmgmt/volume"
	"github.com/gentlemanautomaton/volmgmt/volumeapi"
)

func main() {
	var (
		regexString string
		regex       *regexp.Regexp
	)
	flag.StringVar(&regexString, "match", "", "regular expression for file match")
	flag.Parse()

	if regexString != "" {
		var regexErr error
		regex, regexErr = regexp.Compile(regexString)
		if regexErr != nil {
			fmt.Printf("Unable to compile matching expression \"%s\": %v\n", regexString, regexErr)
			os.Exit(2)
		}
	}

	for _, path := range flag.Args() {
		fmt.Printf("Querying volume information for \"%s\"...\n--------\n", path)

		vol, err := volume.New(path)
		if err != nil {
			fmt.Printf("Unable to create volume handle: %v\n", err)
			continue
		}
		defer vol.Close()

		label, labelErr := vol.Label()
		name, nameErr := vol.Name()
		deviceID, deviceIDErr := vol.DeviceID()
		devicePath, devicePathErr := vol.DevicePath()

		fmt.Printf("Volume Label: %s\n", strOrErr(label, labelErr))
		fmt.Printf("Volume Name: %s\n", strOrErr(name, nameErr))
		fmt.Printf("Device ID: %s\n", strOrErr(fmt.Sprintf("%s", deviceID), deviceIDErr))
		fmt.Printf("NT Namespace Device Path: %s\n", strOrErr(devicePath, devicePathErr))
		fmt.Printf("Device Information: Number %d, Partition %d, Type %d\n", vol.DeviceNumber(), vol.PartitionNumber(), vol.DeviceType())
		fmt.Printf("Device Description: Removable: %t, Vendor: %s, Product: %s, Revision: %s, OS S/N: %s\n", vol.RemovableMedia(), vol.VendorID(), vol.ProductID(), vol.ProductRevision(), vol.SerialNumber())

		paths, err := vol.Paths()
		if err != nil {
			fmt.Printf("%v\n", fmt.Errorf("Unable to ascertain volume paths: %v", err))
		} else if len(paths) > 0 {
			fmt.Printf("Mounts:\n")
			for i, path := range paths {
				root, _ := volumeapi.GetVolumeNameForVolumeMountPoint(path)
				fmt.Printf("  Mount %d on \"%s\": \"%s\"\n", i, root, path)
			}
		}

		journal := vol.Journal()
		defer journal.Close()

		journalData, journalDataErr := journal.Query()
		if journalDataErr != nil {
			fmt.Printf("USN Journal: %v\n", journalDataErr)
			continue
		}

		fmt.Printf("USN Journal: Present, ID: %d, Next USN: %d, Supporting Versions: %d-%d\n", journalData.JournalID, journalData.NextUSN, journalData.MinSupportedMajorVersion, journalData.MaxSupportedMajorVersion)

		cursor, cursorErr := journal.Cursor(nil, usn.ReasonFileCreate|usn.ReasonFileDelete,
			nil, nil)
		if cursorErr != nil {
			fmt.Printf("Unable to create USN journal cursor: %v\n", cursorErr)
			continue
		}
		defer cursor.Close()

		buffer := make([]byte, 262144)
		i := 0
		for {
			records, cursorErr := cursor.Next(buffer)
			if cursorErr != nil {
				if cursorErr != io.EOF {
					fmt.Printf("Unable to retreive USN journal records: %v\n", cursorErr)
				}
				break
			}

			for _, record := range records {
				if regex != nil {
					if !regex.MatchString(record.FileName) {
						continue
					}
				}
				action := "OTHER "
				if record.Reason&usn.ReasonFileCreate != 0 {
					action = "CREATE"
				}
				if record.Reason&usn.ReasonFileDelete != 0 {
					action = "DELETE"
				}
				fmt.Printf("%s  %s  %s\n", record.TimeStamp.Format("2006-01-02 15:04:05.000000"), action, record.FileName)
				i++
			}
		}
	}
	fmt.Printf("--------\n")
}

func strOrErr(s string, err error) string {
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return fmt.Sprintf("\"%s\"", s)
}
