package accessor

import (
	"fmt"
)

type DummyAccessor struct {
	ID int
}

func (o DummyAccessor) Get(path string, paths ...string) (Accessor, error) {
	return nil, fmt.Errorf("this is dummy accessor: %d", o.ID)
}

func (o DummyAccessor) Set(acc Accessor, path string, paths ...string) error {
	return fmt.Errorf("this is dummy accessor: %d", o.ID)
}

func (o DummyAccessor) Unwrap() interface{} {
	return o.ID
}
