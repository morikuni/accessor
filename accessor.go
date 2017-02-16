package accessor

import (
	"strings"
)

func Get(obj Object, path string) (Object, error) {
	if path == "/" {
		return obj, nil
	}

	p, ps, err := ParsePath(path)
	if err != nil {
		return nil, err
	}

	return obj.Get(p, ps...)
}

func Set(dst, obj Object, path string) error {
	if path == "/" {
		return nil
	}

	p, ps, err := ParsePath(path)
	if err != nil {
		return err
	}

	return dst.Set(obj, p, ps...)
}

func ParsePath(path string) (string, []string, error) {
	paths := strings.Split(path, "/")
	if paths[0] == "" {
		paths = paths[1:]
	}
	if last := len(paths) - 1; paths[last] == "" {
		paths = paths[:last]
	}
	if len(paths) == 0 {
		return "", nil, NewInvalidPathError(path)
	}

	return paths[0], paths[1:], nil
}
