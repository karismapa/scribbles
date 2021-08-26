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

		hasKeyWithValue := m.STSPairExist(key, value)
		assert.True(hasKeyWithValue)

		doesntHasKeyWithValue := m.STSPairExist(value, key)
		assert.False(doesntHasKeyWithValue)
	})
}

func TestGetSTAS(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert := assert.New(t)

		key := "key"
		value1 := "value1"
		value2 := "value2"

		m := New()
		m.SetSTAS(map[string][]string{key: {value1, value2}})

		hasMap := m.HasSTAS()
		assert.True(hasMap)

		hasKeyWithValue := m.STASPairExist(key, value1)
		assert.True(hasKeyWithValue)

		doesntHasKeyWithValue := m.STASPairExist(value1, key)
		assert.False(doesntHasKeyWithValue)
	})
}
