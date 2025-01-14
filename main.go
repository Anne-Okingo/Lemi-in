package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lemin/utils"
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



func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	graph, err := utils.ParseInput(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	graph.AllPaths = utils.GetAllPaths(graph.Rooms, graph.StartRoom, graph.EndRoom)
	if len(graph.AllPaths) == 0 {
		fmt.Println("ERROR: No valid paths found")
		return
	}

	fmt.Println("Input File Content:")
	//utils.PrintFileContents(os.Args[1])

	fmt.Println("\nSimulating Ant Movement:")
	utils.SimulateAntMovement(graph.AllPaths, graph.AntCount)
}
