package accessor

import (
	"reflect"
)

// Accessor provides getter/setter to the object.
type Accessor interface {
	// Get finds a object at specific path.
	// NoSuchPathError is returned when no object was found in the path.
	Get(path Path) (Accessor, error)

	// Set set a object into specific path.
	// NoSuchPathError is returned when the path is invalid.
	Set(path Path, value interface{}) error

	// Unwrap unwraps the object and returns actual value.
	Unwrap() interface{}
}

func NewAccessor(acc interface{}) (Accessor, error) {
	if a, ok := acc.(Accessor); ok {
		return a, nil
	}

	rv := reflect.ValueOf(acc)
	switch rv.Kind() {
	case reflect.Map:
		ma := map[string]Accessor{}
		for _, k := range rv.MapKeys() {
			key, ok := k.Interface().(string)
			if !ok {
				return nil, NewInvalidKeyError(k.Interface())
			}
			o, err := NewAccessor(rv.MapIndex(k).Interface())
			if err != nil {
				return nil, err
			}
			ma[key] = o
		}
		return MapAccessor(ma), nil
	case reflect.Slice:
		sa := make([]Accessor, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			o, err := NewAccessor(rv.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			sa[i] = o
		}
		return SliceAccessor(sa), nil
	default:
		return ValueAccessor{acc}, nil
	}
}
