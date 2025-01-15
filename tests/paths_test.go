package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

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

func TestFindPaths_MultipleLinksFromOneRoom(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "A", To: "C"},
		{From: "B", To: "D"},
		{From: "C", To: "D"},
		{From: "D", To: "E"},
	}

	startRoom := "A"
	endRoom := "E"

	expectedPaths := [][]string{
		{"A", "B", "D", "E"},
		{"A", "C", "D", "E"},
	}

	paths := utils.FindPaths(startRoom, endRoom, links)

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for _, expectedPath := range expectedPaths {
		found := false
		for _, path := range paths {
			if reflect.DeepEqual(path, expectedPath) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected path %v not found in result", expectedPath)
		}
	}
}

func TestFindPaths_LargeMaze(t *testing.T) {
	// Create a large number of rooms and links
	numRooms := 1000
	links := make([]models.Link, numRooms-1)
	for i := 0; i < numRooms-1; i++ {
		links[i] = models.Link{
			From: fmt.Sprintf("room%d", i),
			To:   fmt.Sprintf("room%d", i+1),
		}
	}

	startRoom := "room0"
	endRoom := fmt.Sprintf("room%d", numRooms-1)

	// Measure execution time
	start := time.Now()
	paths := utils.FindPaths(startRoom, endRoom, links)
	duration := time.Since(start)

	// Check if the function completes within a reasonable time (e.g., 1 second)
	if duration > time.Second {
		t.Errorf("FindPaths took too long: %v", duration)
	}

	// Verify the result
	if len(paths) != 1 {
		t.Errorf("Expected 1 path, got %d", len(paths))
	}
	if len(paths[0]) != numRooms {
		t.Errorf("Expected path length of %d, got %d", numRooms, len(paths[0]))
	}
}

func TestFindPaths_IsolatedRooms(t *testing.T) {
	links := []models.Link{
		{From: "start", To: "A"},
		{From: "A", To: "B"},
		{From: "B", To: "end"},
		{From: "C", To: "D"}, // Isolated path
	}
	paths := utils.FindPaths("start", "end", links)

	expected := [][]string{
		{"start", "A", "B", "end"},
	}

	if len(paths) != len(expected) {
		t.Fatalf("Expected %d paths, but got %d", len(expected), len(paths))
	}

	for i, path := range paths {
		if !reflect.DeepEqual(path, expected[i]) {
			t.Errorf("Path %d: expected %v, but got %v", i, expected[i], path)
		}
	}
}

func TestFindPaths_MaintainRoomOrder(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
		{From: "A", To: "D"},
		{From: "D", To: "C"},
	}
	startRoom := "A"
	endRoom := "C"

	paths := utils.FindPaths(startRoom, endRoom, links)

	expectedPaths := [][]string{
		{"A", "B", "C"},
		{"A", "D", "C"},
	}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for i, expectedPath := range expectedPaths {
		if !reflect.DeepEqual(paths[i], expectedPath) {
			t.Errorf("Path %d: expected %v, but got %v", i, expectedPath, paths[i])
		}
	}
}
