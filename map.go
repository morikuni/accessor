package accessor

type MapAccessor map[string]Accessor

func (a MapAccessor) Get(path string, paths ...string) (Accessor, error) {
	child, ok := a[path]
	if !ok {
		return nil, NewNoSuchPathError("no such key", path)
	}

	return getFromChild(child, path, paths)
}

func (a MapAccessor) Set(acc Accessor, path string, paths ...string) error {
	child, ok := a[path]
	if !ok {
		return NewNoSuchPathError("no such key", path)
	}

	if len(paths) == 0 {
		a[path] = acc
		return nil
	}

	return setToChild(child, acc, path, paths)
}

func (a MapAccessor) Unwrap() interface{} {
	result := map[string]interface{}{}
	for k, v := range a {
		result[k] = v.Unwrap()
	}
	return result
}
