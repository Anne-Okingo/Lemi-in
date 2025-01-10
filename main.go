package main

import (
	"bufio"
	"fmt"
	"os"
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
	from     string
	to       string
	NumberofAnts int
}

func parseInput(filename string) (int, []Room, []Link, error) {
contents, err := ValidContent(filename)

fmt.Println(contents)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR: oppening file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ants int
	var rooms []Room
	var links []Link
	var nextIsStart, nextIsEnd bool

	for scanner.Scan() {
		line := scanner.Text()
		
		if line == "" {
			continue
		}
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

		if ants == 0 {
			ants, err = strconv.Atoi(line)
			if err != nil || ants <= 0 {
				return 0, nil, nil, fmt.Errorf("invalid number of ants")
			}
			continue
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			links = append(links, Link{from: parts[0], to: parts[1], NumberofAnts: 1})
		} else {
			parts := strings.Fields(line)
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


func ValidContent(filename string)([]string, error){
	fileContent, err := os.Open(filename)
	if err != nil{
		return nil, fmt.Errorf("Error reading file", err)
	}
	defer fileContent.Close()

	ValidContent := []string{}

	scanner := bufio.NewScanner(fileContent)

	for scanner.Scan(){
		lines := scanner.Text()
		if (lines != "" && !strings.HasPrefix(lines, "#")) || strings.HasPrefix(lines,"##end") || strings.HasPrefix(lines, "##start"){
			ValidContent = append(ValidContent, lines)
		}
	}

	if err := scanner.Err();err != nil{
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
	
		argument:= os.Args[1]
		if !strings.HasSuffix(argument, ".txt"){
			fmt.Println("ERROR: invalid data format, inputfile must be .txt")
			return
		}
}