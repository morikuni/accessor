package accessor

// Accessor provides getter/setter to the object.
type Accessor interface {
	// Get finds a object at specific path.
	// NoSuchPathError is returned when no object was found in the path.
	Get(path string, paths ...string) (Accessor, error)

	// Set set a object into specific path.
	// NoSuchPathError is returned when the path is invalid.
	Set(value interface{}, path string, paths ...string) error

	// Unwrap unwraps the object and returns actual value.
	Unwrap() interface{}
}

func NewAccessor(acc interface{}) (Accessor, error) {
	if a, ok := acc.(Accessor); ok {
		return a, nil
	}

	switch t := acc.(type) {
	case map[string]interface{}:
		ma := map[string]Accessor{}
		for k, v := range t {
			o, err := NewAccessor(v)
			if err != nil {
				return nil, err
			}
			ma[k] = o
		}
		return MapAccessor(ma), nil
	case map[interface{}]interface{}:
		ma := map[string]Accessor{}
		for k, v := range t {
			key, ok := k.(string)
			if !ok {
				return nil, NewInvalidKeyError(k)
			}
			o, err := NewAccessor(v)
			if err != nil {
				return nil, err
			}
			ma[key] = o
		}
		return MapAccessor(ma), nil
	case []interface{}:
		sa := make([]Accessor, len(t))
		for i, v := range t {
			o, err := NewAccessor(v)
			if err != nil {
				return nil, err
			}
			sa[i] = o
		}
		return SliceAccessor(sa), nil
	default:
		return ValueAccessor{t}, nil
	}
}
