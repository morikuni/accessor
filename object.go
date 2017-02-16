package undef

// Object represents undefined object (including unstructured types like nil, int, string or something).
type Object interface {
	// Get finds a object at specific path.
	// NoSuchPathError is returned when no object was found in the path.
	Get(path string, paths ...string) (Object, error)

	// Set set a object into specific path.
	// NoSuchPathError is returned when the path is invalid.
	Set(obj Object, path string, paths ...string) error

	// Unwrap unwraps the object and returns actual value.
	Unwrap() interface{}
}
