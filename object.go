package accessor

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

func NewObject(obj interface{}) (Object, error) {
	switch t := obj.(type) {
	case map[string]interface{}:
		mo := map[string]Object{}
		for k, v := range t {
			o, err := NewObject(v)
			if err != nil {
				return nil, err
			}
			mo[k] = o
		}
		return MapObject(mo), nil
	case map[interface{}]interface{}:
		mo := map[string]Object{}
		for k, v := range t {
			key, ok := k.(string)
			if !ok {
				return nil, NewInvalidKeyError(k)
			}
			o, err := NewObject(v)
			if err != nil {
				return nil, err
			}
			mo[key] = o
		}
		return MapObject(mo), nil
	case []interface{}:
		so := make([]Object, len(t))
		for i, v := range t {
			o, err := NewObject(v)
			if err != nil {
				return nil, err
			}
			so[i] = o
		}
		return SliceObject(so), nil
	default:
		return ValueObject{t}, nil
	}
}
