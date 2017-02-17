package accessor

type MapAccessor map[string]Accessor

func (a MapAccessor) Get(path Path) (Accessor, error) {
	child, ok := a[path.Key()]
	if !ok {
		return nil, NewNoSuchPathError("no such key", path)
	}

	return getFromChild(child, path)
}

func (a MapAccessor) Set(path Path, value interface{}) error {
	child, ok := a[path.Key()]
	if !ok {
		return NewNoSuchPathError("no such key", path)
	}

	sub, ok := path.SubPath()
	if !ok {
		acc, err := NewAccessor(value)
		if err != nil {
			return err
		}
		a[path.Key()] = acc
		return nil
	}

	return setToChild(child, value, path.Key(), sub)
}

func (a MapAccessor) Unwrap() interface{} {
	result := map[string]interface{}{}
	for k, v := range a {
		result[k] = v.Unwrap()
	}
	return result
}
