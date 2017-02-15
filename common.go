package undef

type pathPusher interface {
	PushPath(path string)
}

func getFromChild(child Object, path string, paths []string) (Object, error) {
	if len(paths) == 0 {
		return child, nil
	}

	r, err := child.Get(paths[0], paths[1:]...)
	if err != nil {
		if pe, ok := err.(pathPusher); ok {
			pe.PushPath(path)
		}
		return nil, err
	}
	return r, nil
}

// paths must not be empty.
func setToChild(child, obj Object, path string, paths []string) error {
	err := child.Set(obj, paths[0], paths[1:]...)
	if err != nil {
		if pe, ok := err.(pathPusher); ok {
			pe.PushPath(path)
		}
		return err
	}
	return nil
}
