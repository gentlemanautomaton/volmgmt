package volumeapi

import (
	"fmt"
	"path/filepath"
	"strings"
	"syscall"
)

// Handle attempts to acquire a system handle for the volume represented by
// path (or upon which path is mounted).
//
// The given path may be in any of these forms:
//
//  C:\path\to\directory\
//  X:\
//  Y:\MountX\
//  \\.\X:
//  \\?\Volume{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}\
//  \\.\PhysicalDrive0
func Handle(path string, access uint32, mode uint32) (syscall.Handle, error) {
	if len(path) == 0 {
		return 0, syscall.ERROR_FILE_NOT_FOUND
	}

	mount, name, err := MountPoint(path)
	if err != nil {
		return 0, err
	}

	if strings.HasPrefix(name, `\\?\Volume`) {
		name = strings.TrimRight(name, `\`)
	}

	namep, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}

	h, err := syscall.CreateFile(namep, access, mode, nil, syscall.OPEN_EXISTING, 0, 0)
	if err != nil {
		var ref string
		if mount == path {
			ref = fmt.Sprintf("\"%s\"", mount)
		} else {
			ref = fmt.Sprintf("\"%s\" (mounted at \"%s\")", path, mount)
		}
		return 0, fmt.Errorf("unable to create system handle for %s: %v", ref, err)
	}
	return h, nil
}

// MountPoint attempts to find the volume mount point for s. If successful, it
// returns both the volume mount point and the volune name (also known as the
// volume GUID path).
//
// s may be in any of these forms:
//
//  C:\path\to\directory\
//  X:\
//  Y:\MountX\
//  \\.\X:
//  \\?\Volume{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}\
//  \\.\PhysicalDrive0
//
// If successful, the returned volume name will be of this form:
//
//   \\?\Volume{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}\
func MountPoint(s string) (mountPoint, volumeName string, err error) {
	if len(s) > 4 {
		switch s[0:4] {
		case `\\.\`, `\\?\`:
			// The path is already in some form that directly describes a mount point
			return s, s, nil
		}
	}
	if len(s) == 2 {
		if s[1] == ':' {
			s += `\`
		}
	}

	const sep = string(filepath.Separator)

	mountPoint = s
	volumeName, err = GetVolumeNameForVolumeMountPoint(s)
	if err != nil {
		// Assume s is a path. Try walking up the directory tree until we hit a
		// mount point.
		path := filepath.Clean(s)
		for {
			last := path
			// On Windows mount points must have a trailing slash, but filepath.Clean
			// removes it unless it's at the root. We add it back here if it's
			// missing.
			mountPoint = addTrailingSlash(path)
			volumeName, err = GetVolumeNameForVolumeMountPoint(mountPoint)
			if err == nil {
				break
			}
			path = filepath.Dir(path)
			if path == last {
				break
			}
		}
	}
	if err != nil {
		err = fmt.Errorf("unable to determine mount point for \"%s\": %v", s, err)
		mountPoint = ""
	}
	return
}

// addTrailingSlash adds filepath.Separator to the end of the given string if
// it isn't present already.
func addTrailingSlash(path string) string {
	if last := len(path) - 1; last >= 0 && path[last] != filepath.Separator {
		return path + string(filepath.Separator)
	}
	return path
}
