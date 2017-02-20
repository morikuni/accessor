package accessor

type pathPusher interface {
	PushPath(path string)
}

func getFromChild(child Accessor, path Path) (Accessor, error) {
	subPath, ok := path.SubPath()
	if !ok {
		return child, nil
	}

	r, err := child.Get(subPath)
	if err != nil {
		if pe, ok := err.(pathPusher); ok {
			pe.PushPath(path.Key())
		}
		return nil, err
	}
	return r, nil
}

func setToChild(child Accessor, value interface{}, key string, path Path) error {
	err := child.Set(path, value)
	if err != nil {
		if pe, ok := err.(pathPusher); ok {
			pe.PushPath(key)
		}
		return err
	}
	return nil
}

func foreach(child Accessor, key string, f func(Path, interface{}) error) error {
	return child.Foreach(func(path Path, v interface{}) error {
		p := path.PushHead(key)
		return f(p, v)
	})
}
