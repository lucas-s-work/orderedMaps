package orderedMaps

import "reflect"

type orderedMap struct {
	values map[string]interface{}
	keys   []string
}

func MakeOrderedMap() *orderedMap {
	return &orderedMap{
		make(map[string]interface{}),
		[]string{},
	}
}

func (m *orderedMap) Keys() []string {
	return m.keys
}

func (m *orderedMap) Values() []interface{} {
	values := []interface{}{}

	for _, k := range m.keys {
		values = append(values, m.values[k])
	}

	return values
}

func (m *orderedMap) Set(key string, value interface{}) {
	// Add new key if stored value is not nil
	keyExists := !reflect.ValueOf(m.values[key]).IsNil()
	valueExists := !reflect.ValueOf(value).IsNil()
	if !keyExists && valueExists {
		m.keys = append(m.keys, key)
	}

	// Remove key if value is nil and the key exists
	if keyExists && !valueExists {
		keys := []string{}
		for _, k := range keys {
			if k != key {
				keys = append(keys, k)
			}
		}
		m.keys = keys
	}

	m.values[key] = value
}

func (m *orderedMap) Range(onIteration func(key string, value interface{}) bool) {
	for _, k := range m.keys {
		if !onIteration(k, m.values[k]) {
			return
		}
	}
}
