package accessor

import (
	"strconv"
)

type SliceObject []Object

func (o SliceObject) Get(path string, paths ...string) (Object, error) {
	i, err := strconv.Atoi(path)
	if err != nil {
		return nil, NewNoSuchPathError("not a number", path)
	}

	if i < 0 || i >= len(o) {
		return nil, NewNoSuchPathError("index out of range", path)
	}

	return getFromChild(o[i], path, paths)
}

func (o SliceObject) Set(obj Object, path string, paths ...string) error {
	i, err := strconv.Atoi(path)
	if err != nil {
		return NewNoSuchPathError("not a number", path)
	}

	if i < 0 || i >= len(o) {
		return NewNoSuchPathError("index out of range", path)
	}

	if len(paths) == 0 {
		o[i] = obj
		return nil
	}

	return setToChild(o[i], obj, path, paths)
}

func (o SliceObject) Unwrap() interface{} {
	result := make([]interface{}, len(o))
	for i, v := range o {
		result[i] = v.Unwrap()
	}
	return result
}
