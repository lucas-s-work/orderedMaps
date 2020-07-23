package orderedMaps

type OrderedMap struct {
	values map[string]interface{}
	keys   []string
}

func MakeOrderedMap() *OrderedMap {
	return &OrderedMap{
		make(map[string]interface{}),
		[]string{},
	}
}

func (m *OrderedMap) Keys() []string {
	return m.keys
}

func (m *OrderedMap) Values() []interface{} {
	values := []interface{}{}

	for _, k := range m.keys {
		values = append(values, m.values[k])
	}

	return values
}

func (m *OrderedMap) Set(key string, value interface{}) {
	// Check if the key already exists
	keyExists := false
	for _, k := range m.keys {
		if k == key {
			keyExists = true
			break
		}
	}

	if !keyExists {
		m.keys = append(m.keys, key)
	}

	m.values[key] = value
}

func (m *OrderedMap) Remove(key string) {
	m.values[key] = nil

	rIndex := -1
	for i, k := range m.keys {
		if k == key {
			rIndex = i
			break
		}
	}
	m.keys = append(m.keys[:rIndex], m.keys[rIndex+1:]...)
}

func (m *OrderedMap) Range(onIteration func(key string, value interface{}) bool) {
	for _, k := range m.keys {
		if !onIteration(k, m.values[k]) {
			return
		}
	}
}
