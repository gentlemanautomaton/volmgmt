package usn

import (
	"fmt"
	"strings"
)

// Reason describes a kind of USN record, that is the reason for which it was
// recorded. A record may have been created for more than one reason.
type Reason uint32

// USN Journal Reason Codes
//
// https://msdn.microsoft.com/library/windows/desktop/hh802706
const (
	ReasonDataOverwrite       Reason = 0x00000001 // USN_REASON_DATA_OVERWRITE
	ReasonDataExtend          Reason = 0x00000002 // USN_REASON_DATA_EXTEND
	ReasonDataTruncation      Reason = 0x00000004 // USN_REASON_DATA_TRUNCATION
	ReasonNamedDataOverwrite  Reason = 0x00000010 // USN_REASON_NAMED_DATA_OVERWRITE
	ReasonNamedDataExtend     Reason = 0x00000020 // USN_REASON_NAMED_DATA_EXTEND
	ReasonNamedDataTruncation Reason = 0x00000040 // USN_REASON_NAMED_DATA_TRUNCATION
	ReasonFileCreate          Reason = 0x00000100 // USN_REASON_FILE_CREATE
	ReasonFileDelete          Reason = 0x00000200 // USN_REASON_FILE_DELETE
	ReasonEAChange            Reason = 0x00000400 // USN_REASON_EA_CHANGE
	ReasonSecurityChange      Reason = 0x00000800 // USN_REASON_SECURITY_CHANGE
	ReasonRenameOldName       Reason = 0x00001000 // USN_REASON_RENAME_OLD_NAME
	ReasonRenameNewName       Reason = 0x00002000 // USN_REASON_RENAME_NEW_NAME
	ReasonRename              Reason = ReasonRenameOldName | ReasonRenameNewName
	ReasonIndexableChange     Reason = 0x00004000 // USN_REASON_INDEXABLE_CHANGE
	ReasonBasicInfoChange     Reason = 0x00008000 // USN_REASON_BASIC_INFO_CHANGE
	ReasonHardLinkChange      Reason = 0x00010000 // USN_REASON_HARD_LINK_CHANGE
	ReasonCompressionChange   Reason = 0x00020000 // USN_REASON_COMPRESSION_CHANGE
	ReasonEncryptionChange    Reason = 0x00040000 // USN_REASON_ENCRYPTION_CHANGE
	ReasonObjectIDChange      Reason = 0x00080000 // USN_REASON_OBJECT_ID_CHANGE
	ReasonReparsePointChange  Reason = 0x00100000 // USN_REASON_REPARSE_POINT_CHANGE
	ReasonStreamChange        Reason = 0x00200000 // USN_REASON_STREAM_CHANGE
	ReasonTransactedChange    Reason = 0x00400000 // USN_REASON_TRANSACTED_CHANGE
	ReasonIntegrityChange     Reason = 0x00800000 // USN_REASON_INTEGRITY_CHANGE
	ReasonClose               Reason = 0x80000000 // USN_REASON_CLOSE
	ReasonAny                 Reason = 0xffffffff
)

// ReasonFormat describes a format for reason names
type ReasonFormat map[Reason]string

// ReasonFormatConstant maps the official C-style constant strings to reason
// codes.
var ReasonFormatConstant = ReasonFormat{
	ReasonDataOverwrite:       "USN_REASON_DATA_OVERWRITE",
	ReasonDataExtend:          "USN_REASON_DATA_EXTEND",
	ReasonDataTruncation:      "USN_REASON_DATA_TRUNCATION",
	ReasonNamedDataOverwrite:  "USN_REASON_NAMED_DATA_OVERWRITE",
	ReasonNamedDataExtend:     "USN_REASON_NAMED_DATA_EXTEND",
	ReasonNamedDataTruncation: "USN_REASON_NAMED_DATA_TRUNCATION",
	ReasonFileCreate:          "USN_REASON_FILE_CREATE",
	ReasonFileDelete:          "USN_REASON_FILE_DELETE",
	ReasonEAChange:            "USN_REASON_EA_CHANGE",
	ReasonSecurityChange:      "USN_REASON_SECURITY_CHANGE",
	ReasonRenameOldName:       "USN_REASON_RENAME_OLD_NAME",
	ReasonRenameNewName:       "USN_REASON_RENAME_NEW_NAME",
	ReasonIndexableChange:     "USN_REASON_INDEXABLE_CHANGE",
	ReasonBasicInfoChange:     "USN_REASON_BASIC_INFO_CHANGE",
	ReasonHardLinkChange:      "USN_REASON_HARD_LINK_CHANGE",
	ReasonCompressionChange:   "USN_REASON_COMPRESSION_CHANGE",
	ReasonEncryptionChange:    "USN_REASON_ENCRYPTION_CHANGE",
	ReasonObjectIDChange:      "USN_REASON_OBJECT_ID_CHANGE",
	ReasonReparsePointChange:  "USN_REASON_REPARSE_POINT_CHANGE",
	ReasonStreamChange:        "USN_REASON_STREAM_CHANGE",
	ReasonTransactedChange:    "USN_REASON_TRANSACTED_CHANGE",
	ReasonIntegrityChange:     "USN_REASON_INTEGRITY_CHANGE",
	ReasonClose:               "USN_REASON_CLOSE",
}

