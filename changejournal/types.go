package changejournal

import "syscall"

// Volume represents any type that can provide a volume handle.
type Volume interface {
	Handle() syscall.Handle
}
