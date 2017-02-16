package undef

import (
	"encoding"
	"encoding/json"
	"fmt"
)

type BasicTypes struct {
	Value interface{}
}

func (o BasicTypes) Get(path string, paths ...string) (Object, error) {
	return nil, NewPathError(fmt.Sprintf("%[1]T(%[1]v) has no key", o.Value), path)
}

func (o BasicTypes) Set(obj Object, path string, paths ...string) error {
	return NewPathError(fmt.Sprintf("%[1]T(%[1]v) has no key", o.Value), path)
}

func (o BasicTypes) Unwrap() interface{} {
	return o.Value
}

func (o BasicTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

func (o BasicTypes) MarshalText() ([]byte, error) {
	if m, ok := o.Value.(encoding.TextMarshaler); ok {
		return m.MarshalText()
	}
	if s, ok := o.Value.(fmt.Stringer); ok {
		return []byte(s.String()), nil
	}
	return []byte(fmt.Sprint(o.Value)), nil
}

func (o BasicTypes) MarshalYAML() (interface{}, error) {
	return o.Value, nil
}
