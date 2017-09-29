package hsync

import (
	"sync"
	"syscall"
)

// coordinator acts as a central mutex for synchronized access to
// system handles.
type coordinator struct {
	mutex  sync.RWMutex
	handle syscall.Handle
	refs   int // when >= 0 h is in use
}

// Add will change the internal reference count by the given delta. If the
// reference count drops below zero, the system handle will be released. If the
// reference count is already below zero it will panic.
func (c *coordinator) Add(delta int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.refs < 0 {
		panic("hsync: attempted use of closed system handle: reference counter already below zero")
	}
	c.refs += delta
	if c.refs < 0 {
		syscall.CloseHandle(c.handle)
	}
	return nil
}

func (c *coordinator) Handle() syscall.Handle {
	return c.handle
}
