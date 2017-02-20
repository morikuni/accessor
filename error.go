package accessor

import (
	"fmt"
)

func NewNoSuchPathError(message string, key string, keys ...string) *NoSuchPathError {
	var path Path = initialPath{}
	for _, k := range keys {
		path = path.PushKey(k)
	}
	return &NoSuchPathError{message, key, path}
}

type NoSuchPathError struct {
	Message string
	Key     string
	Path    Path
}

func (e *NoSuchPathError) Error() string {
	return fmt.Sprintf("%s: about %q: at %s", e.Message, e.Key, e.Path)
}

func (e *NoSuchPathError) PushKey(key string) {
	e.Path = e.Path.PushKey(key)
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

func NewInvalidPathError(message string) error {
	return &InvalidPathError{message}
}

type InvalidPathError struct {
	Message string
}

func (e *InvalidPathError) Error() string {
	return fmt.Sprintf("path is invalid: %s", e.Message)
}
