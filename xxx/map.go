package xxx

import "reflect"

// MapGet ...
func MapGet(m interface{}, key interface{}) (interface{}, bool) {
	switch m.(type) {
	case map[string]interface{}:
		k, ok := key.(string)
		if !ok {
			return nil, false
		}
		r, ok := m.(map[string]interface{})[k]
		return r, ok
	case map[interface{}]interface{}:
		r, ok := m.(map[interface{}]interface{})[key]
		return r, ok
	default:
		val := reflect.ValueOf(m)
		if val.Kind() != reflect.Map {
			return nil, false
		}
		kVal := reflect.ValueOf(key)
		if !kVal.Type().AssignableTo(val.Type().Key()) {
			return nil, false
		}

		vVal := val.MapIndex(kVal)
		if !vVal.IsValid() {
			return nil, false
		}
		return vVal.Interface(), true
	}
}
