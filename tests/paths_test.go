package test

import (
	"reflect"
	"testing"

	"lemin/models"
	"lemin/utils"
)

func TestGetAllPaths_NoValidPaths(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"A": {Name: "A", Links: []string{"B"}},
		"B": {Name: "B", Links: []string{"A"}},
		"C": {Name: "C", Links: []string{}},
	}

	paths := utils.GetAllPaths(rooms, "A", "C")

	if len(paths) != 0 {
		t.Errorf("Expected empty slice, but got %v paths", len(paths))
	}
}

func TestGetAllPaths_SingleDirectPath(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"start": {Links: []string{"end"}},
		"end":   {Links: []string{"start"}},
	}

	paths := utils.GetAllPaths(rooms, "start", "end")

	expectedPaths := [][]string{{"start", "end"}}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d path, got %d", len(expectedPaths), len(paths))
	}

	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Errorf("Expected paths %v, got %v", expectedPaths, paths)
	}
}

// func TestGetAllPaths(t *testing.T) {
// 	rooms := map[string]*models.ARoom{
// 		"start": {Links: []string{"A", "B"}},
// 		"A":     {Links: []string{"start", "C", "D"}},
// 		"B":     {Links: []string{"start", "D"}},
// 		"C":     {Links: []string{"A", "end"}},
// 		"D":     {Links: []string{"A", "B", "end"}},
// 		"end":   {Links: []string{"C", "D"}},
// 	}

// 	paths := utils.GetAllPaths(rooms, "start", "end")
// 	log.Printf("Paths: %#v", paths)

// 	expectedPaths := [][]string{
// 		{"start", "A", "C", "end"},
// 		{"start", "A", "D", "end"},
// 		{"start", "B", "D", "end"},
// 	}

// 	if len(paths) != len(expectedPaths) {
// 		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
// 	}

// 	for _, expectedPath := range expectedPaths {
// 		found := false
// 		for _, actualPath := range paths {
// 			if equalPaths(expectedPath, actualPath) {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			t.Errorf("Expected path %v not found in the result", expectedPath)
// 		}
// 	}
// }

func equalPaths(path1, path2 []string) bool {
	if len(path1) != len(path2) {
		return false
	}
	for i := range path1 {
		if path1[i] != path2[i] {
			return false
		}
	}
	return true
}

func TestGetAllPaths_WithCycles(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"A": {Name: "A", Links: []string{"B", "C"}},
		"B": {Name: "B", Links: []string{"A", "D"}},
		"C": {Name: "C", Links: []string{"A", "D"}},
		"D": {Name: "D", Links: []string{"B", "C", "E"}},
		"E": {Name: "E", Links: []string{"D"}},
	}

	paths := utils.GetAllPaths(rooms, "A", "E")

	expectedPaths := [][]string{
		{"A", "B", "D", "E"},
		{"A", "C", "D", "E"},
	}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for _, expectedPath := range expectedPaths {
		found := false
		for _, actualPath := range paths {
			if reflect.DeepEqual(expectedPath, actualPath) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected path %v not found in result", expectedPath)
		}
	}
}

func TestGetAllPaths_SameStartEnd(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"A": {Name: "A", Links: []string{}},
	}

	paths := utils.GetAllPaths(rooms, "A", "A")

	if len(paths) != 1 {
		t.Errorf("Expected 1 path, got %d", len(paths))
	}

	if len(paths[0]) != 1 || paths[0][0] != "A" {
		t.Errorf("Expected path [A], got %v", paths[0])
	}
}

func TestGetAllPaths_WithNoOutgoingLinks(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"start":  {Links: []string{"middle"}},
		"middle": {Links: []string{}},
		"end":    {Links: []string{}},
	}

	paths := utils.GetAllPaths(rooms, "start", "end")

	if len(paths) != 0 {
		t.Errorf("Expected 0 paths, but got %d", len(paths))
	}
}

func TestGetAllPaths_WithMultipleLinks(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"start": {Links: []string{"A", "B"}},
		"A":     {Links: []string{"C", "D"}},
		"B":     {Links: []string{"D", "E"}},
		"C":     {Links: []string{"end"}},
		"D":     {Links: []string{"end"}},
		"E":     {Links: []string{"end"}},
		"end":   {Links: []string{}},
	}

	paths := utils.GetAllPaths(rooms, "start", "end")

	expectedPaths := [][]string{
		{"start", "A", "C", "end"},
		{"start", "A", "D", "end"},
		{"start", "B", "D", "end"},
		{"start", "B", "E", "end"},
	}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for _, expectedPath := range expectedPaths {
		found := false
		for _, actualPath := range paths {
			if equalPaths(expectedPath, actualPath) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected path %v not found in result", expectedPath)
		}
	}
}

func TestGetAllPaths_StartRoomDoesNotExist(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"A": {Name: "A", Links: []string{"B"}},
		"B": {Name: "B", Links: []string{"A", "C"}},
		"C": {Name: "C", Links: []string{"B"}},
	}

	paths := utils.GetAllPaths(rooms, "NonExistentStart", "C")

	if len(paths) != 0 {
		t.Errorf("Expected empty slice, but got %v paths", len(paths))
	}
}

func TestGetAllPaths_EndRoomDoesNotExist(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"start": {Name: "start", Links: []string{"A", "B"}},
		"A":     {Name: "A", Links: []string{"start", "B"}},
		"B":     {Name: "B", Links: []string{"start", "A"}},
	}

	paths := utils.GetAllPaths(rooms, "start", "end")

	if len(paths) != 0 {
		t.Errorf("Expected empty slice, got %v", paths)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		path     []string
		room     string
		expected bool
	}{
		{
			name:     "Room exists in path",
			path:     []string{"room1", "room2", "room3"},
			room:     "room2",
			expected: true,
		},
		{
			name:     "Room does not exist in path",
			path:     []string{"room1", "room2", "room3"},
			room:     "room4",
			expected: false,
		},
		{
			name:     "Empty path",
			path:     []string{},
			room:     "room1",
			expected: false,
		},
		{
			name:     "Room exists at the beginning of path",
			path:     []string{"room1", "room2", "room3"},
			room:     "room1",
			expected: true,
		},
		{
			name:     "Room exists at the end of path",
			path:     []string{"room1", "room2", "room3"},
			room:     "room3",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.Contains(tt.path, tt.room)
			if result != tt.expected {
				t.Errorf("Contains(%v, %q) = %v; want %v", tt.path, tt.room, result, tt.expected)
			}
		})
	}
}
