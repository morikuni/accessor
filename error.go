package undef

import (
	"fmt"
	"strings"
)

func NewPathError(message, path string) error {
	return &PathError{message, path, nil}
}

type PathError struct {
	Message string
	Path    string
	Stack   []string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("%s: about %q: at %s", e.Message, e.Path, strings.Join(e.Stack, "/")+"/")
}

func (e *PathError) PushPath(path string) {
	e.Stack = append(e.Stack, path)
}

func NewInvalidKeyError(v interface{}) error {
	return &InvalidKeyError{v}
}

type InvalidKeyError struct {
	Value interface{}
}

func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("%T is not a string", e.Value)
}
