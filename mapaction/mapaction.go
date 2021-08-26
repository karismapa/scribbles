package mapaction

type MapAction interface {
	GetSTS(key, value string) bool
	HasSTS() bool
	SetSTS(mp map[string]string)
}

type mapAction struct {
	sts map[string]string
}

func (m *mapAction) GetSTS(key, value string) bool {
	val := m.sts[key]
	if val == value {
		return true
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

func New() MapAction {
	return &mapAction{}
}