// ReasonFormatBasic maps basic reason strings to reason codes.
var ReasonFormatBasic = ReasonFormat{
	ReasonDataOverwrite:       "DataOverwrite",
	ReasonDataExtend:          "DataExtend",
	ReasonDataTruncation:      "DataTruncation",
	ReasonNamedDataOverwrite:  "NamedDataOverwrite",
	ReasonNamedDataExtend:     "NamedDataExtend",
	ReasonNamedDataTruncation: "NamedDataTruncation",
	ReasonFileCreate:          "FileCreate",
	ReasonFileDelete:          "FileDelete",
	ReasonEAChange:            "EAChange",
	ReasonSecurityChange:      "SecurityChange",
	ReasonRenameOldName:       "RenameOldName",
	ReasonRenameNewName:       "RenameNewName",
	ReasonRename:              "Rename",
	ReasonIndexableChange:     "IndexableChange",
	ReasonBasicInfoChange:     "BasicInfoChange",
	ReasonHardLinkChange:      "HardLinkChange",
	ReasonCompressionChange:   "CompressionChange",
	ReasonEncryptionChange:    "EncryptionChange",
	ReasonObjectIDChange:      "ObjectIDChange",
	ReasonReparsePointChange:  "ReparsePointChange",
	ReasonStreamChange:        "StreamChange",
	ReasonTransactedChange:    "TransactedChange",
	ReasonIntegrityChange:     "IntegrityChange",
	ReasonClose:               "Close",
}

// ReasonFormatShort maps short reason strings to reason codes.
var ReasonFormatShort = ReasonFormat{
	ReasonDataOverwrite:       "Overwrite",
	ReasonDataExtend:          "Extend",
	ReasonDataTruncation:      "Truncation",
	ReasonNamedDataOverwrite:  "NamedDataOverwrite",
	ReasonNamedDataExtend:     "NamedDataExtend",
	ReasonNamedDataTruncation: "NamedDataTruncation",
	ReasonFileCreate:          "Create",
	ReasonFileDelete:          "Delete",
	ReasonEAChange:            "EAChange",
	ReasonSecurityChange:      "SecurityChange",
	ReasonRenameOldName:       "RenameOldName",
	ReasonRenameNewName:       "RenameNewName",
	ReasonRename:              "Rename",
	ReasonIndexableChange:     "IndexableChange",
	ReasonBasicInfoChange:     "BasicInfoChange",
	ReasonHardLinkChange:      "HardLinkChange",
	ReasonCompressionChange:   "CompressionChange",
	ReasonEncryptionChange:    "EncryptionChange",
	ReasonObjectIDChange:      "ObjectIDChange",
	ReasonReparsePointChange:  "ReparsePointChange",
	ReasonStreamChange:        "StreamChange",
	ReasonTransactedChange:    "TransactedChange",
	ReasonIntegrityChange:     "IntegrityChange",
	ReasonClose:               "Close",
}

// ParseReason interprets the given string as one or more reason codes and
// returns a uint32 representing their combined bitmask.
func ParseReason(reason string) (code Reason, err error) {
	reason = strings.ToLower(reason)
	reason = strings.Replace(reason, ",", "|", -1)
	parts := strings.Split(reason, "|")

	for _, part := range parts {
		switch part {
		case "overwrite", "dataoverwrite", "reasondataoverwrite", "usn_reason_data_overwrite":
			code |= ReasonDataOverwrite
		case "extend", "dataextend", "reasondataextend", "usn_reason_data_extend":
			code |= ReasonDataExtend
		case "truncation", "datatruncation", "reasondatatruncation", "usn_reason_data_truncation":
			code |= ReasonDataTruncation
		case "create", "filecreate", "reasonfilecreate", "usn_reason_file_create":
			code |= ReasonFileCreate
		case "delete", "filedelete", "reasonfiledelete", "usn_reason_file_delete":
			code |= ReasonFileDelete
		case "move":
			code |= ReasonRenameOldName
			code |= ReasonRenameNewName
		case "all", "any", "*":
			code |= ReasonAny
		default:
			err = fmt.Errorf("unsupported or unknown reason code: %s", part)
		}
	}

	return
}

// String returns a string representation of the reason code using a default
// format and separator.
func (r Reason) String() string {
	return r.Join("|", ReasonFormatConstant)
}

// Join returns a string representation of the reason code using the given
// separator and reason format.
func (r Reason) Join(sep string, format ReasonFormat) string {
	// If the format has an exact match, just return it
	if value, ok := format[r]; ok {
		return value
	}

	var values []string
	rename := false
	for i := 0; i < 32; i++ {
		code := Reason(1 << uint32(i))
		if r.Match(code) {
			// Special handling for renames/moves
			if code == ReasonRenameOldName && rename {
				continue
			}
			if code == ReasonRenameNewName && r.Match(ReasonRename) {
				if value, ok := format[ReasonRename]; ok {
					rename = true
					values = append(values, value)
					continue
				}
			}

			// Typical handling
			if value, ok := format[code]; ok {
				values = append(values, value)
			}
		}
	}

	return strings.Join(values, sep)
}

// Match reports whether r contains all of the reason codes specified by
// c.
func (r Reason) Match(c Reason) bool {
	return r&c == c
}

// Rename returns true if r specifies both ReasonRenameOldName and
// ReasonRenameNewName.
func (r Reason) Rename() bool {
	return r.Match(ReasonRenameOldName | ReasonRenameNewName)
}
