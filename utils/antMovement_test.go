package utils

import (
	"testing"
)

func TestDistributeAnts(t *testing.T) {
	tests := []struct {
		name          string
		paths         [][]string
		totalAnts     int
		expectedPaths []PathInfo
	}{
		{
			name:      "Simple test with 3 paths and 6 ants",
			paths:     [][]string{{"A", "B", "C"}, {"A", "D", "E"}, {"F", "G", "H"}},
			totalAnts: 6,
			expectedPaths: []PathInfo{
				{Path: []string{"A", "B", "C"}, Length: 2, AntsUsing: 2},
				{Path: []string{"A", "D", "E"}, Length: 2, AntsUsing: 2},
				{Path: []string{"F", "G", "H"}, Length: 2, AntsUsing: 2},
			},
		},
		{
			name:      "Test with 2 paths and 5 ants",
			paths:     [][]string{{"A", "B", "C"}, {"D", "E", "F"}},
			totalAnts: 5,
			expectedPaths: []PathInfo{
				{Path: []string{"A", "B", "C"}, Length: 2, AntsUsing: 2},
				{Path: []string{"D", "E", "F"}, Length: 2, AntsUsing: 3},
			},
		},
		{
			name:          "Edge case with 0 paths",
			paths:         [][]string{},
			totalAnts:     5,
			expectedPaths: nil,
		},

		{
			name:      "Edge case with 1 path and 5 ants",
			paths:     [][]string{{"A", "B", "C", "D"}},
			totalAnts: 5,
			expectedPaths: []PathInfo{
				{Path: []string{"A", "B", "C", "D"}, Length: 3, AntsUsing: 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DistributeAnts(tt.paths, tt.totalAnts)

			// Compare result with expected paths
			if len(result) != len(tt.expectedPaths) {
				t.Errorf("expected %d paths, got %d", len(tt.expectedPaths), len(result))
			}

			for i, p := range result {
				if len(p.Path) != len(tt.expectedPaths[i].Path) || p.AntsUsing != tt.expectedPaths[i].AntsUsing {
					t.Errorf("for path %v: expected %v, got %v", p.Path, tt.expectedPaths[i], p)
				}
			}
		})
	}
}
