package undef

import (
	"fmt"
	"strings"
)

func NewPathError(message, path string) error {
	return &PathError{message, path, nil}
}

type PathError struct {
	message string
	path    string
	stack   []string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("%s: about %q: at %s", e.message, e.path, strings.Join(e.stack, "/")+"/")
}

func (e *PathError) PushPath(path string) {
	e.stack = append(e.stack, path)
}
