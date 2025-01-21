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
