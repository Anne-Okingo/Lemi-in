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
	utils.PrintFileContents(os.Args[1])

	fmt.Println("\nSimulating Ant Movement:")
	utils.SimulateAntMovement(graph.AllPaths, graph.AntCount)
}