package accessor

import (
	"fmt"
)

type DummyAccessor struct {
	ID int
}

func (o DummyAccessor) Get(path Path) (Accessor, error) {
	return nil, fmt.Errorf("this is dummy accessor: %d", o.ID)
}

func (o DummyAccessor) Set(_ interface{}, path Path) error {
	return fmt.Errorf("this is dummy accessor: %d", o.ID)
}

func (o DummyAccessor) Unwrap() interface{} {
	return o.ID
}
