package accessor

import (
	"strconv"
)

type SliceAccessor []Accessor

func (a SliceAccessor) Get(path Path) (Accessor, error) {
	i, err := strconv.Atoi(path.Key())
	if err != nil {
		return nil, NewNoSuchPathError("not a number", path)
	}

	if i < 0 || i >= len(a) {
		return nil, NewNoSuchPathError("index out of range", path)
	}

	return getFromChild(a[i], path)
}

func (a SliceAccessor) Set(path Path, value interface{}) error {
	i, err := strconv.Atoi(path.Key())
	if err != nil {
		return NewNoSuchPathError("not a number", path)
	}

	if i < 0 || i >= len(a) {
		return NewNoSuchPathError("index out of range", path)
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

func (a SliceAccessor) Unwrap() interface{} {
	result := make([]interface{}, len(a))
	for i, v := range a {
		result[i] = v.Unwrap()
	}
	return result
}

func (a SliceAccessor) Foreach(f func(path Path, value interface{}) error) error {
	for i, child := range a {
		err := foreach(child, strconv.Itoa(i), f)
		if err != nil {
			return err
		}
	}
	return nil
}
