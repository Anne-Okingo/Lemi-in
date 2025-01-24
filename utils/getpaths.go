package utils

import (
	// "container/list"
	"lemin/models"
)

func GetAllPaths(rooms map[string]*models.ARoom, start, end string) [][]string {
	var paths [][]string
	queue := [][]string{} // Queue to hold paths
	queue = append(queue, []string{start}) // Initialize the queue with the start room

	if _, exists := rooms[start]; !exists {
		return paths// Return early if the start room doesn't exist
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

	// Return all found paths
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


// Check verifies if a path shares any rooms with already optimized paths (excluding start/end).
func Check(path []string, graph * models.Graph) bool {
	// Use a map for faster lookups
	visitedRooms := make(map[string]struct{})
	for _, optimizedpath := range graph.AllPaths {
		for _, room := range optimizedpath[1 : len(optimizedpath)-1] { // Ignore start and end rooms
			visitedRooms[room] = struct{}{}
		}
	}

	for _, room := range path[1 : len(path)-1] { // Ignore start and end rooms
		if _, found := visitedRooms[room]; found {
			return false
		}
	}
	return true
}


// OptimizedPaths1 filters paths that don't share rooms.
func OptimizedPaths1(graph *models.Graph) [][]string {
	optimized := [][]string{}
	for i := 1; i < len(graph.AllPaths); i++ {
		if Check(graph.AllPaths[i], graph) {
			optimized = append(optimized, graph.AllPaths[i])
		}
	}
	return optimized
}

func deletepath(graph *models.Graph, i int) [][]string {
	// Remove the path at index i from graph.AllPaths
	return append(graph.AllPaths[:i], graph.AllPaths[i+1:]...)
}