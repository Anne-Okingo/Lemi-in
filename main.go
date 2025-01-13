package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Room struct {
	name    string
	x, y    int
	isStart bool
	isEnd   bool
}

type Link struct {
	from         string
	to           string
	NumberofAnts int
}

func parseInput(filename string) (int, []Room, []Link, error) {
	contents, err := ValidContent(filename)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format",err)
	}

	if len(contents) == 0{
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format",err)
	}

	var rooms []Room
	var links []Link
	var nextIsStart, nextIsEnd bool

	// validate number of ants
	ants, err := strconv.Atoi(contents[0])
	if err != nil || ants <= 0 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid number of Ants")
	}

	contents = contents[1:]
	for _, str := range contents {
		if strings.Contains(str, "-") {
			parts := strings.Split(str, "-")
			links = append(links, Link{from: parts[0], to: parts[1], NumberofAnts: 1})
		} else {
			parts := strings.Fields(str)
			fmt.Println(parts)
			if len(parts) != 3 {
				continue
			}
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])
			rooms = append(rooms, Room{
				name:    parts[0],
				x:       x,
				y:       y,
				isStart: nextIsStart,
				isEnd:   nextIsEnd,
			})
			nextIsStart = false
			nextIsEnd = false
		}
	}

	return ants, rooms, links, nil
}

func ValidContent(filename string) ([]string, error) {
	fileContent, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file", err)
	}
	defer fileContent.Close()

	// Validate line content with regular expression
	ValidlineRegex := regexp.MustCompile(`^[A-Za-z0-9\s\-]+$`)

	ValidContent := []string{}

	scanner := bufio.NewScanner(fileContent)

	for scanner.Scan() {
		lines := scanner.Text()
		// ignore empty lines and comments leaving ##start and ##end
		if strings.Contains(lines, "##end") || strings.Contains(lines, "##start") {
			ValidContent = append(ValidContent, lines)
		}

		if lines != "" {
			// check if lines match the regexp
			if ValidlineRegex.MatchString(lines) {
				ValidContent = append(ValidContent, lines)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file", err)
	}

	return ValidContent, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input_file>")
		return
	}

	// Parse input
	ants, rooms, links, err := parseInput(os.Args[1])
	fmt.Println(ants)
	fmt.Println(rooms)
	fmt.Println(links)
	fmt.Println(err)

	argument := os.Args[1]
	if !strings.HasSuffix(argument, ".txt") {
		fmt.Println("ERROR: invalid data format, inputfile must be .txt")
		return
	}
}
