package undef

import (
	"strconv"
)

type SliceObject []Object

func (o SliceObject) Get(path string, paths ...string) (Object, error) {
	i, err := strconv.Atoi(path)
	if err != nil {
		return nil, NewPathError("not a number", path)
	}

	if i < 0 || i >= len(paths) {
		return nil, NewPathError("index out of range", path)
	}

	return getFromChild(o[i], path, paths)
}

func (o SliceObject) Set(obj Object, path string, paths ...string) error {
	i, err := strconv.Atoi(path)
	if err != nil {
		return NewPathError("not a number", path)
	}

	if i < 0 || i >= len(paths) {
		return NewPathError("index out of range", path)
	}

	if len(paths) == 0 {
		o[i] = obj
		return nil
	}

	return setToChild(o[i], path, paths)
}
