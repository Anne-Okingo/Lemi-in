package utils

import (

	// "lemin/utils"
	"lemin/models"
)

func Getturns(pathwithants map[int][]int, graph *models.Graph) int {
	maximumturns := 0

	// get the number of rooms ant will visit excluding the start

	for i, path := range graph.AllPaths {
		rooms := len(path) - 1
		ants := len(pathwithants[i])

		// calculate number of turns for a path
		turns := rooms + ants - 1

		if turns > maximumturns {
			maximumturns = turns
		}

		if maximumturns >= 1000 {
			break
		}
	}

	return maximumturns
}
