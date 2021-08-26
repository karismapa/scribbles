package mapaction

type MapAction interface {
	STSPairExist(key, value string) bool
	HasSTS() bool
	SetSTS(mp map[string]string)

	STASPairExist(key, value string) bool
	HasSTAS() bool
	SetSTAS(mp map[string][]string)
}

type mapAction struct {
	sts  map[string]string
	stas map[string][]string
}

func (m *mapAction) STSPairExist(key, value string) bool {
	if m.sts != nil {
		val := m.sts[key]
		if val == value {
			return true
		}
	}

	return false
}

func (m *mapAction) HasSTS() bool {
	if m.sts != nil {
		return true
	}

	return false
}

func (m *mapAction) SetSTS(mp map[string]string) {
	if mp != nil {
		m.sts = mp
	}
}

func (m *mapAction) STASPairExist(key, value string) bool {
	if m.stas != nil {
		vals := m.stas[key]
		if vals != nil {
			for _, val := range vals {
				if val == value {
					return true
				}
			}
		}
	}

	return false
}

func (m *mapAction) HasSTAS() bool {
	if m.stas != nil {
		return true
	}

	return false
}

func (m *mapAction) SetSTAS(mp map[string][]string) {
	if mp != nil {
		m.stas = mp
	}
}

func New() MapAction {
	return &mapAction{}
}
