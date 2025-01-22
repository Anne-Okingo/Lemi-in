package utils

import (
	"container/list"
	"lemin/models"
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

		// If we've reached the end room, add the path to the result
		if lastRoom == end {
			paths = append(paths, path)
			continue
		}

		// Track visited rooms for the current path
		visited := make(map[string]bool)
		for _, room := range path {
			visited[room] = true
		}

		// Explore neighboring rooms for the next potential path
		for _, nextRoom := range rooms[lastRoom].Links {
			// Allow revisiting 'start' and 'end' rooms, but not others
			if !visited[nextRoom] || nextRoom == start || nextRoom == end {
				// Create a new path by appending the next room
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, nextRoom)

				// Push the new path into the queue for further exploration
				queue.PushBack(newPath)
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
