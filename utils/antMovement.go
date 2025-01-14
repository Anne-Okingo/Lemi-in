package utils

import (
	"fmt"
	"strings"
)

type AntPath struct {
	pathIndex int    // Index of the path being used
	ant       int    // Ant number
	position  int    // Current position in path
	path      []string // The actual path
}

func SimulateAntMovement(paths [][]string, antCount int) {
	// Initialize path assignments
	pathAssignments := make(map[int][]int) // map[pathIndex][]antNumbers
	
	// For each ant, find the optimal path
	for ant := 1; ant <= antCount; ant++ {
		bestPathIndex := findBestPath(paths, pathAssignments)
		if bestPathIndex == -1 {
			return // Error case
		}
		
		// Assign ant to the best path
		pathAssignments[bestPathIndex] = append(pathAssignments[bestPathIndex], ant)
	}

	// Simulate movements
	moveAnts(paths, pathAssignments)
}

func findBestPath(paths [][]string, pathAssignments map[int][]int) int {
	bestPathIndex := 0
	lowestCost := -1

	for i, path := range paths {
		// Calculate cost: rooms in path + ants already assigned
		currentAnts := len(pathAssignments[i])
		cost := len(path) - 1 + currentAnts // -1 because we don't count start room
		
		// Initialize lowestCost or update if this path is better
		if lowestCost == -1 || cost < lowestCost {
			lowestCost = cost
			bestPathIndex = i
		}
	}

	return bestPathIndex
}

func moveAnts(paths [][]string, pathAssignments map[int][]int) {
	// Create initial ant positions
	var activeAnts []AntPath
	for pathIndex, ants := range pathAssignments {
		for _, ant := range ants {
			activeAnts = append(activeAnts, AntPath{
				pathIndex: pathIndex,
				ant:       ant,
				position:  0,
				path:      paths[pathIndex],
			})
		}
	}

	// Continue until all ants reach the end
	for len(activeAnts) > 0 {
		var moves []string
		var remainingAnts []AntPath

		// Track room and link occupancy
		roomOccupancy := make(map[string]bool)
		linkUsage := make(map[string]bool)

		// Process each active ant
		for _, ant := range activeAnts {
			// Check if next room is available
			nextPos := ant.position + 1
			if nextPos >= len(ant.path) {
				continue // Ant has reached the end
			}

			currentRoom := ant.path[ant.position]
			nextRoom := ant.path[nextPos]
			link := fmt.Sprintf("%s-%s", currentRoom, nextRoom)

			// Skip if room is occupied or link is in use (except for end room)
			if roomOccupancy[nextRoom] && nextRoom != ant.path[len(ant.path)-1] {
				remainingAnts = append(remainingAnts, ant)
				continue
			}
			if linkUsage[link] {
				remainingAnts = append(remainingAnts, ant)
				continue
			}

			// Move ant
			moves = append(moves, fmt.Sprintf("L%d-%s", ant.ant, nextRoom))
			roomOccupancy[nextRoom] = true
			linkUsage[link] = true

			// If ant hasn't reached end, keep it active
			if nextPos < len(ant.path)-1 {
				remainingAnts = append(remainingAnts, AntPath{
					pathIndex: ant.pathIndex,
					ant:       ant.ant,
					position:  nextPos,
					path:      ant.path,
				})
			}
		}

		// Output moves for this turn
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}

		activeAnts = remainingAnts
	}
}
