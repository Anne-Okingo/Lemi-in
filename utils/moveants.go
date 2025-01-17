package utils

import (
	"fmt"
)

type Path struct {
	Rooms []string
}

// MoveAnts generates a slice of moves representing the paths taken by each ant.
// Each move is formatted as "L<ant>-<room>", where <ant> is the ant's ID and <room> is the room the ant moves to.
func MoveAnts(paths []Path, antsPerRoom map[int][]int, totalTurns int) []string {
	moves := make([]string, totalTurns)

	for pathIndex, path := range paths {
		ants := antsPerRoom[pathIndex] // Ants assigned to this path
		for antIndex, ant := range ants {
			for turnOffset, room := range path.Rooms[1:] {
				moveIndex := antIndex + turnOffset
				if moveIndex >= totalTurns {
					break // Avoid out-of-bounds issues
				}

				// Use a single string concatenation / alternative to builder
				if moves[moveIndex] != "" {
					moves[moveIndex] += " "
				}
				moves[moveIndex] += fmt.Sprintf("L%d-%s", ant, room)
			}
		}
	}
	return moves
}
