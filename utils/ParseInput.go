package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lemin/models"
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

func ParseInput(filename string) (int, []models.Room, []models.Link, error) {
	contents, err := ValidContent(filename)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format: %v", err)
	}

	if len(contents) == 0 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format: empty file")
	}

	// Parse number of ants
	ants, err := strconv.Atoi(contents[0])
	if err != nil || ants <= 0 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid number of Ants")
	}

	var rooms []models.Room
	var links []models.Link
	var nextIsStart, nextIsEnd bool

	// Parse rooms and links
	for i := 1; i < len(contents); i++ {
		line := contents[i]

		// Handle commands
		if line == "##start" {
			nextIsStart = true
			continue
		}
		if line == "##end" {
			nextIsEnd = true
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Handle links
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) == 2 && parts[0] != "" && parts[1] != "" {
				links = append(links, models.Link{From: parts[0], To: parts[1]})
			}
			continue
		}

		// Handle rooms
		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue
		}

		// Validate coordinates
		x, errX := strconv.Atoi(parts[1])
		y, errY := strconv.Atoi(parts[2])
		if errX != nil || errY != nil {
			continue
		}

		// Check for invalid room names
		if strings.HasPrefix(parts[0], "L") || strings.HasPrefix(parts[0], "#") {
			continue
		}

		room := models.Room{
			Name:    parts[0],
			X:       x,
			Y:       y,
			IsStart: nextIsStart,
			IsEnd:   nextIsEnd,
		}

		rooms = append(rooms, room)
		nextIsStart = false
		nextIsEnd = false
	}

	// Verify start and end rooms
	hasStart := false
	hasEnd := false
	for _, room := range rooms {
		if room.IsStart {
			hasStart = true
		}
		if room.IsEnd {
			hasEnd = true
		}
	}

	if !hasStart || !hasEnd {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, missing start or end room")
	}

	return ants, rooms, links, nil
}
