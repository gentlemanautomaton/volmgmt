package main

import (
	"context"
	"fmt"
	"sort"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gentlemanautomaton/volmgmt/fileapi"
	"github.com/gentlemanautomaton/volmgmt/fileattr"
	"github.com/gentlemanautomaton/volmgmt/usn"
	"github.com/gentlemanautomaton/volmgmt/volume"
	"golang.org/x/sys/windows"
)

func scan(ctx context.Context, path string, settings Settings) (summary Summary) {
	fmt.Printf("Path: \"%s\"\n", path)

	if settingsSummary := settings.Summary(); settingsSummary != "" {
		fmt.Printf(settingsSummary)
	}

	vol, err := volume.New(path)
	if err != nil {
		fmt.Printf("Unable to read \"%s\": %v\n", path, err)
		return
	}
	defer vol.Close()

	volName, err := vol.Name()
	if err != nil {
		fmt.Printf("Unable to determine volume name for \"%s\": %v\n", path, err)
		return
	}
	volHandle := vol.Handle()

	mft := vol.MFT()
	defer mft.Close()

	iter, err := mft.Enumerate(nil, usn.Min, usn.Max)
	if err != nil {
		fmt.Printf("Unable to open MFT: %v\n", err)
		return
	}
	defer iter.Close()

	cache := usn.NewCache()
	{
		fmt.Printf("Scanning MFT...")
		start := time.Now()
		err = cache.ReadFrom(ctx, iter)
		end := time.Now()
		duration := end.Sub(start)
		if err != nil {
			fmt.Printf(" failed: %v. Ran %s.\n", err, duration)
			return
		}
		fmt.Printf(" done. Completed in %s.\n", duration)
	}

	var records []usn.Record
	{
		fmt.Printf("Sorting file paths...")
		start := time.Now()
		records = cache.Records()
		sort.Slice(records, func(i, j int) bool {
			return records[i].Path < records[j].Path
		})
		end := time.Now()
		duration := end.Sub(start)
		fmt.Printf(" done. Completed in %s.\n", duration)
	}
	summary.Sizes = make([]int64, 0, len(records))

	recordFilter := buildRecordFilter(settings)
	fileInfoFilter := buildFileInfoFilter(settings)
	{
		fmt.Printf("Scanning files...\n")
		start := time.Now()
		type token struct{}
		sem := make(chan token, settings.Limit)
		for i, record := range records {
			sem <- token{}
			if settings.Progress && i%5000 == 0 {
				fmt.Printf("Scanning files... (%d/%d) %d%%\n", i, len(records), percent(i, len(records)))
			}
			if ctx.Err() != nil {
				return
			}
			go func(i int, record usn.Record) {
				processRecord(i, record, volHandle, volName, recordFilter, fileInfoFilter, settings.List, settings.Verbose, &summary)
				<-sem
			}(i, record)
		}
		end := time.Now()
		duration := end.Sub(start)
		fmt.Printf("Scanning files... done. Completed in %s.\n", duration)
	}

	return summary
}

func processRecord(index int, record usn.Record, volHandle syscall.Handle, volName string, recordFilter usn.Filter, fileInfoFilter FileInfoFilter, list, verbose bool, summary *Summary) {
	if record.FileAttributes.Match(fileattr.ReparsePoint) {
		summary.Skipped++
		return
	}
	if !recordFilter.Match(record) {
		return
	}
	if record.FileAttributes.Match(fileattr.Directory) {
		summary.Directories++
		return
	}

	const access = uint32(windows.READ_CONTROL)
	const shareMode = uint32(syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE | syscall.FILE_SHARE_DELETE)
	fileHandle, err := fileapi.OpenFileByID(volHandle, record.FileReferenceNumber, access, shareMode, syscall.FILE_FLAG_BACKUP_SEMANTICS)
	if err != nil {
		if verbose {
			fmt.Printf("%10d: %s: can't open file: %v\n", index, record.Path, err)
		}
		summary.Skipped++
		return
	}
	defer syscall.CloseHandle(fileHandle)

	fileInfo := fileapi.FileInfoForHandle{
		FileName: record.FileName,
	}
	fileInfo.ByHandleFileInformation, err = fileapi.GetFileInformationByHandle(fileHandle)
	if err != nil {
		if verbose {
			fmt.Printf("%10d: %s: can't get file info: %v\n", index, record.Path, err)
		}
		summary.Skipped++
		return
	}

	if fileInfo.IsDir() {
		return
	}
	if !fileInfoFilter(fileInfo) {
		return
	}
	size := fileInfo.Size()

	summary.Files++
	summary.TotalBytes += size
	summary.Sizes = append(summary.Sizes, size)
	if list {
		fmt.Printf("%10d: %s: %s\n", index, record.Path, humanize.Bytes(uint64(fileInfo.Size())))
	}
}

func percent(i, total int) int {
	if i <= 0 || total <= 0 {
		return 0
	}
	return (i * 100) / total
}

func printVolume(vol *volume.Volume) {
	label, labelErr := vol.Label()
	name, nameErr := vol.Name()
	devicePath, devicePathErr := vol.DevicePath()

	fmt.Printf("Volume Label: %s\n", strOrErr(label, labelErr))
	fmt.Printf("Volume Name: %s\n", strOrErr(name, nameErr))
	fmt.Printf("NT Namespace Device Path: %s\n", strOrErr(devicePath, devicePathErr))
	fmt.Printf("Device Information: Number %d, Partition %d, Type %d\n", vol.DeviceNumber(), vol.PartitionNumber(), vol.DeviceType())
	fmt.Printf("Device Description: Removable: %t, Vendor: %s, Product: %s, Revision: %s, OS S/N: %s\n", vol.RemovableMedia(), vol.VendorID(), vol.ProductID(), vol.ProductRevision(), vol.SerialNumber())
}

func printStats(total, filtered usn.Stats) {
	var percent float32
	if total.Records > 0 {
		percent = float32(filtered.Records) / float32(total.Records) * 100
	}
	fmt.Printf("Matched:     %d/%d USN journal records (%.2f%%)\n", filtered.Records, total.Records, percent)
	fmt.Printf("First Match: %s\n", filtered.First)
	fmt.Printf("Last Match:  %s\n", filtered.Last)
}

func strOrErr(s string, err error) string {
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return fmt.Sprintf("\"%s\"", s)
}
