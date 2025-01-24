package utils

import (
	"container/list"
	"lemin/models"
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
	
	//paths=SortPaths(paths)
	fmt.Println("unfiltered",len(paths))
	paths=FilterBestPaths(paths,start,end)
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


// // SortPaths sorts the paths by length in ascending order.
// func SortPaths(paths [][]string) [][]string {
//     sort.Slice(paths, func(i, j int) bool {
//         return len(paths[i]) < len(paths[j])
//     })
//     return paths
// }

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

// canAddPath checks if the current path can be added to the existing solution
// without any conflicts, ensuring no repeated rooms except for start and end.
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
