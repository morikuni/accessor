package accessor

import (
	"encoding"
	"encoding/json"
	"fmt"
)

type ValueAccessor struct {
	Value interface{}
}

func (a ValueAccessor) Get(path string, paths ...string) (Accessor, error) {
	return nil, NewNoSuchPathError(fmt.Sprintf("%[1]T(%[1]v) has no key", a.Value), path)
}

func (a ValueAccessor) Set(_ interface{}, path string, paths ...string) error {
	return NewNoSuchPathError(fmt.Sprintf("%[1]T(%[1]v) has no key", a.Value), path)
}

func (a ValueAccessor) Unwrap() interface{} {
	return a.Value
}

func (a ValueAccessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Value)
}

func (a ValueAccessor) MarshalText() ([]byte, error) {
	if m, ok := a.Value.(encoding.TextMarshaler); ok {
		return m.MarshalText()
	}
	if s, ok := a.Value.(fmt.Stringer); ok {
		return []byte(s.String()), nil
	}
	return []byte(fmt.Sprint(a.Value)), nil
}

func (a ValueAccessor) MarshalYAML() (interface{}, error) {
	return a.Value, nil
}
