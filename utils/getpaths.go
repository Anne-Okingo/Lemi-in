package utils

import (
	// "container/list"
	"fmt"
	"lemin/models"
)

// Gets all paths from parsed Input then filters out the ones that conflict to minimize traffic
func GetAllPaths(rooms map[string]*models.ARoom, start, end string) [][]string {
	var paths [][]string
	queue := [][]string{}                  // Queue to hold paths
	queue = append(queue, []string{start}) // Initialize the queue with the start room

	if _, exists := rooms[start]; !exists {
		return paths // Return early if the start room doesn't exist
	}

	// BFS loop
	for len(queue) > 0 {
		// Dequeue the first path from the queue
		currentPath := queue[0]
		queue = queue[1:]

		// Get the last room in the current path
		currentRoom := currentPath[len(currentPath)-1]

		// If the last room is the end room, add the path to the result
		if currentRoom == end {
			paths = append(paths, currentPath)
			continue
		}

		// Explore neighbors of the current room
		for _, nextRoom := range rooms[currentRoom].Links {
			// Avoid revisiting rooms already in the current path
			if !Contains(currentPath, nextRoom) {
				// Create a new path by appending the next room
				newPath := append([]string(nil), currentPath...)
				newPath = append(newPath, nextRoom)

				// Add the new path to the queue
				queue = append(queue, newPath)
			}
		}
	}

	//paths=SortPaths(paths)
	fmt.Println("unfiltered", len(paths))
	paths = FilterBestPaths(paths, start, end)
	fmt.Println("filtered", len(paths))
	return paths
}

func Contains(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

// FilterBestPaths selects the most efficient set of paths for ant movement that doesn't contain conflicts between rooms.
func FilterBestPaths(allPaths [][]string, start string, end string) [][]string {
	bestSolution := [][]string{}

	// Iterate through all paths to find the best solution.
	for i := 0; i < len(allPaths); i++ {
		currentSolution := [][]string{allPaths[i]}
		for j := 0; j < len(allPaths); j++ {
			// Ensure paths are not compared with themselves and check for conflicts.
			if i != j && canAddPath(currentSolution, allPaths[j], start, end) {
				currentSolution = append(currentSolution, allPaths[j])
			}
		}

		// Update the best solution if the current one is longer (better).
		if len(currentSolution) > len(bestSolution) {
			bestSolution = currentSolution
		}
	}

	// Return the best solution found.
	return bestSolution
}

// canAddPath checks if the current path can be added to the existing solution without any conflicts, ensuring no repeated rooms except for start and end.
func canAddPath(paths [][]string, candidate []string, start string, end string) bool {
	// Iterate through the existing paths to check for conflicts.
	for _, path := range paths {
		for _, room := range path {
			// Skip start and end rooms.
			if room == start || room == end {
				continue
			}

			// Check if the room already exists in the candidate path (conflict).
			for _, candidateRoom := range candidate {
				if room == candidateRoom {
					return false // Conflict found, return false.
				}
			}
		}
	}
	// No conflicts found, return true.
	return true
}
