package hsync

import (
	"errors"
	"sync"
	"syscall"
)

// ErrClosed is returned when a handle is already closed.
var ErrClosed = errors.New("handle already closed")

// Handle provides shared access to a system handle by one or more instances
// without fear of the handle being closed prematurely. It should be created
// initially by calling New(). Additional instances are created by calling
// Clone().
//
// Handle maintains a reference counter internally that tracks the number of
// instances. The system handle will be closed when all of the instances relying
// on it have been closed.
//
// It is very important that each instance be closed. Failure to do so may
// result in leaked system handles.
type Handle struct {
	mutex  sync.RWMutex
	source *coordinator
}

// New returns a new Handle that guards access to the given system handle. It
// should only be called once for a particular system handle. Additional copies
// can be created by calling Clone().
//
// When finished with the returned Handle, it is the caller's responsibility to
// release its resources by calling Close().
func New(handle syscall.Handle) *Handle {
	source := &coordinator{
		handle: handle,
	}
	return &Handle{source: source}
}

// Clone returns an independent copy of h. When finished with a clone, it is the
// caller's responsibility to close it.
//
// If h has already been closed calling Clone() will result in a panic.
func (h *Handle) Clone() *Handle {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	if h.source == nil {
		panic("hsync: attempt to clone a handle that has already been closed")
	}
	h.source.Add(1)
	return &Handle{
		source: h.source,
	}
}

// Close prevents further use of h and indicates that h no longer requires
// access to its system handle. If h is the last instance referring to a
// particular system handle the system handle will be closed.
//
// If h has already been closed ErrClosed will be returned.
//
// Once h has been closed it cannot be cloned.
func (h *Handle) Close() (err error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if h.source == nil {
		return ErrClosed
	}
	h.source.Add(-1)
	h.source = nil
	return nil
}

// Handle returns the system handle protected by h.
//
// If h has already been closed calling Handle() will result in a panic.
func (h *Handle) Handle() syscall.Handle {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	if h.source == nil {
		panic("hsync: attempt to retreive a system handle that has already been closed")
	}
	return h.source.Handle()
}
