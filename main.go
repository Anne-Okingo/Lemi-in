package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lemin/models"
	"lemin/utils"
)

func ValidContent(filename string) ([]string, error) {
	fileContent, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	defer fileContent.Close()

	validContent := []string{}
	scanner := bufio.NewScanner(fileContent)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		validContent = append(validContent, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}

	return validContent, nil
}

func findPaths(startRoom string, endRoom string, rooms []models.Room, links []models.Link) [][]string {
	var paths [][]string
	visited := make(map[string]bool)
	currentPath := []string{startRoom}

	var dfs func(current string)
	dfs = func(current string) {
		if current == endRoom {
			newPath := make([]string, len(currentPath))
			copy(newPath, currentPath)
			paths = append(paths, newPath)
			return
		}

		visited[current] = true
		for _, link := range links {
			next := ""
			if link.From == current && !visited[link.To] {
				next = link.To
			} else if link.To == current && !visited[link.From] {
				next = link.From
			}
			if next != "" {
				currentPath = append(currentPath, next)
				dfs(next)
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
		visited[current] = false
	}

	dfs(startRoom)
	return paths
}

func moveAnts(ants int, paths [][]string) [][]string {
	if len(paths) == 0 {
		return nil
	}

	// Sort paths by length
	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}

	// Initialize ants
	type ant struct {
		id       int
		pathIdx  int
		position int
	}

	antList := make([]ant, ants)
	for i := range antList {
		antList[i] = ant{
			id:       i + 1,
			pathIdx:  0,
			position: -1,
		}
	}

	var moves [][]string
	for {
		finished := true
		var currentMoves []string

		for i := range antList {
			if antList[i].position < len(paths[0])-1 {
				finished = false
				antList[i].position++
				move := fmt.Sprintf("L%d-%s", antList[i].id, paths[0][antList[i].position])
				currentMoves = append(currentMoves, move)
			}
		}

		if finished {
			break
		}
		moves = append(moves, currentMoves)
	}

	return moves
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid data format, please provide a file name")
		return
	}

	// Parse input file
	ants, rooms, links, err := utils.ParseInput(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find start and end rooms
	var startRoom, endRoom string
	for _, room := range rooms {
		if room.IsStart {
			startRoom = room.Name
		}
		if room.IsEnd {
			endRoom = room.Name
		}
	}

	// Find all possible paths
	paths := findPaths(startRoom, endRoom, rooms, links)
	if len(paths) == 0 {
		fmt.Println("ERROR: no valid path found between start and end")
		return
	}

	// Print the input data
	content, err := utils.ValidContent(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	// Print number of ants and room configuration
	fmt.Println(content[0])
	for _, line := range content[1:] {
		if !strings.Contains(line, "-") {
			fmt.Println(line)
		}
	}

	// Print links
	for _, line := range content[1:] {
		if strings.Contains(line, "-") {
			fmt.Println(line)
		}
	}
	fmt.Println()

	// Move ants and print moves
	moves := moveAnts(ants, paths)
	for _, move := range moves {
		fmt.Println(strings.Join(move, " "))
	}
}
