package accessor

import (
	"fmt"
)

// NewNoSuchPathError creates a NoSuchPathError.
func NewNoSuchPathError(message string, key string, keys ...string) *NoSuchPathError {
	path := PhantomPath
	for _, k := range keys {
		path = path.PushKey(k)
	}
	return &NoSuchPathError{message, key, path}
}

// NoSuchPathError is returned when no key was found in the object.
type NoSuchPathError struct {
	Message string
	Key     string
	Path    Path
}

func (e *NoSuchPathError) Error() string {
	return fmt.Sprintf("%s: about %q: at %s", e.Message, e.Key, e.Path)
}

// PushKey push key to the head of the stack trace.
func (e *NoSuchPathError) PushKey(key string) {
	e.Path = e.Path.PushKey(key)
}

// NewInvalidKeyError createa InvalidKeyError.
func NewInvalidKeyError(v interface{}) error {
	return &InvalidKeyError{v}
}

// InvalidKeyError is returned when a key type of the map was not a string.
type InvalidKeyError struct {
	Value interface{}
}

func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("%T is not a string", e.Value)
}

// NewInvalidPathError creates a InvalidKeyError.
func NewInvalidPathError(message string) error {
	return &InvalidPathError{message}
}

// InvalidPathError is returned when failing to create a Path.
type InvalidPathError struct {
	Message string
}

func (e *InvalidPathError) Error() string {
	return fmt.Sprintf("path is invalid: %s", e.Message)
}
