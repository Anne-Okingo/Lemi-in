package test

import (
	"testing"

	"lemin/models"
	"lemin/utils"
)

func TestFindPaths_NoPath(t *testing.T) {
	startRoom := "A"
	endRoom := "C"
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "D"},
	}

	paths := utils.FindPaths(startRoom, endRoom, links)

	if len(paths) != 0 {
		t.Errorf("Expected empty slice, but got %v", paths)
	}
}
