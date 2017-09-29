package hsync_test

import (
	"sync"
	"syscall"

	"github.com/gentlemanautomaton/volmgmt/hsync"
)

// acquireHandle acquires a system handle from some source, probably by
// calling a function in syscall.
func acquireHandle() syscall.Handle {
	return 0 // System handles are actually uintptr values
}

// work performs some work that requires access to a system handle.
func work(wg *sync.WaitGroup, h *hsync.Handle) {
	defer wg.Done()
	defer h.Close()
	_ = h.Handle() // Do something with the actual system handle
}

const workers = 5

func Example() {
	var wg sync.WaitGroup
	wg.Add(workers)

	h := hsync.New(acquireHandle())

	for i := 0; i < workers; i++ {
		go work(&wg, h.Clone())
	}

	// We can close our instance as soon as we're done with it, even though
	// the workers may still be relying on their cloned copies.
	h.Close()

	wg.Wait()
}
