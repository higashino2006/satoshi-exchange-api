package lib_test

import (
	"se-api/src/internal/lib/common"
	"testing"
)

func TestRoundDownFrom4DecimalPlaces(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{1.123456, 1.123},
		{3.987654, 3.987},
		{0.000123, 0.000},
		{10.999999, 10.999},
		{-1.987654, -1.987},
	}

	for _, test := range tests {
		result := common.RoundDownFrom4DecimalPlaces(test.input)
		if result != test.expected {
			t.Errorf("For input %f, expected %f, but got %f", test.input, test.expected, result)
		}
	}
}
