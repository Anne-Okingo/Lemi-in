package utils_test

import (
	"os"
	"strings"
	"testing"

	"lemin/utils"
)

func TestParseInput(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		content := `3
##start
A 1 2
##end
B 3 4
C 5 6
A-C
C-B
`
		tmpfile := createTempFile(t, content)
		defer os.Remove(tmpfile)

		graph, err := utils.ParseInput(tmpfile)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Validate AntCount
		if graph.AntCount != 3 {
			t.Errorf("Expected AntCount to be 3, got %d", graph.AntCount)
		}

		// Validate Rooms
		if len(graph.Rooms) != 3 {
			t.Errorf("Expected 3 rooms, got %d", len(graph.Rooms))
		}

		// Validate StartRoom and EndRoom
		if graph.StartRoom != "A" {
			t.Errorf("Expected StartRoom to be 'A', got %s", graph.StartRoom)
		}
		if graph.EndRoom != "B" {
			t.Errorf("Expected EndRoom to be 'B', got %s", graph.EndRoom)
		}

		// Validate Links
		expectedLinks := map[string][]string{
			"A": {"C"},
			"C": {"A", "B"},
			"B": {"C"},
		}
		for room, expected := range expectedLinks {
			actual, exists := graph.Rooms[room]
			if !exists {
				t.Errorf("Room %s not found", room)
				continue
			}
			if len(actual.Links) != len(expected) {
				t.Errorf("Expected %d links for room %s, got %d", len(expected), room, len(actual.Links))
			}
		}
	})

	t.Run("InvalidAntCount", func(t *testing.T) {
		content := `-1
##start
A 1 2
##end
B 3 4
A-B
`
		tmpfile := createTempFile(t, content)
		defer os.Remove(tmpfile)

		_, err := utils.ParseInput(tmpfile)
		if err == nil || err.Error() != "invalid data format, invalid number of ants" {
			t.Errorf("Expected invalid number of ants error, got %v", err)
		}
	})

	t.Run("DuplicateRoom", func(t *testing.T) {
		content := `3
##start
A 1 2
A 3 4
B 5 6
A-B
`
		tmpfile := createTempFile(t, content)
		defer os.Remove(tmpfile)

		_, err := utils.ParseInput(tmpfile)
		if err == nil || !strings.Contains(err.Error(), "duplicate room") {
			t.Errorf("Expected duplicate room error, got %v", err)
		}
	})

	t.Run("InvalidLink", func(t *testing.T) {
		content := `3
##start
A 1 2
##end
B 3 4
A-X
`
		tmpfile := createTempFile(t, content)
		defer os.Remove(tmpfile)

		_, err := utils.ParseInput(tmpfile)
		if err == nil || !strings.Contains(err.Error(), "link references unknown room") {
			t.Errorf("Expected unknown room link error, got %v", err)
		}
	})

	t.Run("NoStartRoom", func(t *testing.T) {
		content := `3
B 1 2
C 3 4
B-C
`
		tmpfile := createTempFile(t, content)
		defer os.Remove(tmpfile)

		_, err := utils.ParseInput(tmpfile)
		if err == nil || !strings.Contains(err.Error(), "no start room found") {
			t.Errorf("Expected no start room error, got %v", err)
		}
	})
}

// Helper function to create temporary file with content
func createTempFile(t *testing.T, content string) string {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	return tmpfile.Name()
}
