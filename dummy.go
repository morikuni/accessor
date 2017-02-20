package accessor

import (
	"fmt"
)

type DummyAccessor struct {
	ID int
}

func (a DummyAccessor) Get(path Path) (Accessor, error) {
	return nil, fmt.Errorf("this is dummy accessor: %d", a.ID)
}

func (a DummyAccessor) Set(path Path, _ interface{}) error {
	return fmt.Errorf("this is dummy accessor: %d", a.ID)
}

func (a DummyAccessor) Unwrap() interface{} {
	return a.ID
}

func (a DummyAccessor) Foreach(f func(path Path, value interface{}) error) error {
	return fmt.Errorf("this is dummy accessor: %d", a.ID)
}
