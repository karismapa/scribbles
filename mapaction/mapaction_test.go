package mapaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSTS(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert := assert.New(t)

		key := "key"
		value := "value"

		m := New()
		m.SetSTS(map[string]string{key: value})

		hasMap := m.HasSTS()
		assert.True(hasMap)

		hasKeyWithValue := m.GetSTS(key, value)
		assert.True(hasKeyWithValue)

		doesntHasKeyWithValue := m.GetSTS(value, key)
		assert.False(doesntHasKeyWithValue)
	})
}
