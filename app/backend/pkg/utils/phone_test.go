package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToNationalFormat(t *testing.T) {
	tests := []struct {
		phone       string
		countryCode string
		expected    string
	}{
		{
			phone:       "+84818311855",
			countryCode: "VN",
			expected:    "0818311855",
		},
		{
			phone:       "+79843768500",
			countryCode: "RU",
			expected:    "8(984)376-85-00",
		},
	}
	for _, test := range tests {
		t.Run(test.phone, func(t *testing.T) {
			result, _ := ToNationalFormat(test.phone, test.countryCode)
			assert.Equal(t, test.expected, result)
		})
	}
}
