package accessor

import (
	"encoding"
	"encoding/json"
	"fmt"
)

type ValueAccessor struct {
	Value interface{}
}

func (a ValueAccessor) Get(path Path) (Accessor, error) {
	return nil, NewNoSuchPathError(fmt.Sprintf("%[1]T(%[1]v) has no key", a.Value), path.Key())
}

func (a ValueAccessor) Set(path Path, _ interface{}) error {
	return NewNoSuchPathError(fmt.Sprintf("%[1]T(%[1]v) has no key", a.Value), path.Key())
}

func (a ValueAccessor) Unwrap() interface{} {
	return a.Value
}

func (a ValueAccessor) Foreach(f func(path Path, value interface{}) error) error {
	return f(initialPath{}, a.Value)
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
