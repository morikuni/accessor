package undef

func Convert(obj interface{}) (Object, error) {
	switch t := obj.(type) {
	case map[string]interface{}:
		mo := map[string]Object{}
		for k, v := range t {
			o, err := Convert(v)
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
			o, err := Convert(v)
			if err != nil {
				return nil, err
			}
			mo[key] = o
		}
		return MapObject(mo), nil
	case []interface{}:
		so := make([]Object, len(t))
		for i, v := range t {
			o, err := Convert(v)
			if err != nil {
				return nil, err
			}
			so[i] = o
		}
		return SliceObject(so), nil
	default:
		return BasicTypes{t}, nil
	}
}
