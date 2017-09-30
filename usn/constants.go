package usn

import (
	"fmt"
	"strings"
)

// USN Journal Source Info Codes
const (
	SourceDataManagement              = 0x00000001 // USN_SOURCE_DATA_MANAGEMENT
	SourceAuxilaryData                = 0x00000002 // USN_SOURCE_AUXILIARY_DATA
	SourceReplicationManagement       = 0x00000004 // USN_SOURCE_REPLICATION_MANAGEMENT
	SourceClientReplicationManagement = 0x00000008 // USN_SOURCE_CLIENT_REPLICATION_MANAGEMENT
)

// USN Journal Reason Codes
//
// https://msdn.microsoft.com/library/windows/desktop/hh802706
const (
	ReasonDataOverwrite       = 0x00000001 // USN_REASON_DATA_OVERWRITE
	ReasonDataExtend          = 0x00000002 // USN_REASON_DATA_EXTEND
	ReasonDataTruncation      = 0x00000004 // USN_REASON_DATA_TRUNCATION
	ReasonNamedDataOverwrite  = 0x00000010 // USN_REASON_NAMED_DATA_OVERWRITE
	ReasonNamedDataExtend     = 0x00000020 // USN_REASON_NAMED_DATA_EXTEND
	ReasonNamedDataTruncation = 0x00000040 // USN_REASON_NAMED_DATA_TRUNCATION
	ReasonFileCreate          = 0x00000100 // USN_REASON_FILE_CREATE
	ReasonFileDelete          = 0x00000200 // USN_REASON_FILE_DELETE
	ReasonEAChange            = 0x00000400 // USN_REASON_EA_CHANGE
	ReasonSecurityChange      = 0x00000800 // USN_REASON_SECURITY_CHANGE
	ReasonRenameOldName       = 0x00001000 // USN_REASON_RENAME_OLD_NAME
	ReasonRenameNewName       = 0x00002000 // USN_REASON_RENAME_NEW_NAME
	ReasonIndexableChange     = 0x00004000 // USN_REASON_INDEXABLE_CHANGE
	ReasonBasicInfoChange     = 0x00008000 // USN_REASON_BASIC_INFO_CHANGE
	ReasonHardLinkChange      = 0x00010000 // USN_REASON_HARD_LINK_CHANGE
	ReasonCompressionChange   = 0x00020000 // USN_REASON_COMPRESSION_CHANGE
	ReasonEncryptionChange    = 0x00040000 // USN_REASON_ENCRYPTION_CHANGE
	ReasonObjectIDChange      = 0x00080000 // USN_REASON_OBJECT_ID_CHANGE
	ReasonReparsePointChange  = 0x00100000 // USN_REASON_REPARSE_POINT_CHANGE
	ReasonStreamChange        = 0x00200000 // USN_REASON_STREAM_CHANGE
	ReasonTransactedChange    = 0x00400000 // USN_REASON_TRANSACTED_CHANGE
	ReasonIntegrityChange     = 0x00800000 // USN_REASON_INTEGRITY_CHANGE
	ReasonClose               = 0x80000000 // USN_REASON_CLOSE
)

// ParseReason interprets the given string as one or more reason codes and
// returns a uint32 representing their combined bitmask.
func ParseReason(reason string) (code uint32, err error) {
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
		default:
			err = fmt.Errorf("unsupported or unknown reason code: %s", part)
		}
	}

	return
}
