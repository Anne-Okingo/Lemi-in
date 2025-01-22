package test

import (
	"testing"

	"lemin/utils"
)

func TestIsValidPath(t *testing.T) {
	tests := []struct {
		name     string
		path     []string
		expected bool
	}{
		{
			name:     "Valid path",
			path:     []string{"A", "B", "C"},
			expected: true,
		},
		{
			name:     "Path with duplicates",
			path:     []string{"A", "B", "A"},
			expected: true,
		},
		{
			name:     "Single room path",
			path:     []string{"A"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsValidPath(tt.path)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
