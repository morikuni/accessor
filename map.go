package undef

type MapObject map[string]Object

func (o MapObject) Get(path string, paths ...string) (Object, error) {
	child, ok := o[path]
	if !ok {
		return nil, NewPathError("no such key", path)
	}

	return getFromChild(child, path, paths)
}

func (o MapObject) Set(obj Object, path string, paths ...string) error {
	child, ok := o[path]
	if !ok {
		return NewPathError("no such key", path)
	}

	if len(paths) == 0 {
		o[path] = obj
		return nil
	}

	return setToChild(child, obj, path, paths)
}
