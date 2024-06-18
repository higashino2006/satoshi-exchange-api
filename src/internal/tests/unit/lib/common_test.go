package lib_test

import (
	"se-api/src/internal/lib/common"
	"testing"
)

func TestRoundDownFrom4DecimalPlaces(t *testing.T) {
	tests := []struct {
		input    float32
		expected float32
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

func TestJoinPaths(t *testing.T) {
	tests := []struct {
		baseURL  string
		subPath  string
		expected string
	}{
		{"http://example.com/", "path", "http://example.com/path"},
		{"http://example.com", "/path", "http://example.com/path"},
		{"http://example.com/", "/path", "http://example.com/path"},
		{"http://example.com", "path", "http://example.com/path"},
	}

	for _, test := range tests {
		result := common.JoinPaths(test.baseURL, test.subPath)
		if result != test.expected {
			t.Errorf("For baseURL %s and subPath %s, expected %s but got %s", test.baseURL, test.subPath, test.expected, result)
		}
	}
}
