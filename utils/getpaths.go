package utils

import (
	"container/list"
	"lemin/models"
	"sort"
	"fmt"
)

func GetAllPaths(rooms map[string]*models.ARoom, start, end string) [][]string {
	var paths [][]string
	queue := list.New()
	queue.PushBack([]string{start})

	if _, exists := rooms[start]; !exists {
		return paths
	}

	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]string)
		lastRoom := path[len(path)-1]

		if lastRoom == end {
			paths = append(paths, path)
			continue
		}

		for _, nextRoom := range rooms[lastRoom].Links {
			if !contains(path, nextRoom) {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, nextRoom)
				queue.PushBack(newPath)
			}
		}
	}
	
	paths=SortPaths(paths)
	fmt.Println("unfiltered",len(paths))
	paths=FilterBestPaths(paths)
	fmt.Println("filtered",len(paths))
	return  paths
}

func contains(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}


// SortPaths sorts the paths by length in ascending order.
func SortPaths(paths [][]string) [][]string {
    sort.Slice(paths, func(i, j int) bool {
        return len(paths[i]) < len(paths[j])
    })
    return paths
}

// FilterBestPaths selects the most efficient set of paths for ant movement
func FilterBestPaths(allPaths [][]string) [][]string {
	// If no paths or only one path, return as is
	if len(allPaths) <= 1 {
		return allPaths
	}

	// Sort paths by length (shortest first)
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	// Initialize best paths with the shortest paths
	bestPaths := [][]string{allPaths[0]}

	// Iterate through remaining paths
	for _, currentPath := range allPaths[1:] {
		// Check for conflicts with existing best paths
		if isPathCompatible(bestPaths, currentPath) {
			bestPaths = append(bestPaths, currentPath)
		}
	}

	return bestPaths
}

// isPathCompatible checks if a path can be added without room conflicts
func isPathCompatible(existingPaths [][]string, newPath []string) bool {
	// Always allow start and end rooms to be shared
	startRoom := newPath[0]
	endRoom := newPath[len(newPath)-1]

	// Check each existing path
	for _, existingPath := range existingPaths {
		// Compare intermediate rooms
		for _, newRoom := range newPath[1 : len(newPath)-1] {
			for _, existingRoom := range existingPath[1 : len(existingPath)-1] {
				// If an intermediate room is the same, paths conflict
				if newRoom == existingRoom && 
				   newRoom != startRoom && 
				   newRoom != endRoom {
					return false
				}
			}
		}
	}
	return true
}
