package utils

import (
	"lemin/models"
	"testing"
)

func TestGetturns(t *testing.T) {
	tests := []struct {
		name          string
		pathwithants  map[int][]int
		graph         *models.Graph
		expectedTurns int
	}{
		{
			name: "Single path, single ant",
			pathwithants: map[int][]int{
				0: {1},
			},
			graph: &models.Graph{
				AllPaths: [][]string{
					{"start", "room1", "end"}, // Path 0: 2 rooms
				},
			},
			expectedTurns: 2, // 2 rooms + 1 ant - 1 = 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Getturns(tt.pathwithants, tt.graph)
			if got != tt.expectedTurns {
				t.Errorf("Getturns() = %v, want %v", got, tt.expectedTurns)
			}
		})
	}
}
