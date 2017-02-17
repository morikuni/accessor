package accessor

func Get(i interface{}, path string) (interface{}, error) {
	if path == "/" {
		return i, nil
	}
	p, err := ParsePath(path)
	if err != nil {
		return nil, err
	}

	acc, err := NewAccessor(i)
	if err != nil {
		return nil, err
	}

	acc, err = acc.Get(p)
	if err != nil {
		return nil, err
	}

	return acc.Unwrap(), nil
}

func Update(i interface{}, path string, value interface{}) (interface{}, error) {
	if path == "/" {
		return nil, NewInvalidPathError(path)
	}
	p, err := ParsePath(path)
	if err != nil {
		return nil, err
	}

	acc, err := NewAccessor(i)
	if err != nil {
		return nil, err
	}

	err = acc.Set(value, p)
	if err != nil {
		return nil, err
	}
	return acc.Unwrap(), nil
}
