package test

import (
	"os"
	"testing"

	"lemin/utils"
)

func TestParseInput(t *testing.T) {
	// Create a temporary file with valid input
	content := `3
##start
0 1 0
##end
1 5 0
2 9 0
0-2
2-1
`
	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	// Call parseInput with the temporary file
	ants, rooms, links, err := utils.ParseInput(tmpfile.Name())
	// Check for errors
	if err != nil {
		t.Errorf("parseInput returned an error: %v", err)
	}
	// Check the number of ants
	if ants != 3 {
		t.Errorf("Expected 3 ants, got %d", ants)
	}
	// Check the number of rooms
	expectedRooms := 3
	if len(rooms) != expectedRooms {
		t.Errorf("Expected %d rooms, got %d", expectedRooms, len(rooms))
	}
	// Check the number of links
	expectedLinks := 2
	if len(links) != expectedLinks {
		t.Errorf("Expected %d links, got %d", expectedLinks, len(links))
	}
	// Check if start and end rooms are correctly marked
	startRoom := rooms[0]
	if !startRoom.IsStart || startRoom.Name != "0" {
		t.Errorf("Start room not correctly identified")
	}
	endRoom := rooms[1]
	if !endRoom.IsEnd || endRoom.Name != "1" {
		t.Errorf("End room not correctly identified")
	}
	// Check if links are correctly parsed
	expectedLinkPairs := [][2]string{{"0", "2"}, {"2", "1"}}
	for i, link := range links {
		if link.From != expectedLinkPairs[i][0] || link.To != expectedLinkPairs[i][1] {
			t.Errorf("Link %d not correctly parsed. Expected %v-%v, got %v-%v",
				i, expectedLinkPairs[i][0], expectedLinkPairs[i][1], link.From, link.To)
		}
	}
}

func TestParseInput_EmptyFile(t *testing.T) {
	// Create a temporary empty file
	tmpfile, err := os.CreateTemp("", "empty_test_file")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Close the file immediately as we want it to be empty
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}
	// Call parseInput with the empty file
	ants, rooms, links, err := utils.ParseInput(tmpfile.Name())
	// Check the results
	if err == nil {
		t.Errorf("Expected error for empty file, got: %v", nil)
	}
	if ants != 0 {
		t.Errorf("Expected 0 ants for empty file, got: %d", ants)
	}
	if len(rooms) != 0 {
		t.Errorf("Expected 0 rooms for empty file, got: %d", len(rooms))
	}
	if len(links) != 0 {
		t.Errorf("Expected 0 links for empty file, got: %d", len(links))
	}
}
