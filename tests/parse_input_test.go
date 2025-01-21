package test

import (
	"os"
	"strings"
	"testing"

	"lemin/utils"
)

func TestParseInputMultiple_EndRooms(t *testing.T) {
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
