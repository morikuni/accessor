package accessor

import (
	"strconv"
)

type SliceAccessor []Accessor

func (a SliceAccessor) Get(path string, paths ...string) (Accessor, error) {
	i, err := strconv.Atoi(path)
	if err != nil {
		return nil, NewNoSuchPathError("not a number", path)
	}

	if i < 0 || i >= len(a) {
		return nil, NewNoSuchPathError("index out of range", path)
	}

	return getFromChild(a[i], path, paths)
}

func (a SliceAccessor) Set(acc Accessor, path string, paths ...string) error {
	i, err := strconv.Atoi(path)
	if err != nil {
		return NewNoSuchPathError("not a number", path)
	}

	if i < 0 || i >= len(a) {
		return NewNoSuchPathError("index out of range", path)
	}

	if len(paths) == 0 {
		a[i] = acc
		return nil
	}

	return setToChild(a[i], acc, path, paths)
}

func (a SliceAccessor) Unwrap() interface{} {
	result := make([]interface{}, len(a))
	for i, v := range a {
		result[i] = v.Unwrap()
	}
	return result
}
