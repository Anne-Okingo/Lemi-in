package test

import (
	"os"
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
