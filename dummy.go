package undef

import (
	"fmt"
)

type DummyObject struct {
	ID int
}

func (o DummyObject) Get(path string, paths ...string) (Object, error) {
	return nil, fmt.Errorf("this is dummy object: %d", o.ID)
}

func (o DummyObject) Set(obj Object, path string, paths ...string) error {
	return fmt.Errorf("this is dummy object: %d", o.ID)
}

func (o DummyObject) Unwrap() interface{} {
	return o.ID
}
