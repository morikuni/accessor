package undef

import (
	"encoding"
	"encoding/json"
	"fmt"
)

type ValueObject struct {
	Value interface{}
}

func (o ValueObject) Get(path string, paths ...string) (Object, error) {
	return nil, NewNoSuchPathError(fmt.Sprintf("%[1]T(%[1]v) has no key", o.Value), path)
}

func (o ValueObject) Set(obj Object, path string, paths ...string) error {
	return NewNoSuchPathError(fmt.Sprintf("%[1]T(%[1]v) has no key", o.Value), path)
}

func (o ValueObject) Unwrap() interface{} {
	return o.Value
}

func (o ValueObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

func (o ValueObject) MarshalText() ([]byte, error) {
	if m, ok := o.Value.(encoding.TextMarshaler); ok {
		return m.MarshalText()
	}
	if s, ok := o.Value.(fmt.Stringer); ok {
		return []byte(s.String()), nil
	}
	return []byte(fmt.Sprint(o.Value)), nil
}

func (o ValueObject) MarshalYAML() (interface{}, error) {
	return o.Value, nil
}
