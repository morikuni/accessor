package accessor

// MapAccessor is the Accessor for a map.
type MapAccessor map[string]Accessor

// Get implements Accessor.
func (a MapAccessor) Get(path Path) (Accessor, error) {
	child, ok := a[path.Key()]
	if !ok {
		return nil, NewNoSuchPathError("no such key", path.Key())
	}

	return getFromChild(child, path)
}

// Set implements Accessor.
func (a MapAccessor) Set(path Path, value interface{}) error {
	child, ok := a[path.Key()]
	if !ok {
		return NewNoSuchPathError("no such key", path.Key())
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

// Unwrap implements Accessor.
func (a MapAccessor) Unwrap() interface{} {
	result := map[string]interface{}{}
	for k, v := range a {
		result[k] = v.Unwrap()
	}
	return result
}

// Foreach implements Accessor.
func (a MapAccessor) Foreach(f func(path Path, value interface{}) error) error {
	for k, child := range a {
		err := foreach(child, k, f)
		if err != nil {
			return err
		}
	}
	return nil
}
