package usnsource

// Format describes a format for USN source information.
type Format map[Info]string

// FormatC maps source information to C-style constant strings.
var FormatC = Format{
	DataManagement:              "USN_SOURCE_DATA_MANAGEMENT",
	AuxilaryData:                "USN_SOURCE_AUXILIARY_DATA",
	ReplicationManagement:       "USN_SOURCE_REPLICATION_MANAGEMENT",
	ClientReplicationManagement: "USN_SOURCE_CLIENT_REPLICATION_MANAGEMENT",
}

// FormatGo maps source information to Go-style constant strings.
var FormatGo = Format{
	Local:                       "Local",
	DataManagement:              "DataManagement",
	AuxilaryData:                "AuxilaryData",
	ReplicationManagement:       "ReplicationManagement",
	ClientReplicationManagement: "ClientReplicationManagement",
}

// FormatShort maps source information to short strings.
var FormatShort = Format{
	Local:                       "LOCAL",
	DataManagement:              "OS",
	AuxilaryData:                "AUX",
	ReplicationManagement:       "REPL",
	ClientReplicationManagement: "CLOUD",
}
