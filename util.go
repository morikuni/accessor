package accessor

import (
	"strings"
)

func Get(acc Accessor, path string) (Accessor, error) {
	if path == "/" {
		return acc, nil
	}

	p, ps, err := ParsePath(path)
	if err != nil {
		return nil, err
	}

	return acc.Get(p, ps...)
}

func Set(dst, acc Accessor, path string) error {
	if path == "/" {
		return nil
	}

	p, ps, err := ParsePath(path)
	if err != nil {
		return err
	}

	return dst.Set(acc, p, ps...)
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
