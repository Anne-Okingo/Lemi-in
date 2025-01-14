package test

import (
	"reflect"
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

func TestFindPaths(t *testing.T) {
	startRoom := "A"
	endRoom := "C"
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
	}

	expectedPaths := [][]string{{"A", "B", "C"}}

	paths := utils.FindPaths(startRoom, endRoom, links)

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d path, but got %d", len(expectedPaths), len(paths))
	}

	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Errorf("Expected paths %v, but got %v", expectedPaths, paths)
	}
}
