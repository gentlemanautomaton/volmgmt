package ioctlcode

// IO Control Code Method Types
const (
	MethodBuffered  = iota // 0
	MethodInDirect         // 1
	MethodOutDirect        // 2
	MethodNeither          // 3
)

// Access Restrictions
const (
	AccessAny       = iota // 0 FILE_ANY_ACCESS
	AccessRead             // 1 FILE_READ_ACCESS
	AccessWrite            // 2 FILE_WRITE_ACCESS
	AccessReadWrite = AccessRead | AccessWrite
	AccessSpecial   = AccessAny
)
