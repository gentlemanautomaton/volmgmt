package changejournal

import "errors"

// Create will create or modify a change journal on the specificied volume
// with the given parameters.
func Create(vol Volume, maximumSize, AllocationDelta uint64) error {
	return errors.New("Not implemented yet")
}
