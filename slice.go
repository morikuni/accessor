package accessor

import (
	"strconv"
)

// SliceAccessor is the Accessor for a slice.
type SliceAccessor []Accessor

// Get implements Accessor.
func (a SliceAccessor) Get(path Path) (Accessor, error) {
	i, err := strconv.Atoi(path.Key())
	if err != nil {
		return nil, NewNoSuchPathError("not a number", path.Key())
	}

	if i < 0 || i >= len(a) {
		return nil, NewNoSuchPathError("index out of range", path.Key())
	}

	return getFromChild(a[i], path)
}

// Set implements Accessor.
func (a SliceAccessor) Set(path Path, value interface{}) error {
	i, err := strconv.Atoi(path.Key())
	if err != nil {
		return NewNoSuchPathError("not a number", path.Key())
	}

	if i < 0 || i >= len(a) {
		return NewNoSuchPathError("index out of range", path.Key())
	}

	sub, ok := path.SubPath()
	if !ok {
		acc, err := NewAccessor(value)
		if err != nil {
			return err
		}
		a[i] = acc
		return nil
	}

	return setToChild(a[i], value, path.Key(), sub)
}

// Unwrap implements Accessor.
func (a SliceAccessor) Unwrap() interface{} {
	result := make([]interface{}, len(a))
	for i, v := range a {
		result[i] = v.Unwrap()
	}
	return result
}

// Foreach implements Accessor.
func (a SliceAccessor) Foreach(f func(path Path, value interface{}) error) error {
	for i, child := range a {
		err := foreach(child, strconv.Itoa(i), f)
		if err != nil {
			return err
		}
	}
	return nil
}
