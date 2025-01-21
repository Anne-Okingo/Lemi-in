package utils

import (
	"container/list"
	"lemin/models"
)

func GetAllPaths(rooms map[string]*models.ARoom, start, end string) [][]string {
	var paths [][]string
	queue := list.New()
	queue.PushBack([]string{start})

	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]string)
		lastRoom := path[len(path)-1]

		if lastRoom == end {
			paths = append(paths, path)
			continue
		}

		for _, nextRoom := range rooms[lastRoom].Links {
			if !Contains(path, nextRoom) {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, nextRoom)
				queue.PushBack(newPath)
			}
		}
	}
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
