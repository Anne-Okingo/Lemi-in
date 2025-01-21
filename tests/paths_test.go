package test

import (
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
