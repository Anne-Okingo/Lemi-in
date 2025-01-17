package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"lemin/models"
)

func ParseInput(filename string) (*models.Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("invalid data format, error reading file: %v", err)
	}
	defer file.Close()

	graph := &models.Graph{
		Rooms: make(map[string]*models.ARoom),
	}
	
	scanner := bufio.NewScanner(file)
	isStartRoom := false
	isEndRoom := false
	startCount := 0
	endCount := 0
	tunnels := make(map[string]bool) // Track unique tunnels

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Skip comments that aren't commands
		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		// Parse number of ants (first non-comment line)
		if graph.AntCount == 0 {
			antCount, err := strconv.Atoi(line)
			if err != nil || antCount <= 0 {
				return nil, fmt.Errorf("invalid data format, invalid number of ants")
			}
			graph.AntCount = antCount
			continue
		}

		// Handle start/end commands
		if line == "##start" {
			isStartRoom = true
			startCount++
			if startCount > 1 {
				return nil, fmt.Errorf("invalid data format, multiple start rooms")
			}
			continue
		}
		if line == "##end" {
			isEndRoom = true
			endCount++
			if endCount > 1 {
				return nil, fmt.Errorf("invalid data format, multiple end rooms")
			}
			continue
		}

		// Parse Rooms
		parts := strings.Fields(line)
		if len(parts) == 3 {
			// Validate room name
			if strings.HasPrefix(parts[0], "L") || strings.HasPrefix(parts[0], "#") {
				return nil, fmt.Errorf("invalid data format, room name cannot start with 'L' or '#'")
			}
			if strings.Contains(parts[0], " ") {
				return nil, fmt.Errorf("invalid data format, room name cannot contain spaces")
			}
			
			// Check for duplicate rooms
			if _, exists := graph.Rooms[parts[0]]; exists {
				return nil, fmt.Errorf("invalid data format, duplicate room: %s", parts[0])
			}

			// Parse coordinates
			x, errX := strconv.Atoi(parts[1])
			y, errY := strconv.Atoi(parts[2])
			if errX != nil || errY != nil {
				return nil, fmt.Errorf("invalid data format, invalid coordinates for room: %s", parts[0])
			}

			room := &models.ARoom{
				Name:        parts[0],
				XCoordinate: x,
				YCoordinate: y,
				Links:       make([]string, 0),
			}

			if isStartRoom {
				graph.StartRoom = room.Name
				isStartRoom = false
			}
			if isEndRoom {
				graph.EndRoom = room.Name
				isEndRoom = false
			}
			
			graph.Rooms[room.Name] = room
			continue
		}

		// Parse Links
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid data format, invalid link format: %s", line)
			}

			from, to := parts[0], parts[1]
			
			// Validate rooms exist
			if _, ok := graph.Rooms[from]; !ok {
				return nil, fmt.Errorf("invalid data format, link references unknown room: %s", from)
			}
			if _, ok := graph.Rooms[to]; !ok {
				return nil, fmt.Errorf("invalid data format, link references unknown room: %s", to)
			}

			// Check for duplicate tunnels
			tunnelKey := fmt.Sprintf("%s-%s", min(from, to), max(from, to))
			if tunnels[tunnelKey] {
				return nil, fmt.Errorf("invalid data format, duplicate tunnel between rooms %s and %s", from, to)
			}
			tunnels[tunnelKey] = true

			graph.Rooms[from].Links = append(graph.Rooms[from].Links, to)
			graph.Rooms[to].Links = append(graph.Rooms[to].Links, from)
			continue
		}

		return nil, fmt.Errorf("invalid data format, invalid line: %s", line)
	}

	// Final validation
	if graph.StartRoom == "" {
		return nil, fmt.Errorf("invalid data format, no start room found")
	}
	if graph.EndRoom == "" {
		return nil, fmt.Errorf("invalid data format, no end room found")
	}
	if len(graph.Rooms) == 0 {
		return nil, fmt.Errorf("invalid data format, no rooms found")
	}

	return graph, nil
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}

func max(a, b string) string {
	if a > b {
		return a
	}
	return b
}