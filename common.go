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
	err := child.Set(value, path)
	if err != nil {
		if pe, ok := err.(pathPusher); ok {
			pe.PushPath(key)
		}
		return err
	}
	return nil
}
