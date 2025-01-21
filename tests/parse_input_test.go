package test

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"lemin/utils"
)

func TestParseInput_MultipleEndRooms(t *testing.T) {
	// Create a temporary file with multiple end rooms
	content := `10
##start
start 0 0
##end
end1 1 1
room1 2 2
##end
end2 3 3
start-room1
room1-end1
room1-end2`

	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	// Test ParseInput function
	_, err = utils.ParseInput(tmpfile.Name())

	// Check if an error is returned
	if err == nil {
		t.Error("Expected an error for multiple end rooms, but got nil")
	}

	// Check if the error message is correct
	expectedError := "invalid data format, multiple end rooms"
	if err != nil && !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected error message to contain '%s', but got: %v", expectedError, err)
	}
}

func TestParseInput_NoAnts(t *testing.T) {
	// Create a temporary file with no ants specified
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write test data to the file
	_, err = tmpfile.WriteString("# This is a comment\n##start\nroom1 0 0\n##end\nroom2 1 1\nroom1-room2\n")
	if err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Call ParseInput with the temporary file
	_, err = utils.ParseInput(tmpfile.Name())

	// Check if an error was returned
	if err == nil {
		t.Error("Expected an error for input with no ants, but got nil")
	}

	// Check if the error message is as expected
	expectedError := "invalid data format, invalid number of ants"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestParseInput_MultipleStartRooms(t *testing.T) {
	// Create a temporary file with multiple start rooms
	content := `2
##start
A 1 1
##start
B 2 2
C 3 3
A-B
B-C`

	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	// Test ParseInput function
	_, err = utils.ParseInput(tmpfile.Name())

	// Check if the error message is correct
	expectedError := "invalid data format, multiple start rooms"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error '%s', but got: %v", expectedError, err)
	}
}

func TestParseInput_ValidFile(t *testing.T) {
	// Create a temporary file with valid input
	content := `10
##start
start 0 0
room1 1 1
room2 2 2
##end
end 3 3
start-room1
room1-room2
room2-end`

	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	// Parse the input file
	graph, err := utils.ParseInput(tmpfile.Name())
	if err != nil {
		t.Fatalf("ParseInput failed: %v", err)
	}

	// Check the parsed data
	if graph.AntCount != 10 {
		t.Errorf("Expected 10 ants, got %d", graph.AntCount)
	}

	if len(graph.Rooms) != 4 {
		t.Errorf("Expected 4 rooms, got %d", len(graph.Rooms))
	}

	if graph.StartRoom != "start" {
		t.Errorf("Expected start room to be 'start', got '%s'", graph.StartRoom)
	}

	if graph.EndRoom != "end" {
		t.Errorf("Expected end room to be 'end', got '%s'", graph.EndRoom)
	}

	expectedRooms := []struct {
		name  string
		x, y  int
		links []string
	}{
		{"start", 0, 0, []string{"room1"}},
		{"room1", 1, 1, []string{"start", "room2"}},
		{"room2", 2, 2, []string{"room1", "end"}},
		{"end", 3, 3, []string{"room2"}},
	}

	for _, er := range expectedRooms {
		room, ok := graph.Rooms[er.name]
		if !ok {
			t.Errorf("Expected room '%s' not found", er.name)
			continue
		}
		if room.XCoordinate != er.x || room.YCoordinate != er.y {
			t.Errorf("Room '%s' coordinates mismatch. Expected (%d, %d), got (%d, %d)", er.name, er.x, er.y, room.XCoordinate, room.YCoordinate)
		}
		if !reflect.DeepEqual(room.Links, er.links) {
			t.Errorf("Room '%s' links mismatch. Expected %v, got %v", er.name, er.links, room.Links)
		}
	}
}

func TestParseInput_RoomNameStartingWithL(t *testing.T) {
	// Create a temporary file with invalid input
	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write test data to the temporary file
	testData := `3
##start
Lroom 1 1
##end
room2 2 2
Lroom-room2
`
	if _, err := tmpfile.Write([]byte(testData)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Call the function with the temporary file
	_, err = utils.ParseInput(tmpfile.Name())

	// Check if the error message is as expected
	expectedError := "invalid data format, room name cannot start with 'L' or '#'"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error '%s', but got: %v", expectedError, err)
	}
}
