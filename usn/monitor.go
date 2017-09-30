package usn

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"syscall"
	"time"

	"github.com/gentlemanautomaton/volmgmt/hsync"
	"github.com/gentlemanautomaton/volmgmt/volumeapi"
)

const (
	// MinimumPollingInterval is the minimum amount of time that a monitor will
	// wait between change journal read requests.
	MinimumPollingInterval = time.Millisecond

	// DefaultPollingInterval is the suggested amount of time that a monitor
	// should wait between change journal read requests.
	DefaultPollingInterval = time.Second
)

var (
	// ErrClosed is returned when a monitor is already closed.
	ErrClosed = errors.New("monitor already closed")

	// ErrRunning is returned when an attempt is made to start an already
	// running monitor.
	ErrRunning = errors.New("the monitor has already been started")

	// ErrStopped is returned when an attempt is made to stop a monitor that isn't
	// running.
	ErrStopped = errors.New("the monitor has already been stopped")
)

// Monitor facilitates monitoring of USN journals.
type Monitor struct {
	cursor *Cursor // Used by m.run without acquiring a lock when it's running

	mutex     sync.RWMutex
	h         *hsync.Handle // Cloned for each cursor when it's created
	listeners []chan Record
	sigstop   chan struct{} // nil when not running, close to stop m.run
	stopped   chan struct{} // nil when not running, closed by m.run when exited
	closed    bool
}

// NewMonitor returns a USN change journal monitor for the volume described by
// path.
func NewMonitor(path string) (monitor *Monitor, err error) {
	const (
		access = syscall.GENERIC_READ
		mode   = syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE
	)

	h, err := volumeapi.Handle(path, access, mode)
	if err != nil {
		return nil, err
	}

	return NewMonitorWithHandle(hsync.New(h)), nil
}

// NewMonitorWithHandle returns a USN journal monitor for the volume with the
// given handle. The returned monitor will be inactive until it has been
// started by a call to Run().
//
// When the monitor is closed its associated handle will also be closed. When
// providing an existing handle that will be used elsewhere be sure to
// clone it first.
//
// It is the caller's responsibility to close the monitor when finished with it.
func NewMonitorWithHandle(handle *hsync.Handle) *Monitor {
	return &Monitor{
		h: handle,
	}
}

// Run will cause the monitor to start observing its USN journal. Reads will
// start from the given update sequence number and will continue until the
// monitor is stopped or closed.
//
// Records retrieved from the journal will be broadcast to all registered
// listeners.
//
// The monitor will start reading records from the update sequence number
// specified by start. If start is zero the monitor will read from the
// beginning of the journal.
//
// If the monitor has already been started an error will be returned.
func (m *Monitor) Run(start USN, interval time.Duration, reasonMask uint32) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.closed {
		return ErrClosed
	}

	if m.sigstop != nil {
		return ErrRunning
	}

	if interval < MinimumPollingInterval {
		interval = MinimumPollingInterval
	}

	if m.cursor == nil {
		var err error
		m.cursor, err = NewCursorWithHandle(m.h.Clone(), reasonMask)
		if err != nil {
			return fmt.Errorf("unable to created cursor for volume handle: %v", err)
		}
	}

	m.cursor.usn = start

	m.sigstop = make(chan struct{})
	m.stopped = make(chan struct{})

	go m.run(interval, m.sigstop, m.stopped)

	return nil
}

func (m *Monitor) run(interval time.Duration, sigstop, stopped chan struct{}) {
	defer close(stopped)

	var (
		buffer [65536]byte
		p      = buffer[:]
	)

	for {
		select {
		case <-sigstop:
			return
		default:
		}

		// TODO: Achieve a zero-allocating solution by updating cursor.Next to
		// receive a buffer of records in addition to the buffer of bytes.
		records, err := m.cursor.Next(p)

		if len(records) > 0 {
			m.broadcast(records)
		}

		switch err {
		case nil:
		case io.EOF:
			t := time.NewTimer(interval)
			select {
			case <-sigstop:
				if !t.Stop() {
					<-t.C
				}
			case <-t.C:
			}
		default:
			// FIXME: Handle expected but infrequent errors, such as journal wraps

			// FIXME: Report non-nil errors in some way?
			//fmt.Printf("monitor cursor error: %v\n", err)
			return
		}
	}
}

// Stop will cause the monitor to stop observing the USN journal.
func (m *Monitor) Stop() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.closed {
		return ErrClosed
	}

	if m.sigstop == nil {
		return ErrStopped
	}

	close(m.sigstop)
	<-m.stopped

	m.sigstop = nil
	m.stopped = nil

	return nil
}

// Close releases any resources consumed by the monitor. All active listeners
// will be closed and the monitor will stop observing the journal.
//
// Once a monitor has been closed it cannot be used.
func (m *Monitor) Close() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.closed {
		return ErrClosed
	}
	m.closed = true

	if m.sigstop != nil {
		close(m.sigstop)
		<-m.stopped
		m.sigstop = nil
		m.stopped = nil
	}
	if m.cursor != nil {
		m.cursor.Close()
	}

	for _, listener := range m.listeners {
		close(listener)
	}
	m.listeners = nil

	return nil
}

// Listen returns a channel on which USN journal updates will be broadcast.
// The channel will be closed when the monitor is closed or when unlisten is
// called for the returned channel.
func (m *Monitor) Listen(chanSize int) <-chan Record {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	ch := make(chan Record, chanSize)
	if m.closed {
		close(ch)
	} else {
		m.listeners = append(m.listeners, ch)
	}

	return ch
}

// Unlisten closes the given listener's channel and removes it from the set of
// listeners that receive records from the monitor.
//
// Unlisten returns false if the listener was not present.
func (m *Monitor) Unlisten(c <-chan Record) (found bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for i := 0; i < len(m.listeners); i++ {
		entry := m.listeners[i]
		if m.listeners[i] == c {
			if i+1 < len(m.listeners) {
				m.listeners = append(m.listeners[:i], m.listeners[i+1:]...)
			} else {
				m.listeners = m.listeners[:i]
			}
			found = true
		}
		i--
		close(entry)
	}
	return
}

func (m *Monitor) broadcast(records []Record) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if m.closed || len(m.listeners) == 0 {
		return
	}

	for r := range records {
		for i := range m.listeners {
			m.listeners[i] <- records[r]
		}
	}
}
