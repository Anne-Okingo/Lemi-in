package utils

import (
	"fmt"
	"sort"
	"strings"
)

type Ant struct {
	ID       int
	Path     []string
	Position int
}

type PathInfo struct {
	Path      []string
	Length    int
	AntsUsing int
}

// Improved distribution algorithm focusing on parallel path usage
func DistributeAnts(paths [][]string, totalAnts int) []PathInfo {
	if len(paths) == 0 || totalAnts <= 0 {
		return nil
	}

	// Initialize path info
	pathInfos := make([]PathInfo, len(paths))
	for i, path := range paths {
		pathInfos[i] = PathInfo{
			Path:      path,
			Length:    len(path) - 1,
			AntsUsing: 0,
		}
	}

	// Sort paths by length in ascending order
	sort.Slice(pathInfos, func(i, j int) bool {
		return pathInfos[i].Length < pathInfos[j].Length
	})

	// Calculate initial distribution
	remainingAnts := totalAnts
	for remainingAnts > 0 {
		// Find the path that would finish earliest with one more ant
		bestIdx := 0
		bestTime := pathInfos[0].Length + pathInfos[0].AntsUsing + 1

		// Find the Best Path to Add one more Ant
		for i := range pathInfos {
			finishTime := pathInfos[i].Length + pathInfos[i].AntsUsing + 1
			if finishTime <= bestTime {
				bestTime = finishTime
				bestIdx = i
			}
		}

		// Assign Ant to Best Path
		pathInfos[bestIdx].AntsUsing++
		remainingAnts--
	}

	// Filter out unused paths
	result := make([]PathInfo, 0)
	for _, p := range pathInfos {
		if p.AntsUsing > 0 {
			result = append(result, p)
		}
	}
	return result
}

func simulateMovements(ants []*Ant) {
	antsByPath := make(map[string][]*Ant)
	// Group ants by their assigned paths
	for _, ant := range ants {
		pathKey := strings.Join(ant.Path, ",")
		antsByPath[pathKey] = append(antsByPath[pathKey], ant)
	}

	activeAnts := make([]*Ant, 0)
	antIndex := 0

	for {
		movements := make([]string, 0)
		occupiedRooms := make(map[string]bool)

		// Process existing active ants
		newActiveAnts := make([]*Ant, 0)
		for _, ant := range activeAnts {
			if ant.Position >= len(ant.Path)-1 {
				continue
			}

			nextRoom := ant.Path[ant.Position+1]
			if !occupiedRooms[nextRoom] || nextRoom == ant.Path[len(ant.Path)-1] {
				ant.Position++
				movements = append(movements, fmt.Sprintf("L%d-%s", ant.ID, nextRoom))
				if nextRoom != ant.Path[len(ant.Path)-1] {
					occupiedRooms[nextRoom] = true
				}
				if ant.Position < len(ant.Path)-1 {
					newActiveAnts = append(newActiveAnts, ant)
				}
			} else {
				newActiveAnts = append(newActiveAnts, ant)
			}
		}

		// Try to add new ants from each path
		for pathKey, pathAnts := range antsByPath {
			if antIndex >= len(pathAnts) {
				continue
			}

			path := strings.Split(pathKey, ",")
			nextRoom := path[1] // First room after start

			if !occupiedRooms[nextRoom] {
				ant := pathAnts[antIndex]
				ant.Position = 1
				movements = append(movements, fmt.Sprintf("L%d-%s", ant.ID, nextRoom))
				occupiedRooms[nextRoom] = true
				newActiveAnts = append(newActiveAnts, ant)
			}
		}

		if len(movements) == 0 {
			break
		}

		// Sort movements for consistent output
		sort.Slice(movements, func(i, j int) bool {
			return movements[i] < movements[j]
		})

		fmt.Println(strings.Join(movements, " "))
		activeAnts = newActiveAnts
		antIndex++
	}
}

func SimulateAntMovement(validPaths [][]string, antCount int) {
	
	// validPaths := filterValidPaths(paths)
	if len(validPaths) == 0 {
		fmt.Println("ERROR: No valid paths available")
		return
	}

	pathInfos := DistributeAnts(validPaths, antCount)

	// Create and assign ants
	var ants []*Ant
	currentAntID := 1

	// Create ants based on distribution
	for _, pathInfo := range pathInfos {
		for i := 0; i < pathInfo.AntsUsing; i++ {
			ants = append(ants, &Ant{
				ID:       currentAntID,
				Path:     pathInfo.Path,
				Position: 0,
			})
			currentAntID++
		}
	}

	simulateMovements(ants)
}