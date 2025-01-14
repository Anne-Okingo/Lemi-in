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

func TestFindPaths_MultipleRoutes(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "A", To: "C"},
		{From: "B", To: "D"},
		{From: "C", To: "D"},
		{From: "D", To: "E"},
	}

	paths := utils.FindPaths("A", "E", links)

	expectedPaths := [][]string{
		{"A", "B", "D", "E"},
		{"A", "C", "D", "E"},
	}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for _, expectedPath := range expectedPaths {
		found := false
		for _, path := range paths {
			if compareSlices(path, expectedPath) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected path %v not found in result", expectedPath)
		}
	}
}

func compareSlices(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

func TestFindPaths_HandlesCyclicPaths(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
		{From: "C", To: "A"},
		{From: "C", To: "D"},
	}
	startRoom := "A"
	endRoom := "D"

	paths := utils.FindPaths(startRoom, endRoom, links)

	expectedPaths := [][]string{
		{"A", "B", "C", "D"},
		{"A", "C", "D"},
	}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for i, path := range paths {
		if !reflect.DeepEqual(path, expectedPaths[i]) {
			t.Errorf("Path %d: expected %v, but got %v", i, expectedPaths[i], path)
		}
	}
}

func TestFindPaths_SameStartAndEnd(t *testing.T) {
	startRoom := "A"
	endRoom := "A"
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
		{From: "C", To: "D"},
	}

	paths := utils.FindPaths(startRoom, endRoom, links)

	if len(paths) != 0 {
		t.Errorf("Expected empty slice, but got %v", paths)
	}
}

func TestFindPaths_BidirectionalLinks(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
		{From: "C", To: "D"},
	}

	paths := utils.FindPaths("A", "D", links)

	expectedPaths := [][]string{
		{"A", "B", "C", "D"},
	}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for i, path := range paths {
		if !reflect.DeepEqual(path, expectedPaths[i]) {
			t.Errorf("Path %d: expected %v, but got %v", i, expectedPaths[i], path)
		}
	}
}
