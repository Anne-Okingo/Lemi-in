package main

import (
	"fmt"
	"os"

	"lemin/utils"
)

func main() {
	// Check for valid number command-line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	// Check is file parsed is a text file
	if !utils.ValidFile(os.Args[1]) {
		fmt.Println("ERROR: Wrong file format")
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

	utils.PrintFileContents(os.Args[1])

	utils.SimulateAntMovement(graph.AllPaths, graph.AntCount)
}
