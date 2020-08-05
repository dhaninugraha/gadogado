package gadogado

type dummyStruct struct{}

type dummyMap map[string]dummyStruct

func newDummyMap() *dummyMap {
	m := make(dummyMap)
	return &m
}

func (m *dummyMap) addKey(k string) {
	if k != "" {
		(*m)[k] = dummyStruct{}
	}
}

func (m *dummyMap) exists(k string) bool {
	if m == nil || k == "" {
		return false
	}

	if m != nil {
		_, ok := (*m)[k]
		if ok {
			return true
		} else {
			return false
		}
	}

	return false
}
