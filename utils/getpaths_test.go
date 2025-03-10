package utils

import (
	"reflect"
	"testing"

	"lemin/models"
)

func TestGetAllPaths_NoValidPaths(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"A": {Name: "A", Links: []string{"B"}},
		"B": {Name: "B", Links: []string{"A"}},
		"C": {Name: "C", Links: []string{}},
	}

	paths := GetAllPaths(rooms, "A", "C")

	if len(paths) != 0 {
		t.Errorf("Expected empty slice, but got %v paths", len(paths))
	}
}

func TestGetAllPaths_SingleDirectPath(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"start": {Links: []string{"end"}},
		"end":   {Links: []string{"start"}},
	}

	paths := GetAllPaths(rooms, "start", "end")

	expectedPaths := [][]string{{"start", "end"}}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d path, got %d", len(expectedPaths), len(paths))
	}

	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Errorf("Expected paths %v, got %v", expectedPaths, paths)
	}
}

func TestGetAllPaths_SameStartEnd(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"A": {Name: "A", Links: []string{}},
	}

	paths := GetAllPaths(rooms, "A", "A")

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

	paths := GetAllPaths(rooms, "start", "end")

	if len(paths) != 0 {
		t.Errorf("Expected 0 paths, but got %d", len(paths))
	}
}


func TestGetAllPaths_StartRoomDoesNotExist(t *testing.T) {
	rooms := map[string]*models.ARoom{
		"A": {Name: "A", Links: []string{"B"}},
		"B": {Name: "B", Links: []string{"A", "C"}},
		"C": {Name: "C", Links: []string{"B"}},
	}

	paths := GetAllPaths(rooms, "NonExistentStart", "C")

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

	paths := GetAllPaths(rooms, "start", "end")

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
			result := Contains(tt.path, tt.room)
			if result != tt.expected {
				t.Errorf("Contains(%v, %q) = %v; want %v", tt.path, tt.room, result, tt.expected)
			}
		})
	}
}
