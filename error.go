package accessor

import (
	"fmt"
	"strings"
)

func NewNoSuchPathError(message string, path Path) error {
	return &NoSuchPathError{message, path.Key(), nil}
}

type NoSuchPathError struct {
	Message string
	Key     string
	Stack   []string
}

func (e *NoSuchPathError) Error() string {
	return fmt.Sprintf("%s: about %q: at %s", e.Message, e.Key, strings.Join(e.Stack, "/")+"/")
}

func (e *NoSuchPathError) PushPath(path string) {
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

func NewInvalidPathError(path string) error {
	return &InvalidPathError{path}
}

type InvalidPathError struct {
	Path string
}

func (e *InvalidPathError) Error() string {
	return fmt.Sprintf("path is invalid: %q", e.Path)
}
