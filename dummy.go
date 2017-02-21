package accessor

import (
	"fmt"
)

// DummyAccessor is the Accessor for testing.
type DummyAccessor struct {
	ID int
}

// Get implements Accessor.
func (a DummyAccessor) Get(path Path) (Accessor, error) {
	return nil, fmt.Errorf("this is dummy accessor: %d", a.ID)
}

// Set implements Accessor.
func (a DummyAccessor) Set(path Path, _ interface{}) error {
	return fmt.Errorf("this is dummy accessor: %d", a.ID)
}

// Unwrap implements Accessor.
func (a DummyAccessor) Unwrap() interface{} {
	return a.ID
}

// Foreach implements Accessor.
func (a DummyAccessor) Foreach(f func(path Path, value interface{}) error) error {
	return f(initialPath{}, a.ID)
}
