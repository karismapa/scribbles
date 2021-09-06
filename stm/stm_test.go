package stm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToStringToSliceOfString(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		correctResult map[string][]string
	}{
		{
			"Parse string to map of string to slice of string",
			"key1:val1;key2:val2",
			map[string][]string{
				"key1": {"val1"},
				"key2": {"val2"},
				"key3": {"val3"},
			},
		},
		{
			"Parse string to map of string to slice of string",
			"key1:val1,val2,val3",
			map[string][]string{
				"key1": {"val1", "val2", "val3"},
			},
		},
		{
			"Parse string to map of string to slice of string",
			"key1:val1,val2;key2:val3,val4;key3:val5,val6",
			map[string][]string{
				"key1": {"val1", "val2"},
				"key2": {"val3", "val4"},
				"key3": {"val5", "val6"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert := assert.New(t)

			m := StringToStringToSliceOfString(testCase.input)

			for key, value := range m {
				assert.Equal(testCase.correctResult[key], value)
			}
		})
	}
}
