package thai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsThai(t *testing.T) {
	testCases := []struct {
		name   string
		term   string
		isThai bool
	}{
		{
			"All thai",
			"น้ำดื่ม",
			true,
		},
		{
			"Not thai",
			"drinking water",
			false,
		},
		{
			"Partially thai",
			"drinking น้ำ",
			true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert := assert.New(t)

			isThai := IsThai(testCase.term)
			assert.Equal(testCase.isThai, isThai)
		})
	}
}
